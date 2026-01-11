# Works correctly with gorm v1.31.0

Note User and Location are both non-nil as expected

```
bigbaby:gorm-preload-bug tadhunt$ grep gorm.io go.mod
	gorm.io/gorm v1.31.0
bigbaby:gorm-preload-bug tadhunt$ make
go mod tidy
go vet
staticcheck
go build -o gorm-bug
bigbaby:gorm-preload-bug tadhunt$ ./gorm-bug

=== Loading shift with Preload ===

Shift ID: shift-001
Shift UserID: 1001
Shift LocationID: 2001
User: Alice Smith (ID: 1001)
Location: Office A (ID: 2001)
```

# Fails with gorm v1.31.1

```
bigbaby:gorm-preload-bug tadhunt$ grep gorm.io go.mod
	gorm.io/gorm v1.31.1
bigbaby:gorm-preload-bug tadhunt$ make
go mod tidy
go vet
staticcheck
go build -o gorm-bug
bigbaby:gorm-preload-bug tadhunt$ ./gorm-bug

=== Loading shift with Preload ===

Shift ID: shift-001
Shift UserID: 1001
Shift LocationID: 2001
ERROR: User is nil - Preload failed!
```
