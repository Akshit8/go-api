/*
 * @File: db.sqlc.main_test.go
 * @Description: entrypoint to init all tests
 * @LastModifiedTime: 2021-01-14 16:47:30
 * @Author: Akshit Sadana (akshitsadana@gmail.com)
 */

package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgres://root:secret@host.docker.internal:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
