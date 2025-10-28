package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func NewGormConnection() *gorm.DB {
	user := os.Getenv("GORM_USER")
	pass := os.Getenv("GORM_PASS")
	protocol := os.Getenv("GORM_PROTOCOL")
	ip := os.Getenv("GORM_IP")
	port := os.Getenv("GORM_PORT")
	database := os.Getenv("GORM_DATABASE")

	dsn := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&multiStatements=true", user, pass, protocol, ip, port, database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}
