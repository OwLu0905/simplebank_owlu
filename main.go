package main

import (
	"database/sql"
	"log"

	"github.com/OwLu0905/simplebank_owlu/api"
	"github.com/OwLu0905/simplebank_owlu/db/sqlc"
	"github.com/OwLu0905/simplebank_owlu/util"

	_ "github.com/lib/pq"
)

var testDB *sql.DB

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := sqlc.NewStore(conn)
	// server := api.NewServer(store)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
