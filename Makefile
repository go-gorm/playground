.PHONY: benchprof
benchprof:
	cp go.mod-upstream go.mod
	go mod tidy
	# this just initializes the DB
	GORM_DIALECT=postgres go test -benchmem -run=^$$ -bench '^BenchmarkFetchAll$$' gorm.io/playground


	GORM_DIALECT=postgres go test -benchmem -run=^$$ -cpuprofile cpu-upstream.prof -memprofile mem-upstream.prof -bench '^BenchmarkFetchAll' gorm.io/playground | tee log-upstream
	go tool pprof -png cpu-upstream.prof > cpu-upstream.png
	go tool pprof -png mem-upstream.prof > memory-upstream.png

	cp go.mod-patched go.mod
	go mod tidy

	GORM_DIALECT=postgres go test -benchmem -run=^$$ -cpuprofile cpu-fixed.prof -memprofile mem-fixed.prof -bench '^BenchmarkFetchAll' gorm.io/playground | tee log-fixed
	go tool pprof -png cpu-fixed.prof > cpu-fixed.png
	go tool pprof -png mem-fixed.prof > memory-fixed.png

	cat log-upstream | grep '^BenchmarkFetchAll' | sed 's/BenchmarkFetchAll/BenchmarkFetchAll-v1.25.6/' > finallog
	cat log-fixed | grep '^BenchmarkFetchAll' |    sed 's/BenchmarkFetchAll/BenchmarkFetchAll-fix    /' >> finallog
	rm -f log-upstream log-fixed
	cat finallog

