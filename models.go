package main

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type (
	BaseModel struct {
		//gorm.Model // `json:"-"` in gorm.Model
		ID        uint           `gorm:"primarykey"`
		CreatedAt time.Time      `json:"-"`
		UpdatedAt time.Time      `json:"-"`
		DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

		Company   *Company `display:"none" visibility:"500" Label:"Компания" json:"-"`
		CompanyID uint     `display:"none" json:"-"`
		User      *User    `display:"none" visibility:"500" Label:"Пользователь" json:"-"`
		UserID    uint     `display:"none" json:"-"`
	}

	Site struct {
		BusinessModel    `Model:"Сайт" Plural:"Сайты"`
		Host             string `Label:"Адрес сайта"`
		Hosts            string `Label:"Дополнительные адреса сайта"` // ; - separator
		Redirects        string `Label:"Адреса перенаправления"`      // ; - separator, - - delimiter
		RedirectFromRoot bool   `Label:"Перенаправлять на свой домен"`
		AllowLogin       bool   `Label:"Разрешение входа пользователя"`
		AllowRegister    bool   `Label:"Разрешение регистрации"`

		BotToken string `Label:"ID и токен телеграм канала"`
		Chat     string `Label:"ID чата телеграм канала"`

		YandexShopID   string `Label:"ID Yandex магазина"`
		YandexPassword string `Label:"Пароль Yandex магазина"`

		//Renderer *echoview.ViewEngine `gorm:"-"`
	}

	User struct {
		Directory
		UserType   string    `Label:"Тип"`
		Email      string    `Label:"Почта" validType:"email"`
		Phone      string    `Label:"Телефон"`
		Password   string    `Label:"Пароль"`
		Language   string    `Label:"Язык"`
		Enabled    bool      `Label:"Включен"` // Блокировка
		PersonID   uint      `Label:"Персона"`
		Person     *Person   `Label:"Персона"`
		PersonName string    `Label:"ФИО"`
		EmployeeID uint      `Label:"Работник"`
		Employee   *Employee `Label:"Работник"`
	}
	Company struct { //Organization
		Directory
		CompanyType string
		PersonID    uint    `Label:"Владелец"`
		Person      *Person `Label:"Владелец"`
		TimeSlotID  uint
		TimeSlot    *TimeSlot `Label:"Расписание работы"`
	}
	BusinessModel struct {
		BaseModel
		Number      string `Label:"Номер"`
		Numb        uint   `Label:"Номер" visibility:"300"`
		Name        string `Label:"Наименование" required:"true"`
		Files       []File `gorm:"polymorphic:Owner;" Label:"Файлы" visibility:"900" index:"502" json:",omitempty"`
		Description string `Label:"Описание" index:"501" json:",omitempty"`
	}
	Directory struct {
		BusinessModel
		DirectoryID uint       `Label:"Группа"`
		Directory   *Directory `Label:"Группа" do:"width:100"`
	}
	File struct {
		BaseModel
		OwnerID   int
		OwnerType string
		Path      string
	}
	TimeSlot struct {
		BusinessModel
		Start time.Time
		End   time.Time
	}
	Person struct {
		Directory	`Plural:"Физические лица" Model:"Физическое лицо"`
		Sex      string    `Label:"Пол"`
		Birthday time.Time `Label:"Дата рождение"`
		Address  string    `Label:"Адрес"`
	}

	Employee struct {
		Directory	`Plural:"Работники" Model:"Работник"`
		PersonID      uint    `Label:"Персона"`
		Person        *Person `Label:"Персона"`
		SubdivisionID uint
		Subdivision   *Subdivision
	}
	Subdivision struct {
		Directory		`Plural:"Подразделения" Model:"Подразделение"`
		BranchOfficeID uint
		BranchOffice   *BranchOffice
		TimeSlotID     uint
		TimeSlot       *TimeSlot `Label:"Расписание работы"`
	}
	BranchOffice struct {
		Directory	`Plural:"Филиалы компании" Model:"Филиал компании"`
		TimeSlotID uint
		TimeSlot   *TimeSlot `Label:"Расписание работы"`
	}
	ProductGroup struct {
		Directory	`Plural:"Группа номенклатуры" Model:"Группы номенклатуры" Category:"Прочее"`
	}
)

//region unused

// User has one `Account` (has one), many `Pets` (has many) and `Toys` (has many - polymorphic)
// He works in a Company (belongs to), he has a Manager (belongs to - single-table), and also managed a Team (has many - single-table)
// He speaks many languages (many to many) and has many friends (many to many - single-table)
// His pet also has one Toy (has one - polymorphic)
type User_PG struct {
	gorm.Model
	Name      string
	Age       uint
	Birthday  *time.Time
	Account   Account
	Pets      []*Pet
	Toys      []Toy `gorm:"polymorphic:Owner"`
	CompanyID *int
	Company   Company
	ManagerID *uint
	Manager   *User
	Team      []User     `gorm:"foreignkey:ManagerID"`
	Languages []Language `gorm:"many2many:UserSpeak"`
	Friends   []*User    `gorm:"many2many:user_friends"`
	Active    bool
}

type Account struct {
	gorm.Model
	UserID sql.NullInt64
	Number string
}

type Pet struct {
	gorm.Model
	UserID *uint
	Name   string
	Toy    Toy `gorm:"polymorphic:Owner;"`
}

type Toy struct {
	gorm.Model
	Name      string
	OwnerID   string
	OwnerType string
}

type Company_PG struct {
	ID   int
	Name string
}

type Language struct {
	Code string `gorm:"primarykey"`
	Name string
}

//endregion