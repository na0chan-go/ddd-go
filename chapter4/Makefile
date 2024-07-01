up:
	docker compose up -d
down:
	docker compose down
destroy:
	docker compose down --volumes
migration:
	docker compose exec app bash db-migration.sh
run:
	docker compose exec app go run main.go
