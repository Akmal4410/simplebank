package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/akmal4410/simple_bank/util"
	_ "github.com/lib/pq"
)

var testQuires *Queries
var testDb *sql.DB

func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("Cannot load config :", err)
	}
	testDb, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db :", err)
	}
	testQuires = New(testDb)
	os.Exit(m.Run())
}
