# Makefile for managing database migrations

MIGRATIONS_DIR := ./migrations
MIGRATE := migrate

.PHONY: create migrate_up migrate_down

# Create a new migration file
create:
	@read -p "Enter migration name: " name; \
	$(MIGRATE) create -ext sql -dir $(MIGRATIONS_DIR) $$name

# Apply all up migrations
migrate_up:
	$(MIGRATE) -database $(DATABASE_URL) -path $(MIGRATIONS_DIR) up

# Roll back all migrations
migrate_down:
	$(MIGRATE) -database $(DATABASE_URL) -path $(MIGRATIONS_DIR) down

# Apply 'n' up migrations
migrate_up_n:
	@read -p "Enter the number of migrations to apply: " n; \
	$(MIGRATE) -database $(DATABASE_URL) -path $(MIGRATIONS_DIR) up $$n

# Roll back 'n' migrations
migrate_down_n:
	@read -p "Enter the number of migrations to roll back: " n; \
	$(MIGRATE) -database $(DATABASE_URL) -path $(MIGRATIONS_DIR) down $$n

# Force the version number to 'n', regardless of the current state
force_version:
	@read -p "Enter the version number to force: " n; \
	$(MIGRATE) -database $(DATABASE_URL) -path $(MIGRATIONS_DIR) force $$n

.PHONY: compile-server
compile-server:
	GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o ./bin/server ./cmd/server/main.go