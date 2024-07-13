package migrations

import (
	"entry-exit-api/schemas"

	"gorm.io/gorm"
)

func MigrateResidentTable(db *gorm.DB) error {
 	if err := db.AutoMigrate(&schemas.Resident{}); err != nil {
	  return err
 	}
 	return nil
}
