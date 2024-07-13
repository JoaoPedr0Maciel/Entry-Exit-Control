package migrations

import (
	"entry-exit-api/schemas"

	"gorm.io/gorm"
)


func AddTableVisitor(db *gorm.DB) error {
	
	if err := db.AutoMigrate(schemas.Visitor{}); err != nil {
		return err
	}

	return nil
}