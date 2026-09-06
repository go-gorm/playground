package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gorm.io/gorm"
	"gorm.io/playground/models"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	if strings.ToLower(DB.Name()) != "postgres" {
		t.Skip("Only runs on Postgres")
	}
	t.Run("DropIndexByField", func(t *testing.T) {
		assert.NoError(t, DB.Migrator().DropIndex(&models.User{}, "Name"))
		assert.False(t, DB.Migrator().HasIndex(&models.User{}, "Name"))
		t.Cleanup(func() {
			require.NoError(t, DB.AutoMigrate(&models.User{}))
			assert.True(t, DB.Migrator().HasIndex(&models.User{}, "Name"))
		})
	})
	t.Run("DropIndexByName", func(t *testing.T) {
		assert.NoError(t, DB.Migrator().DropIndex(&models.User{}, "idx_name"))
		assert.False(t, DB.Migrator().HasIndex(&models.User{}, "Name"))
		t.Cleanup(func() {
			require.NoError(t, DB.AutoMigrate(&models.User{}))
			assert.True(t, DB.Migrator().HasIndex(&models.User{}, "Name"))
		})
	})
	t.Run("RenameIndexByField", func(t *testing.T) {
		t.Skip("Although documented https://gorm.io/docs/migration.html#Indexes this DOES NOT WORK!")
		// Do not do the following, as there is a column Name2, and it will be chosen when trying to find the index!
		// assert.NoError(t, DB.Migrator().RenameIndex(&models.User{}, "Name", "Name2"))
		assert.NoError(t, DB.Migrator().RenameIndex(&models.User{}, "Name", "idx_name2"))
		assert.False(t, DB.Migrator().HasIndex(&models.User{}, "Name"))
		assert.True(t, DB.Migrator().HasIndex(&models.User{}, "idx_name2"))      // Have to use the real name.  Name2 will resolve to the index on the Name2 column
		require.NoError(t, DB.Migrator().DropIndex(&models.User{}, "idx_name2")) // Have to use the real name.  Name2 will resolve to the index on the Name2 column
		assert.False(t, DB.Migrator().HasIndex(&models.User{}, "idx_name2"))     // Have to use the real name.  Name2 will resolve to the index on the Name2 column
		t.Cleanup(func() {
			require.NoError(t, DB.AutoMigrate(&models.User{}))
			assert.True(t, DB.Migrator().HasIndex(&models.User{}, "Name"))
		})
	})
	t.Run("RenameIndexByFieldToExisting", func(t *testing.T) {
		t.Skip("Although documented https://gorm.io/docs/migration.html#Indexes this DOES NOT WORK!")
		err := DB.Migrator().RenameIndex(&models.User{}, "Name", "idx_name_2") // Attempt to rename to an existing index name (on Name2)
		if !assert.Error(t, err) {
			t.Fatal()
		}
		assert.Contains(t, err.Error(), "failed to rename index")
		assert.True(t, DB.Migrator().HasIndex(&models.User{}, "Name"))
		assert.True(t, DB.Migrator().HasIndex(&models.User{}, "Name2"))
		assert.True(t, DB.Migrator().HasIndex(&models.User{}, "idx_name_2"))
	})
	t.Run("RenameIndexByName", func(t *testing.T) {
		assert.NoError(t, DB.Migrator().RenameIndex(&models.User{}, "idx_name", "idx_name2"))
		assert.False(t, DB.Migrator().HasIndex(&models.User{}, "Name"))
		assert.True(t, DB.Migrator().HasIndex(&models.User{}, "Name2"))
		t.Cleanup(func() {
			require.NoError(t, DB.AutoMigrate(&models.User{}))
			assert.True(t, DB.Migrator().HasIndex(&models.User{}, "Name"))
		})
	})
	t.Run("TestWhatHappensWithSchemas", func(t *testing.T) {
		assert.NoError(t, DB.Migrator().AutoMigrate(&PublicSchema{}))
		// CREATE TABLE "public"."public_schema" ("id" bigserial,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz,"std_column" text,PRIMARY KEY ("id"))
		assert.NoError(t, DB.Migrator().DropIndex(&PublicSchema{}, "StdColumn"))
		// DROP INDEX "public"."idx_public_public_schema_std_column"
		assert.NoError(t, DB.Migrator().CreateIndex(&PublicSchema{}, "StdColumn"))
		// CREATE INDEX IF NOT EXISTS "idx_public_public_schema_std_column" ON "public"."public_schema" ("std_column")
		assert.NoError(t, DB.Migrator().RenameIndex(&PublicSchema{}, "idx_public_public_schema_std_column", "idx_public_public_schema_std_column2"))
		// ALTER INDEX "public"."idx_public_public_schema_std_column" RENAME TO "idx_public_public_schema_std_column2"
		assert.NoError(t, DB.Migrator().RenameIndex(&PublicSchema{}, "idx_public_public_schema_std_column2", "idx_public_public_schema_std_column"))
		// ALTER INDEX "public"."idx_public_public_schema_std_column2" RENAME TO "idx_public_public_schema_std_column"
		assert.Error(t, DB.Migrator().AutoMigrate(&DotSchema{}))
		// CREATE TABLE "."dot_schema" ("id" bigserial,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz,"std_column" text,PRIMARY KEY ("id")) // ERROR: syntax error at or near "dot_schema" (SQLSTATE 42601)
		require.NoError(t, DB.Exec("set search_path = ''").Error)
		// set search_path = ''
		require.Error(t, DB.Migrator().AutoMigrate(&NoSchema{}))
		// CREATE TABLE "no_schema" ("id" bigserial,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz,"std_column" text,PRIMARY KEY ("id")) // ERROR: no schema has been selected to create in (SQLSTATE 3F000)
		require.NoError(t, DB.Exec("set search_path = DEFAULT").Error)
		// set search_path = DEFAULT
		require.NoError(t, DB.Migrator().AutoMigrate(&NoSchema{}))
		// CREATE TABLE "no_schema" ("id" bigserial,"created_at" timestamptz,"updated_at" timestamptz,"deleted_at" timestamptz,"std_column" text,PRIMARY KEY ("id"))
		assert.NoError(t, DB.Migrator().DropIndex(&NoSchema{}, "StdColumn"))
		// DROP INDEX "public"."idx_no_schema_std_column"
		assert.NoError(t, DB.Migrator().CreateIndex(&NoSchema{}, "StdColumn"))
		// CREATE INDEX IF NOT EXISTS "idx_no_schema_std_column" ON "no_schema" ("std_column")
		assert.NoError(t, DB.Migrator().RenameIndex(&NoSchema{}, "idx_no_schema_std_column", "idx_no_schema_std_column2"))
		// ALTER INDEX "public"."idx_no_schema_std_column" RENAME TO "idx_no_schema_std_column2"
		assert.NoError(t, DB.Migrator().RenameIndex(&NoSchema{}, "idx_no_schema_std_column2", "idx_no_schema_std_column"))
		// ALTER INDEX "public"."idx_no_schema_std_column2" RENAME TO "idx_no_schema_std_column"
		require.NoError(t, DB.Exec("set search_path = ''").Error)
		// set search_path = ''
		assert.Error(t, DB.Migrator().DropIndex(&NoSchema{}, "StdColumn"))
		// DROP INDEX "idx_no_schema_std_column" // ERROR: index "idx_no_schema_std_column" does not exist (SQLSTATE 42704)
		require.NoError(t, DB.Exec("set search_path = DEFAULT").Error)
		// set search_path = DEFAULT
		assert.NoError(t, DB.Migrator().DropIndex(&NoSchema{}, "StdColumn"))
		// DROP INDEX "public"."idx_no_schema_std_column"
		t.Cleanup(func() {
			if DB.Migrator().HasTable(&PublicSchema{}) { // SELECT count(*) FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'public_schema' AND table_type = 'BASE TABLE'
				require.NoError(t, DB.Migrator().DropTable(&PublicSchema{}))
				// DROP TABLE IF EXISTS "public"."public_schema" CASCADE
			}
			if DB.Migrator().HasTable(&NoSchema{}) { // SELECT count(*) FROM information_schema.tables WHERE table_schema = 'public' AND table_name = 'no_schema' AND table_type = 'BASE TABLE'
				require.NoError(t, DB.Migrator().DropTable(&NoSchema{}))
				// DROP TABLE IF EXISTS "no_schema" CASCADE
			}
			if DB.Migrator().HasTable(&DotSchema{}) { // SELECT count(*) FROM information_schema.tables WHERE table_schema = '' AND table_name = 'dot_schema' AND table_type = 'BASE TABLE'
				require.NoError(t, DB.Migrator().DropTable(&DotSchema{}))
			}
		})
	})
}

type NoSchema struct {
	gorm.Model
	StdColumn string `gorm:"index"`
}

func (t *NoSchema) TableName() string {
	return "no_schema"
}

type DotSchema struct {
	gorm.Model
	StdColumn string `gorm:"index"`
}

func (t *DotSchema) TableName() string {
	return ".dot_schema"
}

type PublicSchema struct {
	gorm.Model
	StdColumn string `gorm:"index"`
}

func (t *PublicSchema) TableName() string {
	return "public.public_schema"
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
