package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/SmoothWay/simpleBank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("error loading config:", err)
	}
	testDB, err = sql.Open(config.DBDRiver, config.DSN)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
