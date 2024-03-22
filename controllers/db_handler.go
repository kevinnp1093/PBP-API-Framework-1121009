package controllers

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Conn *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/db_latihan_pbp")
	if err != nil {
		panic(err.Error())
	}

	Conn = db
}
