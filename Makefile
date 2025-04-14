DB_NAME=pvz-db

DB_URL=postgres://postgres:qwerty123@localhost:5432/$(DB_NAME)?sslmode=disable

ifeq ($(os), win)
#win path format
MIGRATIONS_PATH=.\internal\storage\migrations
else 
#unix-like path format (default)
MIGRATIONS_PATH=./internal/storage/migrations
endif

MIGRATE=docker run --rm -v $(MIGRATIONS_PATH):/migrations --network host migrate/migrate -path=/migrations/

TEST_NET_NAME=testnet-pvz
TEST_DB_NAME=test-pvz-db
TEST_SERV_NAME=test-pvz-server 

all: build run

build:
	docker-compose build server

run:
	docker-compose up

stop:
	docker-compose stop

remove:
	docker-compose down

integration:
	cd tests/integration && go test -v -db=$(TEST_DB_NAME)

clear-for-integration:
	docker rm -f $(TEST_DB_NAME)

migrate-up:
	$(MIGRATE) -database $(DB_URL) up
migrate-down:
	$(MIGRATE) -database $(DB_URL) down -all