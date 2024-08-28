package main

import (
	"context"
	"testing"

	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	// Create a wrapping transaction

	err := Transaction(context.Background(), DB, func(ctx context.Context, tx *gorm.DB) error {
		// Within this, create a transaction that retrieves a user, and then
		// creates another transaction that retrieves the account, but errors,
		// and bubble that error up.
		_ = Transaction(ctx, tx, func(ctx context.Context, tx *gorm.DB) error {
			user := User{Name: "jinzhu"}
			var account Account
			err := DB.Create(&user).Error
			if err != nil {
				return err
			}

			// Since we propagate the error, we'll now also try to
			// rollback this nested transaction, using the same
			// savepoint ID, since the `Transaction` helper
			// creates a single closure that always has the
			// same address.
			return Transaction(ctx, tx, func(ctx context.Context, tx *gorm.DB) error {
				// We haven't created an account, so we return an error, which does
				// a rollback of the inner transaction using a savepoint.
				err := tx.First(&account, 1).Error
				if err != nil {
					return err
				}

				return nil
			})
		})
		// We discard the inner transaction error, which allows us to commit this outer
		// transaction (which does nothing), even if the inner fails.
		return nil
	})

	// Since we have rolled back the inner transaction, we expect that
	// no user was created, since we should have rolled that back.
	var user User
	err = DB.First(&user, "name = ?", "jinzhu").Error
	if err == nil {
		t.Errorf("User was created, despite erroring inside the transaction")
	}
}
