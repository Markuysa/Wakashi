prod: logs metrics
	mkdir -p ./app/internal/logs/data
	go run ./app/cmd/main/main.go 2>&1 | tee ./app/internal/logs/data/log.txt

.PHONY: logs
logs:
	mkdir -p ./app/internal/logs/data/
	touch ./app/internal/logs/data/log.txt
	touch ./app/internal/logs/data/offsets.yaml
	chmod -R 777 ./app/internal/logs/data
	cd ./app/internal/logs && docker compose up

.PHONY: metrics
metrics:
	mkdir -p ./app/internal/metrics/data
	chmod -R 777 ./app/internal/metrics/data
	cd ./app/internal/metrics/data && docker compose up
