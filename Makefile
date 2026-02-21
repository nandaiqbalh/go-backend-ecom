# Simple Makefile for building, testing, running and managing database migrations

build:
	@go build -o bin/go-backend-ecom cmd/main.go

# run unit tests for all packages
test:
	@go test -v ./...

# compile and run the binary
run: build
	@./bin/go-backend-ecom

# create a new migration file using the golang-migrate CLI
#
# Usage:
#   make migration add-user-table
#
# The target looks for the migrate binary in $(go env GOPATH)/bin so you
# don't need to have that directory in your PATH when running make. If the
# binary is missing, you'll receive instructions on how to install it.

# computed location of the migrate CLI
MIGRATE_BIN := $(shell go env GOPATH)/bin/migrate

migration:
	@# the first word after "make migration" is treated as the name
	@# for the new migration file; any other arguments are ignored.
	@# $(filter-out $@,$(MAKECMDGOALS)) expands to that name.
	@if [ ! -x "$(MIGRATE_BIN)" ]; then \
		echo "Error: 'migrate' CLI not found at $(MIGRATE_BIN)."; \
		echo "Install it with:"; \
		echo "  go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest"; \
		echo "and ensure \$$GOPATH/bin is on your PATH or rerun make."; \
		exit 1; \
	fi
	@$(MIGRATE_BIN) create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

# apply all pending migrations by running the migration tool located in cmd/migrate/main.go
migrate-up:
	@go run cmd/migrate/main.go up

# rollback the last migration
migrate-down:
	@go run cmd/migrate/main.go down