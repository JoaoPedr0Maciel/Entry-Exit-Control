package database

import (
	"entry-exit-api/migrations"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DbConnection() (*gorm.DB, error) {
	connection_string := "jpdev:12345678@tcp(localhost:3306)/entry_exit_control?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(connection_string), &gorm.Config{})
	if err != nil {
		println(err)
		return nil, err
	}

	err = migrations.MigrateResidentTable(db)
	if err != nil {
		println(err)
		return nil, err
	}

	err = migrations.AddTableVisitor(db)
	if err != nil {
		println(err)
		return nil, err
	}

	return db, nil

}
 