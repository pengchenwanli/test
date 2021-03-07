package main

import (
	"gorm.io/gorm"
	"test/model"
)

func autoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Account{},
		&model.Asset{},
		&model.Token{},
		&model.Admin{},
	)
}
