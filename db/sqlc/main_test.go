package sqlc

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/OwLu0905/simplebank_owlu/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries

// NOTE : In order to reuse the db connection, we should declare a new global variable (then we could export this variable)
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())

}
