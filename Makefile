postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=owlu0905 -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres15 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:owlu0905@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:owlu0905@localhost:5432/simple_bank?sslmode=disable" -verbose down

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
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test retest start-postgres server mock
