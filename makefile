	
migrateup:
	migrate -path repository/migration -database "${POSTGRES_URL}" -verbose up

migratedown:
	migrate -path repository/migration -database "${POSTGRES_URL}" -verbose down

dropdb:
	docker exec -it postgres12 dropdb laundry

dockerup:
	docker compose up -d

sqlc:
	sqlc generate

	

.PHONY: migrateup migratedown dropdb dockerup sqlc 

include .env