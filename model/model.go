package model

import (
	"regexp"
	"time"

	"github.com/jinzhu/gorm"
)

//Model - used for embedding common database fields
type Model struct {
	ID        uint64     `json:"id" gorm:"primary_key"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

var emailRegexp *regexp.Regexp

func init() {
	const emailPattern = `^(([^<>()\[\]\.,;:\s@\"]+(\.[^<>()\[\]\.,;:\s@\"]+)*)|(\".+\"))@(([^<>()[\]\.,;:\s@\"]+\.)+[^<>()[\]\.,;:\s@\"]{2,})$`
	emailRegexp = regexp.MustCompile(emailPattern)
}

//MigrateModels - auto migrate orm models
func MigrateModels(db *gorm.DB) {
	db.AutoMigrate(
		Tag{},
		NewPageContent{},
		SigninRequest{})

	db.Exec("CREATE INDEX signin_request_index ON signin_request USING hash (username);")
}
