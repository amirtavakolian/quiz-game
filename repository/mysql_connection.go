package repository

import (
	"database/sql"
	"github.com/amirtavakolian/quiz-game/pkg/configloader"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func NewMysqlConnection() *sql.DB {
	cfgLoader := configloader.NewConfigLoader()
	dbConfig := cfgLoader.SetPrefix("APP_").SetDelimiter(".").SetDivider("_").Build()
	dsn := dbConfig.String("mysql.user") + ":" + dbConfig.String("mysql.password") + "@/" + dbConfig.String("mysql.database")

	db, err := sql.Open(dbConfig.String("mysql.dialect"), dsn)

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
