/*
 * FeedFlow
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

//SigninRequest model
type SigninRequest struct {
	model

	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`

	Salt []byte
}
