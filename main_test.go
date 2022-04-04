package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

type CompanyEmployeeInfo struct {
	Company  *Company  `gorm:"embedded"`
	Employee *Employee `gorm:"embedded"`
}

func TestGORM(t *testing.T) {
	RunMigrations()

	company1 := Company{Name: "Apple"}
	company2 := Company{Name: "Meta"}
	company3 := Company{Name: "Amazon"}
	company4 := Company{Name: "Tesla"}

	DB.Create(&company1)
	DB.Create(&company2)
	DB.Create(&company3)
	DB.Create(&company4)

	employee1 := Employee{Name: "Harold"}
	employee2 := Employee{Name: "John"}
	employee3 := Employee{Name: "Tom"}
	employee4 := Employee{Name: "Harry"}

	DB.Create(&employee1)
	DB.Create(&employee2)
	DB.Create(&employee3)
	DB.Create(&employee4)

	DB.Create(&CompanyEmployeeJunction{CompanyID: company1.ID, EmployeeID: employee3.ID})
	DB.Create(&CompanyEmployeeJunction{CompanyID: company1.ID, EmployeeID: employee4.ID})
	DB.Create(&CompanyEmployeeJunction{CompanyID: company2.ID, EmployeeID: employee1.ID})
	DB.Create(&CompanyEmployeeJunction{CompanyID: company2.ID, EmployeeID: employee3.ID})

	var companyEmployeeInfoList []*CompanyEmployeeInfo
	if err := DB.Select("companies.*, employees.*").
		Table("companies").
		Joins("inner join company_employee_junctions on companies.id = company_employee_junctions.company_id").
		Joins("inner join employees on company_employee_junctions.employee_id = employees.id").
		Find(&companyEmployeeInfoList).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	companyMatched := false
	employeeMatched := false

	for _, info := range companyEmployeeInfoList {
		if info.Company == nil {
			t.Fatalf("Failed, expected not nil, got company nil")
		}
		if info.Employee == nil {
			t.Fatalf("Failed, expected not nil, got employee nil")
		}
		if info.Company.ID == company1.ID {
			companyMatched = true
			if info.Company.Name != company1.Name {
				t.Fatalf("Failed, expected %v, got %v", company1.Name, info.Company.Name)
			}
		}
		if info.Employee.ID == employee1.ID {
			employeeMatched = true
			if info.Employee.Name != employee1.Name {
				t.Fatalf("Failed, expected %v, got %v", company1.Name, info.Company.Name)
			}
		}
	}

	if !companyMatched {
		t.Fatalf("Failed, no company matched")
	}
	if !employeeMatched {
		t.Fatalf("Failed, no employee matched")
	}
}
