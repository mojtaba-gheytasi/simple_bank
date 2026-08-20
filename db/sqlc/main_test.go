package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/mojtaba-gheytasi/simple-bank/util"
)

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error

	config, err := util.LoadConfig("../../.")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDb, err = sql.Open(config.DbDriver, config.DbSource)

	if err != nil {
		log.Fatal(err)
	}

	if err := testDb.Ping(); err != nil {
		log.Fatal(err)
	}

	testQueries = New(testDb)

	code := m.Run()

	testDb.Close()

	os.Exit(code)

}