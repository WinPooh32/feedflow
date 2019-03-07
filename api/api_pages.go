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

	//Try to bind form data to person
	if err := ctx.ShouldBind(&person); err != nil || !model.ValidSigninRequest(db, &person) {
		ctx.AbortWithStatus(http.StatusNotAcceptable)
		return
	}

	//Generate random salt using crypto/rand
	salt := make([]byte, 8)
	if _, err := rand.Read(salt); err != nil {
		ctx.AbortWithStatus(http.StatusNotAcceptable)
		return
	}

	//Concatinate user password and salt
	saltedHash := append([]byte(person.Password), salt...)

	//Calc slated password hash
	hash, err := bcrypt.GenerateFromPassword(saltedHash, bcrypt.MinCost+2)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	person.Salt = salt
	person.PasswordHash = hash
	person.DeletedAt = nil // gorm sets time for form post

	db.Create(&person)

	if err := loginSessionUpgrade(person, ctx); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}

// Login - Logs in and returns the authentication cookie.
func Login(ctx *gin.Context) {
	var form model.LoginRequest
	var person model.SigninRequest

	//Check user session previlegies, escape if privileged
	store := ginsession.FromContext(ctx)
	rawID, ok := store.Get("user_id")

	if ok && rawID.(float64) != 0 {
		//User has been already logged in
		ctx.Status(http.StatusOK)
		return
	}

	db, ok := database.FromContext(ctx)
	if !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	//Try to bind form data
	if ctx.ShouldBind(&form) != nil {
		ctx.AbortWithStatus(http.StatusNotAcceptable)
		return
	}

	//Select person from database
	db.First(&person, "username = ?", form.Username)
	if person.ID == 0 {
		ctx.AbortWithStatus(http.StatusNotAcceptable)
		return
	}

	//Concatinate user password and salt
	salted := append([]byte(form.Password), person.Salt...)

	if err := bcrypt.CompareHashAndPassword(person.PasswordHash, salted); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err := loginSessionUpgrade(person, ctx); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func loginSessionUpgrade(person model.SigninRequest, ctx *gin.Context) error {
	store := ginsession.FromContext(ctx)

	//Upgrade user session previlegies
	store, err := ginsession.Refresh(ctx)
	if err != nil {
		return err
	}

	hits, _ := store.Get("visit_hits")
	store.Flush()

	store.Set("visit_hits", hits)
	store.Set("user_id", person.ID)

	if err := store.Save(); err != nil {
		return err
	}

	return nil
}
