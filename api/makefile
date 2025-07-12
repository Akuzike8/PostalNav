include .env
export

DATABASE_URL := $(DBUSER):$(DBPASS)@tcp($(DBHOST):$(DBPORT))/$(DBNAME)
VERSION := latest
MIGRATIONS_DIR := internal/database/migrations
NAME := latest

up:
	@echo "Running migrations up..."
	migrate -path internal/database/migrations -database "mysql://$(DATABASE_URL)" up

down:
	@echo "Reverting migrations down..."
	migrate -path internal/database/migrations -database "mysql://$(DATABASE_URL)" down

create:
	@if [ -z "$(NAME)" ]; then echo "Error: NAME is required. Use: make migrate-create NAME=name"; exit 1; fi;
	@echo "Creating new migration..."
	migrate create -ext sql -dir internal/database/migrations -seq $(NAME)

force:
	@if [ -z "$(VERSION)" ]; then echo "Error: VERSION is required. Use: make migrate-force VERSION=version"; exit 1; fi
	$(MIGRATE) -path $(MIGRATIONS_DIR) -database "mysql://$(DATABASE_URL)" force $(VERSION)
