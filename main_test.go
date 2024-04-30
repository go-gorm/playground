package main

import (
	"testing"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	setup(t)

	tests := []struct {
		name              string
		inputQuery        string
		inputArgs         []any
		wantLen           int
		wantAccountNumber string
	}{
		// These all fail
		{
			name:              "single arrow true bool arg",
			inputQuery:        "config->'$.enabled' = ?",
			inputArgs:         []any{true},
			wantLen:           1,
			wantAccountNumber: "AC12345",
		},
		{
			name:              "json_extract true bool arg",
			inputQuery:        "json_extract(config, '$.enabled') = ?",
			inputArgs:         []any{true},
			wantLen:           1,
			wantAccountNumber: "AC12345",
		},
		{
			name:              "single arrow false bool arg",
			inputQuery:        "config->'$.enabled' = ?",
			inputArgs:         []any{false},
			wantLen:           1,
			wantAccountNumber: "AC67890",
		},
		{
			name:              "json_extract false bool arg",
			inputQuery:        "json_extract(config, '$.enabled') = ?",
			inputArgs:         []any{false},
			wantLen:           1,
			wantAccountNumber: "AC67890",
		},

		// These work
		{
			name:              "single arrow no args with true bool",
			inputQuery:        "config->'$.enabled' = true",
			wantLen:           1,
			wantAccountNumber: "AC12345",
		},
		{
			name:              "single arrow no args with false bool",
			inputQuery:        "config->'$.enabled' = false",
			wantLen:           1,
			wantAccountNumber: "AC67890",
		},
		{
			name:              "double arrow with true string arg",
			inputQuery:        "config->>'$.enabled' = ?",
			inputArgs:         []any{"true"},
			wantLen:           1,
			wantAccountNumber: "AC12345",
		},
		{
			name:              "double arrow with false string arg",
			inputQuery:        "config->>'$.enabled' = ?",
			inputArgs:         []any{"false"},
			wantLen:           1,
			wantAccountNumber: "AC67890",
		},
		{
			name:              "double arrow no args with string",
			inputQuery:        "config->>'$.enabled' = 'true'",
			wantLen:           1,
			wantAccountNumber: "AC12345",
		},
		{
			name:              "json_extract no args with true bool",
			inputQuery:        "json_extract(config, '$.enabled') = true",
			wantLen:           1,
			wantAccountNumber: "AC12345",
		},
		{
			name:              "json_extract no args with false bool",
			inputQuery:        "json_extract(config, '$.enabled') = false",
			wantLen:           1,
			wantAccountNumber: "AC67890",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var results []Account
			if err := DB.Where(tc.inputQuery, tc.inputArgs...).Find(&results).Error; err != nil {
				t.Fatalf("Failed, got error: %v", err)
			}

			if len(results) != tc.wantLen {
				t.Fatalf("Failed, len(results) want=%d got=%d", tc.wantLen, len(results))
			}

			if results[0].Number != tc.wantAccountNumber {
				t.Fatalf("Failed, account.Number want=%s got=%s", tc.wantAccountNumber, results[0].Number)
			}
		})
	}
}

func setup(t *testing.T) {
	err := DB.Exec(`TRUNCATE TABLE accounts`).Error
	if err != nil {
		t.Fatalf("unable to truncate table: %v", err)
	}

	accounts := []Account{
		{
			Number: "AC12345",
			Config: AccountConfig{
				Enabled: true,
				Foo:     "abc",
				Bar:     123,
			},
		},
		{
			Number: "AC67890",
			Config: AccountConfig{
				Enabled: false,
				Foo:     "def",
				Bar:     456,
			},
		},
	}
	DB.Create(&accounts)
}
