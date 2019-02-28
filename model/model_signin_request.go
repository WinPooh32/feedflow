/*
 * FeedFlow
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

import (
	"github.com/jinzhu/gorm"
)

//SigninRequest model
type SigninRequest struct {
	Base
	LoginRequest

	Email string `json:"email" form:"email"`

	Salt         []byte
	PasswordHash []byte
}

//ValidSigninRequest - validate SigninRequest
func ValidSigninRequest(db *gorm.DB, sr *SigninRequest) bool {

	if len(sr.Password) < 10 {
		return false
	}

	if !emailRegexp.MatchString(sr.Email) {
		return false
	}

	var found SigninRequest
	db.First(&found, "Username = ? OR Email = ?", sr.Username, sr.Email)

	return found.ID == 0
}
