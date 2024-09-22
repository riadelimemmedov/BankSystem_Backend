package db

import (
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func Connect() (*pgxpool.Pool, error) {
	testDB, err := NewDBPool()
	if err != nil {
		return nil, err
	}
	return testDB, nil
}

func TestMain(m *testing.M) {
	testDB, err := Connect()
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer testDB.Close()

	testQueries = New(testDB)

	code := m.Run()
	os.Exit(code)
}
