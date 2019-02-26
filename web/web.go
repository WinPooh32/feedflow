package web

import (
	"log"
	"net/http"
	"reflect"

	"github.com/WinPooh32/feedflow/model"
	ginsession "github.com/go-session/gin-session"

	"github.com/WinPooh32/feedflow/database"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
)

func index(ctx *gin.Context) {
	if db, ok := database.Extract(ctx); ok {

		db.First(model.NewPageContent{})

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

		err := store.Save()
		if err != nil {
			ctx.AbortWithError(500, err)
			return
		}

		gintemplate.HTML(ctx, http.StatusOK, "index", gin.H{
			"title": "Hello, web!",
			"hits":  hits,
		})
	} else {
		ctx.AbortWithStatus(http.StatusInternalServerError)
	}
}

func notFound(ctx *gin.Context) {
	gintemplate.HTML(ctx, http.StatusNotFound, "404.html", gin.H{})
}
