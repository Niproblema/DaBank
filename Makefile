postgres:
	docker run --name DaBankPG -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=pspass -d postgres

createdb:
	docker exec -it DaBankPG createdb --username=root --owner=root DaBank

dropdb:
	docker exec -it DaBankPG dropdb DaBank

migrateup: 
	migrate -path db/migration -database "postgresql://root:pspass@localhost:5432/DaBank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:pspass@localhost:5432/DaBank?sslmode=disable" -verbose down

sqlc: 
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc