package main

import (
	"database/sql"
	"log"

	"github.com/SmoothWay/simpleBank/api"
	db "github.com/SmoothWay/simpleBank/db/sqlc"
	"github.com/SmoothWay/simpleBank/util"

	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("error loading config:", err)
	}

	conn, err := sql.Open(config.DBDRiver, config.DSN)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	err = conn.Ping()
	if err != nil {
		log.Fatal(err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start the server:", err)
	}
}
