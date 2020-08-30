package main

import (
	"os"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/playground/pkg/model/migrator"
	v42 "gorm.io/playground/pkg/model/v42"
	v43 "gorm.io/playground/pkg/model/v43"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: postgres

func TestGORM(t *testing.T) {
	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{TablePrefix: "t_", SingularTable: true},
		Logger:         logger.Default.LogMode(logger.Info),
		PrepareStmt:    false,
	}

	dbDSN := os.Getenv("GORM_DSN")
	var err error
	var db *gorm.DB
	for intento := 1; intento < 5; intento++ {
		db, err = gorm.Open(postgres.Open(dbDSN), config)
		if err != nil {
			logrus.Errorf("connection failed: %s (%s)", err.Error(), dbDSN)
			logrus.Infof("Retrying in %d seconds", intento)
			time.Sleep(time.Duration(intento) * time.Duration(3) * time.Second)
		} else {
			break
		}
	}
	// termina la aplicación si no fue posible conectar con la base de
	// datos
	if err != nil {
		t.Fatalf("cannot connect with database: %v", err)
	}

	m := migrator.New(db, append(v42.Migrations, v43.Migrations...))
	err = m.MigrateAll()
	if err != nil {
		t.Fatal(err)
	}
}
