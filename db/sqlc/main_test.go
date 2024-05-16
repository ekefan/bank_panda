package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/ekefan/bank_panda/utils"
)


var testStore Store
var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Can not connect to db", err)
	}
	testQueries = New(conn)

	testStore = NewStore(conn)
	// fmt.Println("connection established with the database")
	os.Exit(m.Run())

}
