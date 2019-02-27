package web

import (
	"log"
	"net/http"
	"reflect"

	ginsession "github.com/go-session/gin-session"

	"github.com/WinPooh32/feedflow/database"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func index(ctx *gin.Context) {
	if _, ok := database.FromContext(ctx); ok {

		// db.First(model.NewPageContent{})

		store := ginsession.FromContext(ctx)

		var hits float64
		hitsI, ok := store.Get("visit_hits")
		if ok {
			tmp, _ := hitsI.(float64)
			hits = tmp
			log.Println("HITS: ", hits, reflect.TypeOf(hitsI))
		}

		hits++
		store.Set("visit_hits", hits)

		userID, ok := store.Get("user_id")

		err := store.Save()
		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}

		gintemplate.HTML(ctx, http.StatusOK, "index", gin.H{
			"title":   "Hello, web!",
			"hits":    hits,
			"user_id": userID,
		})
	} else {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
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
