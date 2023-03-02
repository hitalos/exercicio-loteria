build:
	@go build -o main main.go

build-with-prof:
	@go build -pgo cpu.prof -o main-optimized main.go

view-prof:
	@go tool pprof -http localhost:8000 ./main cpu.prof

compare: build
	@echo "Default build:"
	@bash -c 'time ./main > /dev/null'

	@echo "Optimized build:"
	@./main -profile > /dev/null
	@make build-with-prof
	@bash -c 'time ./main-optimized > /dev/null'
