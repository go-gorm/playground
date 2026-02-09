package main

import (
	"testing"

	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
	"gorm.io/playground/models"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	userId := uuid.NullUUID{Valid: true, UUID: uuid.Must(uuid.NewV4())}
	projectId := uuid.Must(uuid.NewV4())
	layouts := []models.Layout{}

	expectedSQL := "SELECT * FROM `layouts` WHERE `layouts`.`project_id` = \"" +
		projectId.String() + "\" AND (private_for = \"" + userId.UUID.String() + "\" OR private_for IS NULL)"

	t.Run("OR as function", func(t *testing.T) {
		sql := DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
			byUser := func(dbScope *gorm.DB) *gorm.DB {
				if userId.Valid {
					return dbScope.Where(dbScope.Where("private_for = ?", userId).Or("private_for IS NULL"))
				}
				return dbScope
			}
			return tx.Table("layouts").Scopes(byUser).Find(&layouts, map[string]any{"project_id": projectId})
		})
		if sql != expectedSQL {
			t.Errorf("Expected SQL: %v, got: %v", expectedSQL, sql)
		}
	})

	t.Run("OR as raw SQL", func(t *testing.T) {
		sql := DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
			byUser := func(dbScope *gorm.DB) *gorm.DB {
				if userId.Valid {
					return dbScope.Where("private_for = ? OR private_for IS NULL", userId)
				}
				return dbScope
			}
			return tx.Table("layouts").Scopes(byUser).Find(&layouts, map[string]any{"project_id": projectId})
		})
		if sql != expectedSQL {
			t.Errorf("Expected SQL: %v, got: %v", expectedSQL, sql)
		}
	})

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
