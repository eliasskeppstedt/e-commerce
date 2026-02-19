dbup:
	docker compose up -d

dbdown:
	docker compose down

dbreset:
	docker compose down -v
	docker compose up -d

dbreset:
	goose reset
	goose up

run:
	go run cmd/server/main.go
