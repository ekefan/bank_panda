package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbConnSignal = "postgresql://root:secret@localhost:5432/simpleBank?sslmode=disable"
	dbDriver     = "postgres"
)

var testStore *Store
var testQueries *Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbConnSignal)
	if err != nil {
		log.Fatal("Can not connect to db", err)
	}
	testQueries = New(conn)

	testStore = NewStore(conn)
	// fmt.Println("connection established with the database")
	os.Exit(m.Run())

}
