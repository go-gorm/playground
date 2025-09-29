package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"golang.org/x/sync/errgroup"
	gormpostgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Standalone reproduction of the behavior demonstrated in the suite test
// Test_UpdateColums_insert_deleted_associations: when using UpdateColumns on the
// parent, GORM re-inserts deleted child associations (deleted concurrently) due to
// association handling in that update path.

type Parent struct {
	gorm.Model
	Name     string
	Children []*Child `gorm:"foreignKey:ParentID"`
}
type Child struct {
	gorm.Model
	ParentID uint `gorm:"type:uuid;primaryKey"` // composite PK (ParentID, ID)
}

func setupDB(t *testing.T) *gorm.DB {
	t.Helper()
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "postgres:16-alpine",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_PASSWORD": "test",
			"POSTGRES_USER":     "test",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForAll(
			wait.ForListeningPort("5432/tcp"),
			wait.ForLog("database system is ready to accept connections"),
		).WithDeadline(5 * time.Second),
	}

	pgC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{ContainerRequest: req, Started: true})
	if err != nil {
		t.Fatalf("failed to start postgres container: %v", err)
	}
	// ensure cleanup
	t.Cleanup(func() { _ = pgC.Terminate(context.Background()) })

	host, err := pgC.Host(ctx)
	if err != nil {
		t.Fatalf("container host: %v", err)
	}
	port, err := pgC.MappedPort(ctx, "5432/tcp")
	if err != nil {
		t.Fatalf("container mapped port: %v", err)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=test password=test dbname=testdb sslmode=disable TimeZone=UTC", host, port.Port())

	// retry opening (container might not yet be fully accepting connections even after wait strategy)
	var gormDB *gorm.DB

	gormDB, err = gorm.Open(postgresdriver(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err == nil {
		// ensure low-level ping succeeds
		sqldb, derr := gormDB.DB()
		if derr == nil {
			pingErr := sqldb.PingContext(ctx)
			err = errors.Join(err, pingErr)
		}
	}

	if err := gormDB.AutoMigrate(&Parent{}, &Child{}); err != nil {
		t.Fatalf("failed to automigrate: %v", err)
	}
	return gormDB
}

// Helper to satisfy linter (explicit name to avoid import alias confusion if we later add other drivers)
func postgresdriver(dsn string) gorm.Dialector { return gormpostgres.Open(dsn) }

func Test_UpdateColums_insert_deleted_associations(t *testing.T) { // keep original name (typo included) for clarity
	DB = setupDB(t)

	parent := Parent{Name: "parent1"}
	if err := DB.Create(&parent).Error; err != nil {
		t.Fatalf("create parent: %v", err)
	}

	child1 := Child{ParentID: parent.ID}
	child2 := Child{ParentID: parent.ID}
	if err := DB.Create(&child1).Error; err != nil {
		t.Fatalf("create child1: %v", err)
	}
	if err := DB.Create(&child2).Error; err != nil {
		t.Fatalf("create child2: %v", err)
	}

	tx2Done := make(chan struct{})
	tx1loaded := make(chan struct{})
	start := make(chan struct{})
	var eg errgroup.Group

	// Transaction 1: load parent & children, then after deletion attempt UpdateColumns
	eg.Go(func() error {
		<-start
		return DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Preload("Children").First(&parent, "id = ?", parent.ID).Error; err != nil {
				return err
			}
			if len(parent.Children) != 2 {
				return fmt.Errorf("expected 2 children, got %d", len(parent.Children))
			}
			close(tx1loaded)
			<-tx2Done // wait for delete
			if err := tx.Model(&parent).UpdateColumns(map[string]interface{}{"name": "parent1-updated"}).Error; err != nil {
				return err
			}
			return nil
		}, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	})

	// Transaction 2: delete one child (unscoped)
	eg.Go(func() error {
		defer close(tx2Done)
		<-start
		<-tx1loaded
		return DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Unscoped().Delete(&Child{}, "parent_id = ? AND id = ?", parent.ID, child1.ID).Error; err != nil {
				return err
			}
			return nil
		}, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
	})

	close(start)
	if err := eg.Wait(); err != nil {
		if testing.Verbose() {
			t.Logf("errgroup wait returned error: %v", err)
		}
	}

	var children []Child
	if err := DB.Find(&children, "parent_id = ?", parent.ID).Error; err != nil {
		t.Fatalf("query children: %v", err)
	}
	if len(children) != 2 {
		ids := make([]uint, 0, len(children))
		for _, c := range children {
			ids = append(ids, c.ID)
		}
		// Keep concise fatal now
		// Behavior changed or anomaly not reproduced
		t.Fatalf("expected 2 children (resurrection) got %d; child IDs: %v", len(children), ids)
	}
	if err := DB.Preload("Children").First(&parent, "id = ?", parent.ID).Error; err != nil {
		t.Fatalf("query parent: %v", err)
	}
	if parent.Name != "parent1-updated" {
		t.Fatalf("expected parent name to be 'parent1-updated', got '%s'", parent.Name)
	}
}
