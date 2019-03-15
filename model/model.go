package model

import (
	"regexp"
	"time"

	"github.com/jinzhu/gorm"
)

//Base - used for embedding common database fields
type Base struct {
	ID        uint64     `json:"id" gorm:"primary_key"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

//BaseNoJSON - used for embedding common database fields without json fields
type BaseNoJSON struct {
	ID        uint64     `json:"-" gorm:"primary_key"`
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

	//Hash indeces for username in lower case
	db.Exec("CREATE INDEX signin_request_index ON signin_request USING hash ((lower(username)));")

	//Hash indeces for tag values
	db.Exec("CREATE INDEX tag_index ON tag USING hash ((lower(value)));")
}
