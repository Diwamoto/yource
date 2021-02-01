package server

import (
	//標準ライブラリ
	"net/http"
	"os"

	//自作ライブラリ
	"main/controller"

	//githubライブラリ
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	//corsの設定
	router.Use(cors.Default())

	router.LoadHTMLGlob("view/*.html")

	//apikeyの認証を行う
	//router.Use(CheckApiKey())

	//ユーザルーティング
	router.GET("/users", controller.GetAllUserAction)
	router.POST("/users", controller.CreateUserAction)
	router.GET("/users/:id", controller.GetUserAction)
	router.PUT("/users/:id", controller.UpdateUserAction)
	router.DELETE("/users/:id", controller.DeleteUserAction)

	router.POST("/users/login", controller.LoginAction)

	return router
}

//ミドルウェア
//リクエストヘッダのapiキーを確認し、なければそれ以降の処理を中断する
func CheckApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {

		err := godotenv.Load(os.Getenv("ENV_PATH"))
		if err != nil {
			panic(err.Error())
		}

		if c.Request.Header.Get("Apikey") == os.Getenv("APIKEY") {
			c.Set("Authorized", true)
		} else {
			c.Set("Authorized", false)
			if c.Request.Header.Get("Apikey") == "" {
				c.JSON(http.StatusBadRequest, "Auth failed: Apikey not found")
			} else {
				c.JSON(http.StatusBadRequest, "Auth failed: Invalid Apikey")
			}
			//ルーティング以降の処理を中断する
			c.Abort()
		}

	}
}
