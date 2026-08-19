package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	_ "github.com/lib/pq"
)

const dbDriver = "postgres"
const dsn = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"

var testQueries *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDb, err = sql.Open(dbDriver, dsn)

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