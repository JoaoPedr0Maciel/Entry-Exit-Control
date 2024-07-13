package database

import (
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitDB() error {
	var err error

	// initialize database
	db, err = DbConnection()

	if err != nil {
		return err
	}

	return nil
}

func GetMysql() *gorm.DB {
	return db
}