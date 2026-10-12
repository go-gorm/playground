# PostgreSQL batch map insert reproduction

This standalone module uses only GORM, its PostgreSQL driver, a small model,
and synthetic data. It does not use the playground models, mocks, or an explicit
`RETURNING` clause. **Use a disposable database: the program drops `users`.**

```sh
cd repro
go mod download
export GORM_DSN='host=127.0.0.1 port=55439 user=gorm dbname=gorm_repro sslmode=disable'
go run .
```

For a temporary native PostgreSQL server (no containers), from any directory:

```sh
PG_REPRO_DIR=$(mktemp -d)
initdb -D "$PG_REPRO_DIR/data" -U gorm -A trust --no-locale -E UTF8
pg_ctl -D "$PG_REPRO_DIR/data" -l "$PG_REPRO_DIR/server.log" \
  -o "-h 127.0.0.1 -p 55439 -k $PG_REPRO_DIR" -w start
createdb -h 127.0.0.1 -p 55439 -U gorm gorm_repro
# Run the module above, then stop this temporary server:
pg_ctl -D "$PG_REPRO_DIR/data" -m fast -w stop
```

Expected: one persisted row, `name=alice`, `age=18`, no panic.
Observed with the pinned release:

```text
INSERT INTO "users" ("age","name") VALUES (18,'alice') RETURNING "id"
panic: reflect: slice index out of range
gorm.io/gorm.(*DB).CreateInBatches.func1
    gorm/finisher_api.go:61
```

To test the master commit checked on 2026-10-12 without changing the pinned module:

```sh
cp go.mod master.mod
cp go.sum master.sum
go mod edit -modfile=master.mod \
  -require=gorm.io/gorm@v1.31.3-0.20260914162115-b3d3bf219f02
go run -mod=mod -modfile=master.mod .
```

## Verified environment and observations

- Go 1.26.3, macOS arm64, native PostgreSQL 17.10.
- Latest GORM release confirmed: v1.31.2; driver: v1.6.3.
- Master: `b3d3bf219f0283f8e2e985bac509cb643170f729`.
- Both release and master panic with this single-row reproduction.
- `main_test.go` in the playground root covers values and pointers, direct
  `CreateInBatches` and `CreateBatchSize`, one row/batch size 500, and five rows/
  batch size 2. `TestBatchControls` verifies struct inserts and table-only map
  inserts without automatic `RETURNING`.
- SQLite also reproduces the single-row panic. PostgreSQL and SQLite pass the
  same regression cases with the proposed production fix.
- With one batch and default transaction settings, the row is committed before
  the panic occurs in the outer copy-back loop. Do not infer rollback from panic.
- With five rows and batch size 2, returned ID-only maps overwrite the next
  batch in the shared backing array. The next insert can fail with a duplicate
  key error, and the default outer transaction rolls back all rows. This is an
  observation on the current release, not a claim about historical v1.22.5 bugs.

The new panic is related to #7768's addressable batch destination and copy-back
loop. Testing its parent `33229296228898b311e868090f30eed169f1ae7a` gives an
`unsupported Scan` error for this reproduction rather than this panic; this
report does not claim batched maps were fully functional before #7768.

To run the root playground tests without its container-oriented script:

```sh
# From the playground root:
git clone https://github.com/go-gorm/gorm.git gorm
go test -mod=mod -run '^TestBatchControls$' -v
GORM_DIALECT=postgres GORM_DSN="$GORM_DSN" \
  go test -mod=mod -run '^TestGORM$/pointer=false/configured=false/count=1$' -v
```

The standalone module pins driver v1.6.3; the root playground uses upstream driver declarations, and its normal
`test.sh` refreshes them. An unused `gorm.io/cmd/gorm` indirect requirement was
removed because its unavailable vanity import blocked CI before tests could run.
