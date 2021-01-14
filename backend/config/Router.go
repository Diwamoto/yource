package config

import (
	//標準ライブラリ

	//自作ライブラリ
	"main/controller"

	//githubライブラリ 
	"github.com/gin-gonic/gin"
)


func GetRouter() *gin.Engine {
    router := gin.Default()
	router.LoadHTMLGlob("view/*.html")
 
	//
	router.POST("/user", controller.CreateUserAction)
	router.GET("/user/:id", controller.GetUserAction)
    return router
}