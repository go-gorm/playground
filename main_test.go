package main

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

import (
	"errors"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

var models = []interface{}{&BaseModel{}, &Site{}, &User{}, &Company{}, &BusinessModel{}, &Directory{},
	&File{}, &TimeSlot{}, &Person{}, &Employee{}, &Subdivision{}, &BranchOffice{}, &ProductGroup{}}

//var models = GetModels()
var testDB *gorm.DB

func InitDB(migrate bool, customConfig bool) {
	dbPath := filepath.Join(os.TempDir(), "gorm.test.db")
	os.Remove(dbPath)
	if customConfig {
		DB, _ = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	} else {
		DB, _ = gorm.Open(sqlite.Open(dbPath), &gorm.Config{SkipDefaultTransaction: false, PrepareStmt: true, DisableForeignKeyConstraintWhenMigrating: true})
	}
	testDB = DB
	if migrate {
		testDB.Migrator().DropTable(models...)
		testDB.AutoMigrate(models...)
	}

}

func TestGORM_AutoMigrate(t *testing.T) { // OK
	InitDB(false, false)
	if err := testDB.AutoMigrate(models...); err != nil {
		t.Errorf("List AutoMigrate Failed, got error: %v", err)
	}
	if err := testDB.Migrator().DropTable(models...); err != nil {
		t.Errorf("DropTable Failed, got error: %v", err)
	}
}

func TestGORM_AutoMigrate2(t *testing.T) { // OK
	InitDB(false, false)
	for _, model := range models {
		fmt.Println("AutoMigrate - " + reflect.ValueOf(model).Elem().Type().Name())
		if err := testDB.AutoMigrate(model); err != nil {
			t.Errorf("AutoMigrate Failed, got error: %v", err)
		}
	}
	for _, model := range models {
		fmt.Println("DropTable " + reflect.ValueOf(model).Elem().Type().Name())
		if err := testDB.Migrator().DropTable(model); err != nil {
			t.Errorf("DropTable Failed, got error: %v", err)
		}
	}
}

func TestGORM_StackOverflowInCreate(t *testing.T) {
	if os.Getenv("SKIP") != "" {
		t.Skip()
	}
	InitDB(true, false)

	admin := &User{PersonName: " Admin", Enabled: true, UserType: "admin", Password: "123"}
	admin.Name = "Admin"
	admin.Description = "Администратор сайта"
	admin.User = admin
	testDB.Create(admin)
}

func TestGORM_StackOverflowInSave(t *testing.T) {
	if os.Getenv("SKIP") != "" {
		t.Skip()
	}
	InitDB(true, false)

	admin := &User{PersonName: " Admin", Enabled: true, UserType: "admin", Password: "123"}
	admin.Name = "Admin"
	admin.Description = "Администратор сайта"
	testDB.Create(admin)
	admin.User = admin
	testDB.Save(admin)
}

func TestGORM_Series(t *testing.T) {
	Series(false, t) // data corruption occurs
	Series(true, t) // in my project data corruption occurs in custom config case
}

func CreateSite(site *Site, admin *User, company *Company) (*Site, *User, *Company) {
	site.Name = strings.Split(site.Host, ".")[0]
	site.Number = site.Number
	if admin == nil {
		admin = &User{PersonName: site.Name + " Admin", Enabled: true, UserType: "admin", Password: "HashAndSalt(site.Name + time.Now().String())"}
		admin.Name = site.Number + "Admin"
		admin.Description = "Администратор сайта - " + site.Host
		//admin.User = admin // lineA
		testDB.Create(admin) // not work 1 if lineA uncommented
		//admin.User = admin // lineB
		//testDB.Save(admin) // not work 2 if lineB uncommented
		testDB.Model(admin).Update("user_id", admin.ID) // ok in lineB - commented
	}

	if company == nil {
		company = &Company{}
		company.Name = site.Name
		company.Number = site.Number
		company.Description = "Адрес сайта - " + site.Host
		company.User = admin // not stored (stored as miracle)
		//company.Company=company // not work
		testDB.Create(company) // miracle = in this line in table "users" (not companies! as I suppose) field company_id becomes = admin.ID in line "company.User = admin" uncommented
		testDB.Model(company).Update("company_id", company.ID)
		testDB.Model(company).Update("user_id", admin.ID)
	}

	//if admin.Company == nil {
	//	admin.Company = company
	//	testDB.Save(admin) // stack overflow
	//}
	if admin.CompanyID == 0 {
		testDB.Model(admin).Update("company_id", company.ID)
	}

	//site.Company = company // stored as supposed in CompanyID
	//site.User = admin // stored as supposed in UserID
	site.CompanyID = company.ID // stored as supposed in CompanyID
	site.UserID = admin.ID
	testDB.Create(site) // work well
	return site, admin, company
}

func Series(customConfig bool, t *testing.T) {
	InitDB(true, customConfig)

	sites := []*Site{}
	users := []*User{}
	companies := []*Company{}

	GtxBotToken := "12495358:AAGuC3QJg"
	GtxChat := "70347"
	CreateSite(&Site{BusinessModel: BusinessModel{Name: "My Chef"}, Host: "mychefeda.ru", Hosts: "edaa.ga;127.1.1.2", Redirects: "edarot.ru;edaa.ga<samon.tk",
		BotToken: "1383:AAEzuBSV2AeA", Chat: "-100173315", AllowLogin: true, AllowRegister: true,
		YandexShopID: "7405", YandexPassword: "test_MCcL1wRWFKHOEFx6bC8"}, nil, nil)
	CreateSite(&Site{BusinessModel: BusinessModel{Name: "asamon"}, Host: "asamon.ru", Hosts: "127.1.1.3",
		BotToken: GtxBotToken, Chat: GtxChat, AllowLogin: true, AllowRegister: false}, nil, nil)
	CreateSite(&Site{BusinessModel: BusinessModel{Name: "1cluga"}, Host: "1clu.ga", Hosts: "127.1.1.4",
		BotToken: GtxBotToken, Chat: GtxChat, AllowLogin: true, AllowRegister: false}, nil, nil)
	CreateSite(&Site{BusinessModel: BusinessModel{Name: "Тестовый сайт"}, Host: "ustas.ml", Hosts: "127.1.1.5",
		BotToken: GtxBotToken, Chat: GtxChat, AllowLogin: true, AllowRegister: false,
		YandexShopID: "7415", YandexPassword: "test_pSFqmi6t0ZgbgJbXkLNoQ"}, nil, nil)
	CreateSite(&Site{BusinessModel: BusinessModel{Name: "Тестовый домашний 1cka.ml"}, Host: "1cka.ml", BotToken: GtxBotToken, Chat: GtxChat,
		AllowLogin: true, AllowRegister: false, YandexShopID: "7354", YandexPassword: "test_ApA9lu8wIkB27E9llkTo"}, nil, nil)

	for i := 0; i < 50; i++ {
		site, user, company := CreateSite(&Site{BusinessModel: BusinessModel{Name: "Метамагазин"}, Host: "metamagazine.ru", Hosts: "127.1.1.1", Redirects: "metamagazin.ru;127.5.5.5",
			RedirectFromRoot: true, BotToken: GtxBotToken, Chat: GtxChat, AllowLogin: true, AllowRegister: false}, nil, nil)

		sites = append(sites, site)
		users = append(users, user)
		companies = append(companies, company)

		CreateSite(&Site{BusinessModel: BusinessModel{Name: "Метамагазин"}, Host: "metamagazine.ru", Hosts: "127.1.1.1", Redirects: "metamagazin.ru;127.5.5.5",
			RedirectFromRoot: true, BotToken: GtxBotToken, Chat: GtxChat, AllowLogin: true, AllowRegister: false}, nil, nil)
	}

	for i := 0; i < 50; i++ {
		site := sites[i]
		user := users[i]
		company := companies[i]

		if user.UserID != user.ID {
			t.Errorf("TestGORM_Series: %v, customConfig: %v", errors.New("user.UserID != user.ID"), customConfig)
		}
		if user.CompanyID != company.ID {
			t.Errorf("TestGORM_Series: %v, customConfig: %v", errors.New("user.CompanyID != company.ID"), customConfig)
		}
		if company.UserID != user.ID {
			t.Errorf("TestGORM_Series: %v, customConfig: %v", errors.New("company.UserID != user.ID"), customConfig)
		}
		if company.CompanyID != company.ID {
			t.Errorf("TestGORM_Series: %v, customConfig: %v", errors.New("company.CompanyID != company.ID"), customConfig)
		}
		if site.CompanyID != company.ID {
			t.Errorf("TestGORM_Series: %v, customConfig: %v", errors.New("site.CompanyID != company.ID"), customConfig)
		}
		if site.UserID != user.ID {
			t.Errorf("TestGORM_Series: %v, customConfig: %v", errors.New("site.UserID != user.ID"), customConfig)
		}
	}

	var q_sites []*Site
	testDB.Find(&q_sites)

	var q_users []*User
	testDB.Find(&q_users)

	var q_companies []*Company
	testDB.Find(&q_companies)

	if len(q_companies) != 105 {
		t.Errorf("TestGORM_Series: %v", errors.New("len(sites)!=105"))
	}
	if len(q_users) != 105 {
		t.Errorf("TestGORM_Series: %v", errors.New("len(users)!=105"))
	}
	if len(q_companies) != 105 {
		t.Errorf("TestGORM_Series: %v", errors.New("len(companies)!=105"))
	}

}
