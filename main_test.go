package main

import (
	"testing"
	"time"

	. "gorm.io/playground/models"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: v1.31.2
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

// Config matches the empty config used by TestAssociationMany2ManyAppendMap.
type Config struct{}

func GetUser(name string, _ Config) *User {
	birthday := time.Now().Round(time.Second)
	return &User{
		Name:     name,
		Age:      18,
		Birthday: &birthday,
	}
}

func AssertAssociationCount(t *testing.T, data interface{}, name string, result int64, reason string) {
	t.Helper()
	if count := DB.Model(data).Association(name).Count(); count != result {
		t.Fatalf("invalid %v count %v, expects: %v got %v", name, reason, result, count)
	}

	var newUser User
	if user, ok := data.(User); ok {
		DB.Find(&newUser, user.ID)
	} else if user, ok := data.(*User); ok {
		DB.Find(&newUser, user.ID)
	}

	if newUser.ID != 0 {
		if count := DB.Model(&newUser).Association(name).Count(); count != result {
			t.Fatalf("invalid %v count %v, expects: %v got %v", name, reason, result, count)
		}
	}
}

// TestAssociationMany2ManyAppendMap fails with NamingStrategy.NoLowerCase.
// Column names stay Code and Name, and Append writes those values, but the
// join row is stored with an empty LanguageCode. The association count stays
// 1 instead of 3 on sqlite, mysql, postgres, and sqlserver.
func TestAssociationMany2ManyAppendMap(t *testing.T) {
	user := *GetUser("assoc_m2m_append_map", Config{})
	if err := DB.Create(&user).Error; err != nil {
		t.Fatalf("create user: %v", err)
	}

	// Append single map
	if err := DB.Model(&user).Association("Languages").Append(map[string]interface{}{
		"Code": "am2m_map_1", "Name": "AppendMap1",
	}); err != nil {
		t.Fatalf("append map: %v", err)
	}
	AssertAssociationCount(t, user, "Languages", 1, "after append 1 map")

	// Append more maps individually
	if err := DB.Model(&user).Association("Languages").Append(map[string]interface{}{"Code": "am2m_map_2", "Name": "AppendMap2"}); err != nil {
		t.Fatalf("append map 2: %v", err)
	}
	if err := DB.Model(&user).Association("Languages").Append(map[string]interface{}{"Code": "am2m_map_3", "Name": "AppendMap3"}); err != nil {
		t.Fatalf("append map 3: %v", err)
	}
	AssertAssociationCount(t, user, "Languages", 3, "after append 3 maps total")

	// Verify codes exist
	var langs []Language
	if err := DB.Model(&user).Association("Languages").Find(&langs); err != nil {
		t.Fatalf("find languages: %v", err)
	}
	codeSet := map[string]bool{}
	for _, l := range langs {
		codeSet[l.Code] = true
	}
	for _, c := range []string{"am2m_map_1", "am2m_map_2", "am2m_map_3"} {
		if !codeSet[c] {
			t.Fatalf("expected language code %s present", c)
		}
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
