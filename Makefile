watch:
	go tool github.com/air-verse/air -c .air.toml

build: templ sqlc
	go build -o ./tmp/main .

templ:
	go tool github.com/a-h/templ/cmd/templ generate
	
sqlc:
	go tool github.com/sqlc-dev/sqlc/cmd/sqlc generate