package main

import (
	"testing"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

var (
	postgresDSN = "user=gorm password=gorm dbname=gorm host=localhost port=9920 sslmode=disable TimeZone=Asia/Shanghai"
)

// https://github.com/go-gorm/gorm/blob/654b5f20066737fd7a7e62662b12bdf9cedba178/tests/migrate_test.go#L257
func TestMigrateWithUniqueIndex(t *testing.T) {
	// DB, err := gorm.Open(mysql.Open(mysqlDSN), &gorm.Config{})
	// DB, err := gorm.Open(sqlite.Open(filepath.Join(os.TempDir(), "gorm.db")), &gorm.Config{})
	DB, err := gorm.Open(postgres.Open(postgresDSN), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open postgres database, got error %v", err)
	}
	type UserWithUniqueIndex struct {
		ID    int
		Name  string    `gorm:"size:20;index:idx_name,unique"`
		Date  time.Time `gorm:"index:idx_name,unique"`
		UName string    `gorm:"uniqueIndex;size:255"`
		Token string    `gorm:"index:uidx_user_with_unique_indices_token,unique,where:token!=''"`
	}

	DB.Migrator().DropTable(&UserWithUniqueIndex{})
	if err := DB.AutoMigrate(&UserWithUniqueIndex{}); err != nil {
		t.Fatalf("failed to migrate, got %v", err)
	}

	if !DB.Migrator().HasIndex(&UserWithUniqueIndex{}, "idx_name") {
		t.Errorf("Failed to find created index")
	}

	if !DB.Migrator().HasIndex(&UserWithUniqueIndex{}, "idx_user_with_unique_indices_u_name") {
		t.Errorf("Failed to find created index")
	}

	if err := DB.AutoMigrate(&UserWithUniqueIndex{}); err != nil {
		t.Fatalf("failed to migrate, got %v", err)
	}

	if !DB.Migrator().HasIndex(&UserWithUniqueIndex{}, "idx_user_with_unique_indices_u_name") {
		t.Errorf("Failed to find created index")
	}

	if !DB.Migrator().HasIndex(&UserWithUniqueIndex{}, "uidx_user_with_unique_indices_token") ||
		DB.Migrator().HasIndex(&UserWithUniqueIndex{}, "idx_user_with_unique_indices_token") {
		t.Errorf("Failed to find created index")
	}
}
