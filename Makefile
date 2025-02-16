watch:
	air -c .air.toml

dev:
	go run .

build: sqlc
	go build -o ./tmp/main .

install:
	go install github.com/air-verse/air@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

sqlc:
	sqlc generate