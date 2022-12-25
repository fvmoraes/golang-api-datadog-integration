package migrations

import (
	"api-sample/models"

	"gorm.io/gorm"
)

func RunMigrations(db gorm.DB) {
	db.AutoMigrate(models.Person{})
}
