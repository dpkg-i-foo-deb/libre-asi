package database

import (
	"database/sql"
	"libre-asi-api/util"
)

var CreatePersonStatement *sql.Stmt
var CreateInterviewerStatement *sql.Stmt
var CreatePatientStatement *sql.Stmt

func preparePersonStatements() {
	CreatePersonStatement, err = DB.Prepare(`INSERT INTO person
												(id, first_name, second_name, first_surname, second_surname, age, birthdate, personal_id,created_at, updated_at)
												VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10) RETURNING id;`)

	util.HandleErrorStop(err)

	CreateInterviewerStatement, err = DB.Prepare(`INSERT INTO public.interviewer
													(person, "RMA", created_at, updated_at, profession)
													VALUES($1,$2,$3,$4,$5);
	`)

	util.HandleErrorStop(err)

	CreatePatientStatement, err = DB.Prepare(`INSERT INTO public.patient
												(id, created_at, updated_at, social_security_number, race, religious_preference)
													VALUES($1, $2, $3, $4, $5, $6);
	`)

	util.HandleErrorStop(err)
}
