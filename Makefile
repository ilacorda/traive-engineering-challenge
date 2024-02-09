DC := docker-compose
APP_BINARY := app

# Define the default make target
.PHONY: default
default: run

# Build the application binary
.PHONY: build
build:
	go build -o $(APP_BINARY) cmd/main.go

.PHONY: db-up
db-up:
	$(DC) up -d postgres

.PHONY: db-down
db-down:
	$(DC) down

# Run the application
.PHONY: run
run: db-up build
	./$(APP_BINARY)

# Run the application without rebuilding the binary
.PHONY: run-fast
run-fast: db-up
	./$(APP_BINARY)

.PHONY: clean
clean:
	rm -f $(APP_BINARY)

.PHONY: test
test:
	go test ./... -v

gofumpt:
	gofumpt -w -l .
