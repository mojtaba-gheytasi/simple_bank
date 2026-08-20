package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/mojtaba-gheytasi/simple-bank/api"
	db "github.com/mojtaba-gheytasi/simple-bank/db/sqlc"
	"github.com/mojtaba-gheytasi/simple-bank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}


	conn, err := sql.Open(config.DbDriver, config.DbSource)

	if err != nil {
		log.Fatal(err)
	}

	if err := conn.Ping(); err != nil {
		log.Fatal(err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}


