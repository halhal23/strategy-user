package infrastructure

import (
	"strategy-user/interface/controller"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()
	controller := controller.NewUserController(NewSqlHandler())
	router.GET("/", func(c *gin.Context){ c.JSON(200, "Hello world in strategy.")})
	router.GET("/users", func(c *gin.Context){ controller.Index(c)})
	router.GET("/users/:id", func(c *gin.Context){ controller.Show(c)})
	router.POST("/users", func(c *gin.Context){ controller.Create(c)})
	Router = router
}