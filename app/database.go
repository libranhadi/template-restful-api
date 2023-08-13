package app

import(
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"template-restful-api/helper"
	"time"
)

func NewDb() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/template-restful-api")

	helper.PanicIfError(err)
	db.SetConnMaxLifetime(time.Minute * 10)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db
}