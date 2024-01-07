package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQuires *Queries
var testDb *sql.DB

const (
	driverName = "postgres"
	sourceName = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error
	testDb, err = sql.Open(driverName, sourceName)
	if err != nil {
		log.Fatal("Cannot connect to db :", err)
	}
	testQuires = New(testDb)
	os.Exit(m.Run())
}
