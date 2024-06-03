package main

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sync"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

type dbConnections struct {
	aDB *gorm.DB
	tDB map[string]*gorm.DB
}

var db dbConnections

// TestSelectResultCaching
// Before running this test execute this query
// DELETE FROM departments
// WHERE tenant_id = ':1'
// AND department_id = ':2'
func TestSelectResultCaching(t *testing.T) {
	initDB()
	if err := initConstraints(); err != nil {
		t.Errorf(err.Error())
	}
	newD := Department{
		TenantID:       "T1",
		DepartmentID:   uuid.New().String(),
		DepartmentName: "FANCY DEPARTMENT NAME",
		Employees: []Employee{
			{EmployeeID: "001", EmployeeName: "Jinzhu"},
			{EmployeeID: "002", EmployeeName: "Muchen"},
			{EmployeeID: "003", EmployeeName: "Mingze"},
			{EmployeeID: "004", EmployeeName: "Yichen"},
			{EmployeeID: "005", EmployeeName: "Muyang"},
		},
	}
	tDB, err := getDB(newD.TenantID)
	if err != nil {
		t.Errorf(err.Error())
		return
	}
	dbDept, err := getDept(tDB, newD.TenantID, newD.DepartmentID)
	if err != nil {
		t.Errorf("Get department before create error: %v", err)
		return
	}
	if len(dbDept.DepartmentID) > 0 {
		t.Errorf("Get department before create but it already exists: " + dbDept.DepartmentID)
		return
	}
	dbDept, err = createDept(tDB, newD)
	if err != nil {
		t.Errorf("Create department error: %v", err)
		return
	}
	dbDept, err = getDept(tDB, dbDept.TenantID, dbDept.DepartmentID)
	if err != nil {
		t.Errorf("Get department after create error: %v", err)
		return
	}
}

func getDept(tDB *gorm.DB, tenantID string, departmentID string) (*Department, error) {
	var departments []Department
	output := tDB.Model(Department{}).
		Preload("Employees").
		Where("tenant_id = ?", tenantID).
		Where("department_id = ?", departmentID).
		Find(&departments)
	if output.Error != nil || output.RowsAffected == 0 {
		return &Department{}, output.Error
	}
	if output.RowsAffected == 1 {
		return &departments[0], nil
	}
	return &Department{}, errors.New(fmt.Sprintf("too many rows affected %v, check indexes and keys", output.RowsAffected))
}

func createDept(tDB *gorm.DB, newD Department) (*Department, error) {
	tx := tDB.Begin()
	defer tx.Rollback()
	output := tx.
		Model(&Department{}).
		Create(newD)
	if output.Error != nil || output.RowsAffected == 0 {
		return &Department{}, output.Error
	}
	audit := Audit{
		AuditID:   uuid.New().String(),
		AuditDesc: "created department " + newD.DepartmentID,
	}
	tx.
		Model(Audit{}).
		Create(&audit)
	if output.Error != nil || output.RowsAffected == 0 {
		return &Department{}, output.Error
	}
	tx.Commit()
	return &newD, output.Error
}

func initDB() {
	db.aDB = DB
	db.tDB = make(map[string]*gorm.DB)
}

const grantSQL = "grant delete, insert, references, select, trigger, truncate, update on %s to %s;"
const policySQL = "CREATE POLICY %s ON %s using (tenant_id = current_setting('myapp.current_tenant_id')) WITH CHECK (tenant_id = current_setting('myapp.current_tenant_id'));"
const rlsSQL = "ALTER TABLE %s ENABLE ROW LEVEL SECURITY;"

func initConstraints() error {
	aDB, err := getDB("")
	if err != nil {
		return err
	}

	nonAdminUser := os.Getenv("NON_ADMIN_USER")

	err = constrain(aDB, nonAdminUser, Employee{}.TableName())
	if err != nil {
		return err
	}
	err = constrain(aDB, nonAdminUser, Department{}.TableName())
	if err != nil {
		return err
	}
	return nil
}

func constrain(aDB *gorm.DB, nonAdminUser string, tName string) error {
	//1. Grant
	grantQuery := fmt.Sprintf(grantSQL, tName, nonAdminUser)
	err := aDB.Exec(grantQuery).Error
	if err != nil {
		return err
	}
	//2. policy
	policyQuery := fmt.Sprintf(policySQL, tName+"_rls_policy", tName)
	err = aDB.Exec(policyQuery).Error
	if err != nil {
		return err
	}
	//3. enable RLS
	rlsQuery := fmt.Sprintf(rlsSQL, tName)
	err = aDB.Exec(rlsQuery).Error
	if err != nil {
		return err
	}
	return nil
}

var lock sync.Mutex

func getDB(tenantID string) (*gorm.DB, error) {
	if len(tenantID) == 0 {
		return db.aDB, nil
	}
	tDB, ok := db.tDB[tenantID]
	if ok {
		return tDB, nil
	}
	lock.Lock()
	defer lock.Unlock()
	tDB, ok = db.tDB[tenantID]
	if ok {
		return tDB, nil
	}
	var err error
	tDB, err = gorm.Open(postgres.Open(os.Getenv("TDSN")), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	setQuery := fmt.Sprintf("SET myapp.current_tenant_id = '%s'", tenantID)
	dbConn := tDB.Exec(setQuery)
	db.tDB[tenantID] = dbConn

	tDB = db.tDB[tenantID]
	return tDB, nil
}
