package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Conn() *gorm.DB {
	dsn := "test:test@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
