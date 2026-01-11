all:
	go mod tidy
	go vet
	staticcheck
	go build -o gorm-bug

clean:
	rm -f gorm-bug
