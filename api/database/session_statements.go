package database

import (
	"database/sql"
	"libre-asi-api/util"
)

var LoginStatement *sql.Stmt
var SignUpStatement *sql.Stmt
var err error

func prepareSessionStatements() {
	LoginStatement, err = DB.Prepare(`SELECT email,"password" 
											FROM "user" u 
												WHERE email = $1`)

	util.HandleErrorStop(err)

}
