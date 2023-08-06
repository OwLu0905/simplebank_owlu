package main

import (
	"database/sql"
	"log"

	"github.com/OwLu0905/simplebank_owlu/api"
	"github.com/OwLu0905/simplebank_owlu/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver       = "postgres"
	dbSource       = "postgresql://root:owlu0905@localhost:5432/simple_bank?sslmode=disable"
	serverADdresss = "0.0.0.0:8888"
)

var testDB *sql.DB

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := sqlc.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverADdresss)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
