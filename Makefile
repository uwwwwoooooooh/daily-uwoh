APP_NAME=daily-uwoh
CMD_PATH=./cmd/server
include app.env

all: build

build:
	go build -o bin/$(APP_NAME) $(CMD_PATH)

run:
	go run $(CMD_PATH)

test:
	go test ./... -v

clean:
	rm -rf bin

tidy:
	go mod tidy

sqlc:
	sqlc generate

dockerdb:
	docker run --name uwohdb --network uwoh-network -p 5432:5432 -e POSTGRES_PASSWORD=shiratama -d postgres:18-alpine

createdb:
	docker exec -it uwohdb createdb --username=root --owner=root dailyuwoh

dropdb:
	docker exec -it uwohdb dropdb dailyuwoh

migrateup:
	migrate -path internal/db/migration -database "$(DATABASE_URL)" -verbose up

migratedown:
	migrate -path internal/db/migration -database "$(DATABASE_URL)" -verbose down

server:
	go run $(CMD_PATH)

mock:
	mockgen -package mockdb -destination internal/db/mock/store.go github.com/uwwwwoooooooh/daily-uwoh/internal/db/sqlc Store

.PHONY: all build run test clean tidy sqlc dockerdb dockerserver createdb dropdb migrateup migratedown server mock
