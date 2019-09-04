# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=bookstore_platform
BINARY_UNIX=$(BINARY_NAME)_unix
PKGS := $(shell go list ./... | grep -v /vendor)

all: test build

build-current: 
	$(GOBUILD) -o $(BINARY_NAME) -v

build-unix: 
	$(GOBUILD) -o $(BINARY_UNIX) -v

build-unix-2:
	CGO_ENABLED=0 GOOS=linux go build -a -o bookstore_platform_linux

dev:
	$(GOCMD) run main.go

ps: 
	docker ps -a

build: 
	docker-compose build

start:
	docker-compose up -d
	docker ps -a

start-a:
	docker-compose up

stop:
	docker-compose down
	docker ps -a

restart: stop start

test:
	$(GOTEST) -v ./...

lint:
	golangci-lint run

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

tidy:
	go mod tidy

# make new migration: sql-migrate new -config=config/migrations.yml

up:
	sql-migrate up -config=config/migrations.yml

down:
	sql-migrate down -config=config/migrations.yml

down-50:
	seq 50 | xargs -I -- sql-migrate down -config=config/migrations.yml

.PHONY: models
models:
	sqlboiler -c config/sqlboiler.yml --wipe --no-tests psql

models-with-tests:
	sqlboiler --wipe -c config/sqlboiler.yaml psql

populate-db:
	sh populate_db.sh

docs-create:
	swag init

.PHONY: docs
docs:
	swagger serve docs/swagger.yaml

exec-db:
	docker exec -it bookstore_db_dev /bin/bash
