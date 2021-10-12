package app

import (
	"database/sql"
	"github.com/zaenalarifin12/golang-restfull-api/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:rahasia@tcp(localhost:3306)/golang_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime( 60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}


