prod:
	mkdir -p app/internal/logs/data
	go run app/cmd/main/main.go 2>&1 | tee app/internal/logs/data/log.txt