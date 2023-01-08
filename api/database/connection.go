package database

import (
	"database/sql"
	"libre-asi-api/util"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func init() {
	var err error

	DB, err = sql.Open("postgres", os.Getenv("CONNECTION_STRING"))

	util.HandleErrorStop(err)

	prepareSessionStatements()
	preparePersonStatements()
}
