run: templ sqlc
	go run .

.PHONY: test
test:
	go test -v ./...

templ:
	go tool github.com/a-h/templ/cmd/templ generate
	
sqlc:
	go tool github.com/sqlc-dev/sqlc/cmd/sqlc generate