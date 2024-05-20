package main

import (
	"log"
	"database/sql"
	_ "github.com/lib/pq"
	db "github.com/ekefan/bank_panda/db/sqlc"
	api "github.com/ekefan/bank_panda/api"
	"github.com/ekefan/bank_panda/utils"
)

func main() {
	config, err :=  utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	dbConn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	store := db.NewStore(dbConn)
	server, err := api.NewServer(store, config)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot Start Server: ", err)
	}



}