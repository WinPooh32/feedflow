package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type model struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

//MigrateModels - auto migrate orm models
func MigrateModels(db *gorm.DB) {
	db.AutoMigrate(
		Error{},
		LoginRequest{},
		Tag{},
		NewPageContent{},
		SigninRequest{},
		UploadImg{})
}