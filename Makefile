ifneq (,$(wildcard ./.env))
    include .env
    export DB_HOST DB_PORT DB_DATABASE DB_USER DB_PASSWORD
endif

DB_STRING="postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=disable"

dev:
	air -c .air.toml

gen:
	jet -dsn=${DB_STRING} -schema=public -path=./.gen

db-status:
	goose -dir=".gen/migrations" postgres ${DB_STRING} status

db-up:
	goose -dir=".gen/migrations" postgres ${DB_STRING} up

db-down:
	goose -dir=".gen/migrations" postgres ${DB_STRING} down

db-redo:
	goose -dir=".gen/migrations" postgres ${DB_STRING} redo

db-new:
	goose -dir=".gen/migrations" postgres ${DB_STRING} create ${name} sql
