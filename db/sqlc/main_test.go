/*
 * @File: db.sqlc.main_test.go
 * @Description: entrypoint to all db tests
 * @LastModifiedTime: 2021-01-14 16:47:30
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Akshit8/go-api/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	log.Printf("dbdriver: %s", config.DBDriver)
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}
