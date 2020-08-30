package main

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/playground/pkg/model/migrator"
)

func main() {
	config := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{TablePrefix: "t_", SingularTable: true},
		Logger:         logger.Default.LogMode(logger.Info),
		PrepareStmt:    false,
	}

	url := "postgres://gorm:gorm@postgres:9920/gorm?sslmode=disable"
	var err error
	var db *gorm.DB
	for intento := 1; intento < 5; intento++ {
		db, err = gorm.Open(postgres.Open(url), config)
		if err != nil {
			logrus.Errorf("connection failed: %s (%s)", err.Error(), url)
			logrus.Infof("Retrying in %d seconds", intento)
			time.Sleep(time.Duration(intento) * time.Duration(3) * time.Second)
		} else {
			break
		}
	}
	// termina la aplicación si no fue posible conectar con la base de
	// datos
	if err != nil {
		logrus.Panic("cannot connect with the database")
	}

	m := migrator.New(db, append(v42.Migrations, v43.Migrations...))
	err = m.MigrateAll()
	if err != nil {
		logrus.Panicf("cannot migrate: %v", err)
	}
}
