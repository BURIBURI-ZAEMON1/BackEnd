package database

import (
	"backend/app/model"
	"gorm.io/gorm"
)

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.User{},
		&model.Post{},
		&model.Report{},
	)
}
