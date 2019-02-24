package web

import (
	"net/http"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func index(ctx *gin.Context) {
	gintemplate.HTML(ctx, http.StatusOK, "index", gin.H{
		"title": "Hello, web!",
	})
}

func notFound(ctx *gin.Context) {
	gintemplate.HTML(ctx, http.StatusNotFound, "404.html", gin.H{})
}
