PHONY: all

test:
	go test -v ./...

up:
	docker-compose up -d

up_api:
	docker-compose up db_mysql api

migrate:
	docker-compose up db_mysql migrations

generate_sql:
	docker run --rm -it -v $(PWD):/src -w /src kjconroy/sqlc generate