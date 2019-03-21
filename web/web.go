package web

import (
	"log"
	"net/http"

	"github.com/WinPooh32/feedflow/database"
	"github.com/WinPooh32/feedflow/user"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func index(ctx *gin.Context) {
	if _, ok := database.FromContext(ctx); !ok {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	myuser := user.New(ctx)
	myuser.SessionHit()
	myuser.SessionSave()

	log.Println("User", myuser.SessionGetID(), " HITS: ", myuser.SessionGetHits())

	gintemplate.HTML(ctx, http.StatusOK, "index", gin.H{
		"title": "Hello, web!",
		"hits":  myuser.SessionGetHits(),
		// "user_id": userID,
	})
}

func signin(ctx *gin.Context) {
	gintemplate.HTML(ctx, http.StatusOK, "signin", gin.H{})
}

func login(ctx *gin.Context) {
	gintemplate.HTML(ctx, http.StatusOK, "login", gin.H{})
}

func notFound(ctx *gin.Context) {
	gintemplate.HTML(ctx, http.StatusNotFound, "404.html", gin.H{})
}
