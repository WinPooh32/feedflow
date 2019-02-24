/*
 * FeedFlow
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package model

//NewPageContent model
type NewPageContent struct {
	model

	Title   string `json:"title"`
	Content string `json:"content"`
	Tags    []Tag  `json:"tags" gorm:"auto_preload"`
}

//Tag model
type Tag struct {
	model

	Value            string
	NewPageContentID uint `json:"-"`
}