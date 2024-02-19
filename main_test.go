package main

import (
	"context"
	"errors"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

// Use this command to test:
// GORM_DIALECT=postgres go test

func TestGORM(t *testing.T) {
	user := User{Active: true, Name: "jinzhu"}

	session := DB.Session(&gorm.Session{})
	for i := 0; i < 8000; i++ {
		user.ID = uint(i + 1)
		session.WithContext(context.Background()).Create(&user)
	}

	err := DB.Transaction(func(tx *gorm.DB) error {

		g, ctx := errgroup.WithContext(context.Background())

		for i := 0; i < 8000; i++ {
			i := i

			g.Go(
				func() error {

					if i == 6000 {
						return errors.New("fake error")
					}

					var userName string

					result := session.
						WithContext(ctx).
						Table("users").
						Where("id = ?", i).
						Pluck("name", &userName)

					// When the driver returns `conn busy` and the context is canceled, Pluck panics:
					// panic: runtime error: index out of range [0] with length 0
					// github.com/jackc/pgx/v4/stdlib.(*Rows).Next(0x140000a2320, {0x10549fb80, 0x0, 0x104896cdc?})
					//        external/com_github_jackc_pgx_v4/stdlib/sql.go:767 +0x14d0
					// database/sql.(*Rows).nextLocked(0x140004a4d00)
					//        GOROOT/src/database/sql/sql.go:2974 +0x160
					// database/sql.(*Rows).Next.func1()
					//        GOROOT/src/database/sql/sql.go:2952 +0x30
					// database/sql.withLock({0x104ea0a58, 0x140004a4d30}, 0x140005692a8)
					//        GOROOT/src/database/sql/sql.go:3405 +0x7c
					// database/sql.(*Rows).Next(0x140004a4d00)
					//        GOROOT/src/database/sql/sql.go:2951 +0x64
					// gorm.io/gorm.Scan({0x104ea5890, 0x140004a4d00}, 0x140003b0540, 0x0)
					//        external/io_gorm_gorm/scan.go:159 +0x15c8
					// gorm.io/gorm/callbacks.Query(0x140003b0540)
					//        external/io_gorm_gorm/callbacks/query.go:26 +0xec
					// gorm.io/gorm.(*processor).Execute(0x14000447130, 0x1400001aab0?)
					//        external/io_gorm_gorm/callbacks.go:130 +0x3d0
					// gorm.io/gorm.(*DB).Pluck(0x140007a2180?, {0x104b663b1, 0x2}, {0x104d3a3c0?, 0x140006228c0})
					//        external/io_gorm_gorm/finisher_api.go:542 +0x24c

					if result.Error != nil {
						t.Errorf("Failed during read, got error: %v", result.Error)
						return result.Error
					}

					result = session.Table("users").
						WithContext(ctx).
						Where("active = ?", true).
						Update("name", "new_name")
					if result.Error != nil {
						t.Errorf("Failed during update, got error: %v", result.Error)
						return result.Error
					}

					return nil

				})
		}

		if err := g.Wait(); err != nil {
			t.Errorf("Failed, got error: %v", err)
		}

		return nil
	})

	if err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}
