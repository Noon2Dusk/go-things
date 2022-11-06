PHONY: all

test:
	go test -v ./...

up:
	docker-compose up -d