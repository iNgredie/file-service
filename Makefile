export

DB_MIGRATE_URL = postgres://login:pass@localhost:5432/db-name?sslmode=disable
MIGRATE_PATH = ./migration/postgres/asset

up:
	docker compose  up --build -d --force-recreate
	docker compose logs -f

down:
	docker compose down

run: mod
	go run ./cmd/app

mod:
	go mod tidy

mod-update:
	go get -u all
	go mod tidy

lint:
	golangci-lint run

test:
	go test -v -cover ./...

test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

migrate-install:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.1

migrate-create:
	migrate create -ext sql -dir "$(MIGRATE_PATH)" magration-name

migrate-up:
	migrate -database "$(DB_MIGRATE_URL)" -path "$(MIGRATE_PATH)" up

migrate-down:
	migrate -database "$(DB_MIGRATE_URL)" -path "$(MIGRATE_PATH)" down -all