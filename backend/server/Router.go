package server

import (
	//標準ライブラリ
	"net/http"

	//自作ライブラリ
	"main/controller"

	//githubライブラリ
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	router.LoadHTMLGlob("view/*.html")

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome to Yource API v0.1")
	})

	//ユーザルーティング
	router.POST("/user", controller.CreateUserAction)
	router.GET("/user/:id", controller.GetUserAction)
	router.PUT("/user/:id", controller.UpdateUserAction)
	router.DELETE("/user/:id", controller.DeleteUserAction)
	return router
}
