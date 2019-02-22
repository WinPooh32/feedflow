package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

type model struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

func migrateModels(db *gorm.DB) {
	db.AutoMigrate(
		Error{},
		LoginRequest{},
		NewPageContent{},
		SigninRequest{},
		UploadImg{})
}
