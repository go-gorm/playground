package migrator

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// MigrateFunc tipo que describe la funcion de migracion
type MigrateFunc func(tx *gorm.DB, m gorm.Migrator) error

// RollbackFunc tipo que describe la funcion de retroceso
type RollbackFunc func(tx *gorm.DB, m gorm.Migrator) error

// Migration es la tabla de migraciones
type Migration struct {
	ID string `gorm:"primaryKey"`
}

// SingleMigration es una unica migracion
type SingleMigration struct {
	ID       string
	Migrate  MigrateFunc
	Rollback RollbackFunc
}

func (s *SingleMigration) isMigrated(db *gorm.DB) (bool, error) {
	var migration Migration
	err := db.Find(&migration, "id = ?", s.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}

	return migration.ID == s.ID, err
}

func (s *SingleMigration) isMe(id string) bool {
	return s.ID == id
}

// New retorna un nuevo Migrator
func New(db *gorm.DB, migrations []*SingleMigration) *Migrator {
	return &Migrator{
		db:         db,
		migrations: migrations,
	}
}

// Migrator contiene las migraciones
type Migrator struct {
	db         *gorm.DB
	migrations []*SingleMigration
}

// beforeMigrate migra la tabla de migraciones de go-gormmigrate/gormmigrate,
// esto es importante porque gorm tiene varias opciones de configuracion para
// los nombres de tablas y ya no usa TableName().
func (m *Migrator) beforeMigrate() error {
	// migra la tabla de migraciones
	if m.db.Migrator().HasTable("migrations") {
		return m.db.Migrator().RenameTable("migrations", &Migration{})
	}
	// crea la tabla de migraciones si no existia antes
	if !m.db.Migrator().HasTable(&Migration{}) {
		return m.db.AutoMigrate(&Migration{})
	}
	return nil
}

// MigrateTo realiza la migracion en orden hasta determinado ID (inclusivo)
func (m *Migrator) MigrateTo(id string) error {
	ckErr := m.beforeMigrate()
	if ckErr != nil {
		return ckErr
	}

	for index, single := range m.migrations {
		migrated, err := single.isMigrated(m.db)
		if err != nil {
			return fmt.Errorf("Migrator: no se pudo migrar ID %s (%d): %v", single.ID, index, err)
		}
		if migrated {
			logrus.WithField("id", single.ID).Info("migracion ya aplicada")
			continue
		}
		// realiza la migracion
		tx := m.db.Begin()
		gormMigrator := tx.Migrator()

		err = single.Migrate(tx, gormMigrator)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("Migrator: hubo un error al aplicar la migracion ID %s: %v", single.ID, err)
		}
		// registra la migracion
		err = tx.Create(&Migration{ID: single.ID}).Error
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("Migrator: hubo un error al registrar la migracion ID %s: %v", single.ID, err)
		}
		tx.Commit()

		if single.isMe(id) {
			break
		}
	}
	return nil
}

// MigrateAll aplica todas las migraciones en orden
func (m *Migrator) MigrateAll() error {
	single := m.migrations[len(m.migrations)-1]

	return m.MigrateTo(single.ID)
}
