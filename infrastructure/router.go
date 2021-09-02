package infrastructure

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context){
		ctx.String(200, "hello strategy-user")
	})
	return router
}