dev:
	go run .

build: templ
	go build -o ./tmp/main .

install:
	go install github.com/a-h/templ/cmd/templ@latest

templ:
	templ generate