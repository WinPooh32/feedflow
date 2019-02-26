package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Model - used for embedding common database fields
type Model struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

//MigrateModels - auto migrate orm models
func MigrateModels(db *gorm.DB) {
	db.AutoMigrate(
		Tag{},
		NewPageContent{},
		SigninRequest{})
}
