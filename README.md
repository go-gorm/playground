# GORM Playground

GORM Playground can be used to play GORM and reports issues, if you encounter a bug in GORM, please report it at [https://github.com/go-gorm/gorm/issues](https://github.com/go-gorm/gorm/issues) with the Playground Pull Request's link

[![test status](https://github.com/go-gorm/playground/workflows/tests/badge.svg?branch=master "test status")](https://github.com/go-gorm/playground/actions)

### Quick Start

##### 1. [Fork this repo](https://docs.github.com/en/free-pro-team@latest/github/getting-started-with-github/fork-a-repo)

##### 2. [Clone the forked repo to your local](https://docs.github.com/en/free-pro-team@latest/github/creating-cloning-and-archiving-repositories/cloning-a-repository)

##### 3. Install docker-compose

https://docs.docker.com/compose/install/

##### 4. Run tests with lastest GORM and all drivers

```bash
# run tests (retry if failed due to preparing the testing database)
./test.sh

# Run tests with cached GORM and latest drivers
GORM_ENABLE_CACHE=true ./test.sh

# Before running go test directly, prepare the local GORM checkout and dependencies
git clone --depth 1 https://github.com/go-gorm/gorm.git gorm
go mod download

# Run tests with specfied database
GORM_DIALECT=mysql go test
```

The local `gorm` checkout is required by the replacement in `go.mod`; skip the
clone command if it already exists. Keep `go.sum` under version control so direct
Go commands can verify dependencies without missing-checksum errors.

##### 5. Modify tests and make it fail

##### 6. [Create Playground Pull Request](https://docs.github.com/en/free-pro-team@latest/github/collaborating-with-issues-and-pull-requests/creating-a-pull-request) and [Create a GORM issue](https://github.com/go-gorm/gorm/issues/new?template=bug_report.md) with the link

### Advanced Usage

We are using the following configuration run your code (GORM's latest master branch, latest database drivers: sqlite, mysql, postgres, sqlserver), you could change the configuration in file [main_test.go](https://github.com/go-gorm/playground/edit/master/main_test.go)

```go
// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver
```

We have prepared some structs with relationships in [https://github.com/go-gorm/playground/blob/master/models.go](https://github.com/go-gorm/playground/blob/master/models.go) that you can use for your tests

### GORM Gen regression reproduction

The regression test is in [gen_issue_repro_test.go](gen_issue_repro_test.go).
It uses the playground's PostgreSQL connection, creates and drops the
`memberships` table itself, and skips other dialects. Use a disposable test
database: the playground also recreates its standard model tables at startup.

Run it from PowerShell against a PostgreSQL test database:

```powershell
docker run --rm --name gorm-gen-repro -e POSTGRES_PASSWORD=pw -p 55432:5432 -d postgres:17
# Wait until this reports "accepting connections".
docker exec gorm-gen-repro pg_isready -U postgres
$env:GORM_DIALECT = "postgres"
$env:GORM_DSN = "host=localhost port=55432 user=postgres password=pw dbname=postgres sslmode=disable"
go test -mod=readonly -run '^TestUpdateLeaksPrimaryKeyIntoSharedModel$' -count=1 -v .
docker stop gorm-gen-repro
```

By default, this uses the local GORM checkout configured in `go.mod`. To reproduce
with the released versions listed below instead, first run:

```powershell
go mod edit '-dropreplace=gorm.io/gorm'
go get gorm.io/gorm@v1.31.2 gorm.io/driver/postgres@v1.6.3 gorm.io/plugin/dbresolver@v1.6.2 gorm.io/gen@v0.3.29
go mod tidy
```

Repeat the `go test` command after selecting each Gen version with
`go get gorm.io/gen@v0.3.27`, `go get gorm.io/gen@v0.3.28`, or
`go get gorm.io/gen@v0.3.29`, followed by `go mod tidy`.
Do not use `test.sh` for this version comparison: it runs `go get -u -t ./...`,
which upgrades dependencies. The explicit `dbresolver` version is needed because
Gen's older transitive dependency does not compile with this GORM version.

#### Issue draft for go-gorm/gen

**Title:** Update() since v0.3.28 mutates the shared UseModel instance, leaking a primary-key condition into later subqueries

**Description**

Starting with `gorm.io/gen v0.3.28`, `DO.Update(column, value)` writes the
updated value into the model instance registered via `UseModel(&Membership{})`.
Generated Query objects reuse this model instance. After updating a primary-key
field, a later subquery built from the same shared DO gains an unintended
primary-key condition.

**Reproduction**

See [gen_issue_repro_test.go](gen_issue_repro_test.go) in this playground.
The test uses PostgreSQL and a composite primary key:
`memberships(group_id text, user_id text, primary key(group_id, user_id))`.
It builds a subquery, executes an unrelated update from group `X` to `Y` through
the same shared DO, then builds another subquery for group `B`.
The table can be empty; the model is mutated even when the update affects zero rows.

**Actual**

```sql
-- Before Update; parameters: [A]
SELECT * FROM "users" WHERE id IN (SELECT "memberships"."user_id" FROM "memberships" WHERE "memberships"."group_id" = $1)
-- After Update; parameters: [B Y]
SELECT * FROM "users" WHERE id IN (SELECT "memberships"."user_id" FROM "memberships" WHERE "memberships"."group_id" = $1 AND "memberships"."group_id" = $2)
```

The extra `$2` is the value `Y` from the earlier update. The shared model has
changed from `Membership{}` to `Membership{GroupID: "Y"}`.

**Expected**

The SQL structure stays unchanged after `Update`, with only the requested
group filter as a parameter (`[B]`). The shared model registered through
`UseModel` remains unchanged.

**Verified versions**

| gorm.io/gen | Result |
|---|---|
| v0.3.27 | Passes |
| v0.3.28 | Fails: extra primary-key condition and mutated model |
| v0.3.29 | Fails: extra primary-key condition and mutated model |

All three versions were tested against PostgreSQL 17 with
`gorm.io/gorm v1.31.2`, `gorm.io/driver/postgres v1.6.3`, and
`gorm.io/plugin/dbresolver v1.6.2`, without the local GORM replacement.

**Impact**

When a generated Query object is shared for the lifetime of the process, one
update can contaminate later subquery-based reads. Queries for another group
then contain conflicting group filters and return empty results.

## Happy Hacking!
