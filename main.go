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

var (
	dbDriver      = "postgres"
	dbSource      = util.SetPostgresHost()
	serverAddress = "0.0.0.0:8000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
