package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var testStore Store

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:pspass@localhost:5432/DaBank?sslmode=disable"
)

func TestMain(m *testing.M) {
	// conn, err := sql.Open(dbDriver, dbSource)

	connPool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("Cannot conntect to the db", err)
	}

	testStore = NewStore(connPool)

	os.Exit(m.Run())
}
