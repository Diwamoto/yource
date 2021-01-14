package config

import (
	//標準ライブラリ

	//自作ライブラリ
	"main/controller"

	//githubライブラリ 
	"github.com/gin-gonic/gin"
)


func GetRouter() *gin.Engine {    // *gin.Engineの表記は返り値の型
    router := gin.Default()
	router.LoadHTMLGlob("view/*.html")
 
    // router.GET("/", controller.IndexDisplayAction)
	router.GET("/", controller.Get)
    return router
}