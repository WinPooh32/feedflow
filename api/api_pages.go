/*
 * FeedFlow
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

import (
	"crypto/rand"
	"net/http"

	"github.com/WinPooh32/feedflow/database"
	ginsession "github.com/go-session/gin-session"
	"golang.org/x/crypto/bcrypt"

	"github.com/WinPooh32/feedflow/model"

	"github.com/gin-gonic/gin"
)

// Add - Add a new page.
func Add(ctx *gin.Context) {
	pagecontent := model.NewPageContent{Tags: make([]model.Tag, 0)}

	db, ok := database.FromContext(ctx)
	if !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if ctx.ShouldBind(&pagecontent) == nil && model.ValidNewPageContent(&pagecontent) {
		db.Create(&pagecontent)
		ctx.Status(http.StatusOK)
		return
	}

	ctx.AbortWithStatus(http.StatusNotAcceptable)
}

// ImgUpload - Upload a new image.
func ImgUpload(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// Remove - Move page to archive.
func Remove(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// Signin - Singns in.
func Signin(ctx *gin.Context) {
	var person model.SigninRequest

	db, ok := database.FromContext(ctx)
	if !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if ctx.ShouldBind(&person) == nil &&
		model.ValidSigninRequest(db, &person) {

		salt := make([]byte, 8)
		_, err := rand.Read(salt)

		if err == nil {
			mixed := append([]byte(person.Password), salt...)
			hash, _ := bcrypt.GenerateFromPassword(mixed, bcrypt.MinCost+2)

			person.Salt = salt
			person.PasswordHash = hash
			person.DeletedAt = nil // gorm sets time for form post

			db.Create(&person)

			store := ginsession.FromContext(ctx)
			store.Set("user_id", person.ID)

			err := store.Save()
			if err != nil {
				ctx.AbortWithError(http.StatusInternalServerError, err)
				return
			}

			ctx.Status(http.StatusOK)
			return
		}
	}

	ctx.AbortWithStatus(http.StatusNotAcceptable)
}

// Login - Logs in and returns the authentication cookie.
func Login(ctx *gin.Context) {
	var form model.LoginRequest
	var person model.SigninRequest

	db, ok := database.FromContext(ctx)
	if !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if ctx.ShouldBind(&form) == nil {
		db.First(&person, "username = ?", form.Username)

		if person.ID != 0 {
			mixed := append([]byte(form.Password), person.Salt...)

			if bcrypt.CompareHashAndPassword(person.PasswordHash, mixed) == nil {
				store := ginsession.FromContext(ctx)

				hits, _ := store.Get("visit_hits")
				store.Flush()

				store.Set("visit_hits", hits)
				store.Set("user_id", person.ID)
				err := store.Save()
				if err != nil {
					ctx.AbortWithError(http.StatusInternalServerError, err)
					return
				}

				ctx.Status(http.StatusOK)
				return
			}
		}
	}

	ctx.AbortWithStatus(http.StatusNotAcceptable)
}
