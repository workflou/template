dev:
	air

make-migration:
	goose create NAME go

migrate:
	go run ./cmd/migrate/main.go

rollback:
	go run ./cmd/rollback/main.go

reset: rollback migrate

.PHONY: tailwind
tailwind:
	tailwind -i ./cmd/web/static/main.css -o ./cmd/web/static/dist.css --watch

testuser:
	go run ./cmd/testuser/main.go