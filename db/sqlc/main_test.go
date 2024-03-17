package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:pspass@localhost:5432/DaBank?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot conntect to the db", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
