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
	router.GET("/users", controller.GetAllUserAction)
	router.POST("/users", controller.CreateUserAction)
	router.GET("/users/:id", controller.GetUserAction)
	router.PUT("/users/:id", controller.UpdateUserAction)
	router.DELETE("/users/:id", controller.DeleteUserAction)
	return router
}
