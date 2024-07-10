ifneq (,$(wildcard ./.env))
    include .env
    export DB_HOST DB_PORT DB_DATABASE DB_USER DB_PASSWORD
endif

DB_STRING="postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=disable"

.swag:
	./cmd/swag fmt
	./cmd/swag init -q

.wait-for-pg:
	./cmd/wait-for-postgres.sh

up:
	docker-compose up -d

down:
	docker-compose down --remove-orphans

dev: up .swag .wait-for-pg db-up
	./cmd/air -c .air.toml

build:
	go build -o ./cmd/api main.go

start:
	./cmd/api

gen:
	./cmd/jet -dsn=${DB_STRING} -schema=public -path=./.gen

drop-db:
	docker-compose up -d postgres
	make .wait-for-pg
	docker-compose exec postgres dropdb -U ${DB_USER} --if-exists ${DB_DATABASE}
	docker-compose exec postgres createdb -U ${DB_USER} ${DB_DATABASE}

reset-db: drop-db db-up

db-status:
	./cmd/goose -dir=".gen/migrations" postgres ${DB_STRING} status

db-up:
	./cmd/goose -dir=".gen/migrations" postgres ${DB_STRING} up

db-down:
	./cmd/goose -dir=".gen/migrations" postgres ${DB_STRING} down

db-redo:
	./cmd/goose -dir=".gen/migrations" postgres ${DB_STRING} redo

db-new:
	./cmd/goose -dir=".gen/migrations" postgres ${DB_STRING} create ${name} sql
