/*
 * @File: main.go
 * @Description: main package
 * @LastModifiedTime: 2021-01-16 08:34:54
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

package main

import (
	"database/sql"
	"log"

	"github.com/Akshit8/go-api/api"
	db "github.com/Akshit8/go-api/db/sqlc"
	"github.com/Akshit8/go-api/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
