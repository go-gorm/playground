package main

import (
	"strconv"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type unsignedAutoIncrementRecord struct {
	ID   uint64 `gorm:"primaryKey;autoIncrement"`
	Name string
}

func TestGORM(t *testing.T) {
	if DB.Dialector.Name() != "mysql" {
		t.Skip("MySQL-specific LastInsertId reproduction")
	}

	const (
		tableName  = "unsigned_auto_increment_records"
		expectedID = uint64(1) << 63
	)

	if err := DB.Exec("DROP TABLE IF EXISTS `" + tableName + "`").Error; err != nil {
		t.Fatalf("failed to drop test table: %v", err)
	}
	t.Cleanup(func() {
		DB.Exec("DROP TABLE IF EXISTS `" + tableName + "`")
	})

	if err := DB.Exec(`CREATE TABLE ` + "`" + tableName + "`" + ` (
        id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
        name VARCHAR(32) NOT NULL,
        PRIMARY KEY (id)
    ) AUTO_INCREMENT = 9223372036854775808`).Error; err != nil {
		t.Fatalf("failed to create test table: %v", err)
	}

	record := &unsignedAutoIncrementRecord{Name: "test"}
	result := DB.Table(tableName).CreateInBatches([]*unsignedAutoIncrementRecord{record}, 100)
	if result.Error != nil {
		t.Fatalf("create returned an unexpected error: %v", result.Error)
	}
	if result.RowsAffected != 1 {
		t.Fatalf("expected one affected row, got %d", result.RowsAffected)
	}

	var persistedID string
	if err := DB.Raw("SELECT CAST(id AS CHAR) FROM `"+tableName+"` WHERE name = ?", record.Name).Scan(&persistedID).Error; err != nil {
		t.Fatalf("failed to query the inserted ID: %v", err)
	}
	if persistedID != strconv.FormatUint(expectedID, 10) {
		t.Fatalf("expected database ID %d, got %s", expectedID, persistedID)
	}
	if record.ID != expectedID {
		t.Fatalf("database inserted ID %s, but GORM populated ID %d", persistedID, record.ID)
	}
}

// func TestGORMGen(t *testing.T) {
// 	user := models.User{Name: "jinzhu2"}
// 	ctx := context.Background()

// 	gorm.G[models.User](DB).Create(ctx, &user)

// 	if u, err := gorm.G[models.User](DB).Where(g.User.ID.Eq(user.ID)).First(ctx); err != nil {
// 		t.Errorf("Failed, got error: %v", err)
// 	} else if u.Name != user.Name {
// 		t.Errorf("Failed, got user name: %v", u.Name)
// 	}
// }
