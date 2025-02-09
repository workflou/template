watch:
	air -c .air.toml

dev:
	go run .

build: templ sqlc
	go build -o ./tmp/main .

install:
	go install github.com/air-verse/air@latest
	go install github.com/a-h/templ/cmd/templ@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

templ:
	templ generate
	
sqlc:
	sqlc generate