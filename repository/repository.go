package repository

import (
	"database/sql"

	"github.com/mFsl16/go-restapi-example/utils"
)

type Database struct {
	Mysql *sql.DB
}

func NewDB() *Database {
	database := Database{}
	d, err := sql.Open("mysql", "user:secret@tcp(localhost:3306)/learn")
	utils.PanicIfErr(err)
	database.Mysql = d
	return &database
}
