/*
 * FeedFlow
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

//LoginRequest model
type LoginRequest struct {
	Username string `json:"username" form:"username" sql:"index"`
	Password string `json:"password" form:"password" gorm:"-" `
}
