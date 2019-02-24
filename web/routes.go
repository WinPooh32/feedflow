package web

import "github.com/gin-gonic/gin"

//RouteWeb - Define web pages routes
func RouteWeb(router *gin.Engine) {
	router.GET("/", index)
}
