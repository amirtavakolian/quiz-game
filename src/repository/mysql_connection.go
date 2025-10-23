package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

func NewMysqlConnection() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	)

	db, err := sql.Open(os.Getenv("MYSQL_DIALECT"), dsn)
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		panic(err.Error())
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
