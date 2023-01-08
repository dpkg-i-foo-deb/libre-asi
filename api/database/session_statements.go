package database

import (
	"database/sql"
	"libre-asi-api/util"
)

var LoginAdminStatement *sql.Stmt
var LoginInterviewerStatement *sql.Stmt
var CheckAdminStatement *sql.Stmt
var CreateUserStatement *sql.Stmt
var CreateAdminStatement *sql.Stmt
var err error

func prepareSessionStatements() {
	LoginAdminStatement, err = DB.Prepare(`select u.email , u."password" 
												from administrator a 
													join "user" u on u.id = a.id 
														where u.email = $1`)

	util.HandleErrorStop(err)

	LoginInterviewerStatement, err = DB.Prepare(`select u.email , u."password"
													from "user" u
														join interviewer i on i.person  = u.id 
															where u.email = $1
													`)
	util.HandleErrorStop(err)

	CheckAdminStatement, err = DB.Prepare(`SELECT id
												FROM administrator`)

	util.HandleErrorStop(err)

	CreateUserStatement, err = DB.Prepare(`INSERT INTO public."user"
													(email, username, "password", created_at, updated_at)
													VALUES($1, $2, $3, $4, $5) RETURNING id`)

	util.HandleErrorStop(err)

	CreateAdminStatement, err = DB.Prepare(`INSERT INTO public.administrator
												(id, created_at, updated_at)
												VALUES($1, $2, $3);
												`)
	util.HandleErrorStop(err)
}
