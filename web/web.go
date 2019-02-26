package web

import (
	"net/http"

	"github.com/WinPooh32/feedflow/database"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func index(ctx *gin.Context) {
	if _, ok := database.Extract(ctx); ok {

		gintemplate.HTML(ctx, http.StatusOK, "index", gin.H{
			"title": "Hello, web!",
		})
	} else {
		ctx.Writer.WriteHeader(http.StatusInternalServerError)
	}
}

func notFound(ctx *gin.Context) {
	gintemplate.HTML(ctx, http.StatusNotFound, "404.html", gin.H{})
}
