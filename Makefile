postgres:
	docker run --name postgres15 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=owlu0905 -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres15 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:owlu0905@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:owlu0905@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:owlu0905@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:owlu0905@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test :
	go test -v -cover ./...

retest :
	go clean -testcache && go test -v -cover ./...

start-postgres:
	docker exec -it postgres15 /bin/sh

server: 
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/OwLu0905/simplebank_owlu/db/sqlc Store

production:
	docker run --name simplebank --network bank-network -p 8080:8080 -e DB_SOURCE="postgresql://root:owlu0905@postgres15:5432/simple_bank?sslmode=disable" -e GIN_MODE=release simplebank:latest 
.PHONY: postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test retest start-postgres server mock production

