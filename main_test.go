package main

import (
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {

	DB.Set("gorm:auto_preload", true)
	DB.AutoMigrate(&Organisation{})
	if _, err := SetupOrganisation("name", DB); err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

type Organisation struct {
	gorm.Model
	OrgName      string
	Forward      bool
	InfluxData   *Data `gorm:"embedded"`
	InfluxDataID uint
	Grafana      *OrgSettings `gorm:"embedded"`
	GrafanaID    uint
}

type OrgSettings struct {
	GrafanaName      string
	ID               uint `gorm:"primarykey"`
	GrafanaID        string
	GrafanaAuthToken string
	GrafanaUsers     []User `gorm:"embedded"`
}

// User holds user data
type GrafUser struct {
	Name          string
	ID            uint
	Email         string
	Password      string
	OrgSettingsID uint
}

type Data struct {
	gorm.Model
	//OrganisationID uint   // foreign key for gorm
	Org      string // Organisation as defined by influxdb
	OrgID    string // Influxdb's organisation id
	Bucket   string
	BucketID string
}

// SetupOrganisation sets up a new organisation
func SetupOrganisation(orgName string, db *gorm.DB) (orgModel *Organisation, err error) {
	// Create influx organisation
	inflData := &Data{
		Org:      "coca-cola",
		OrgID:    "aabb12341234",
		Bucket:   "cola-bucket",
		BucketID: "ccdd56785678",
	}

	org := &OrgSettings{
		GrafanaName:  orgName,
		ID:           123,
		GrafanaID:    "123",
		GrafanaUsers: []GrafUser{},
	}

	// Create a new auth token
	//token, err := p.Grafana.CreateAPIToken(org)
	org.GrafanaAuthToken = "super-secret-token"

	// Create new admin user and add him to the organisation
	//admin, err := p.Grafana.CreateUser(org, "Admin", "admin")
	admin := &GrafUser{
		Name:  "Jesse-admin",
		Email: "jesse.geens@gmail.com",
	}
	org.GrafanaUsers = append(org.GrafanaUsers, *(admin))

	// Create new viewer user and add him to the organisation
	//viewer, err := p.Grafana.CreateUser(org, "Viewer", "viewer")
	viewer := &GrafUser{
		Name:  "Jesse-viewer",
		Email: "jesse.geens.viewer@gmail.com",
	}
	org.GrafanaUsers = append(org.GrafanaUsers, *(viewer))

	// Save details
	resOrg := &Organisation{
		OrgName:    orgName,
		InfluxData: inflData,
		Grafana:    org,
	}

	result := db.Create(resOrg) // --> This is where the bug happens!
	return resOrg, result.Error
}
