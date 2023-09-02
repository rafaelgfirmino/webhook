docker-build:
	@docker build -t webhook:latest .

up-build:
	@docker-compose up -d --build

up:
	@docker-compose up -d

down:
	@docker-compose down

test:
	@go test -v -cover -short ./...

godoc:
	@godoc -http=:6060 | echo "godoc is running on http://localhost:6060"



.PHONY: up-build up down test godoc