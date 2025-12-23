APP_NAME=daily-uwoh
CMD_PATH=./cmd/server
DB_URL=postgresql://postgres:shiratama@localhost:5432/dailyuwoh?sslmode=disable

all: build

build:
	go build -o bin/server.exe $(CMD_PATH)

run:
	go run $(CMD_PATH)

test:
	go test ./... -v

clean:
	if exist bin rmdir /s /q bin

tidy:
	go mod tidy

sqlc:
	sqlc generate

createdb:
	docker exec -it postgres18 createdb --username=postgres --owner=postgres dailyuwoh

dropdb:
	docker exec -it postgres18 dropdb dailyuwoh

migrateup:
	migrate -path internal/db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path internal/db/migration -database "$(DB_URL)" -verbose down

server:
	go run $(CMD_PATH)

.PHONY: all build run test clean tidy sqlc createdb dropdb migrateup migratedown server
