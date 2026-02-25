MIGRATION_PATH=migrations

include .env
export

.PHONY: migrate-up migrate-down migrate-create

migrate-up:
	migrate -path $(MIGRATION_PATH) -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path $(MIGRATION_PATH) -database "$(DATABASE_URL)" down 1

migrate-create:
	migrate create -ext sql -dir $(MIGRATION_PATH) -seq $(name)
