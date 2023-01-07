package database

import (
	"database/sql"
	"libre-asi-api/util"
)

var LoginStatement *sql.Stmt
var CheckAdminStatement *sql.Stmt
var CreateUserAdminStatement *sql.Stmt
var CreateAdminStatement *sql.Stmt
var err error

func prepareSessionStatements() {
	LoginStatement, err = DB.Prepare(`SELECT email,"password" 
											FROM "user" u 
												WHERE email = $1`)

	util.HandleErrorStop(err)

	CheckAdminStatement, err = DB.Prepare(`SELECT id
												FROM administrator`)

	util.HandleErrorStop(err)

	CreateUserAdminStatement, err = DB.Prepare(`INSERT INTO public."user"
													(email, username, "password", created_at, updated_at)
													VALUES($1, $2, $3, $4, $5) RETURNING id`)

	util.HandleErrorStop(err)

	CreateAdminStatement, err = DB.Prepare(`INSERT INTO public.administrator
												(id, created_at, updated_at)
												VALUES($1, $2, $3);
												`)
	util.HandleErrorStop(err)
}
