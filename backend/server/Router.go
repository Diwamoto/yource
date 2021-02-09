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

	//apiはhttps://hogehoge.com/api/v1以下のルーティングで判断する
	v1 := router.Group("/api/v1")
	{
		//ユーザルーティング
		v1.GET("/users", controller.GetAllUserAction)
		v1.POST("/users", controller.CreateUserAction)
		v1.GET("/users/:id", controller.GetUserAction)
		v1.PUT("/users/:id", controller.UpdateUserAction)
		v1.DELETE("/users/:id", controller.DeleteUserAction)

		//プロフィールルーティング
		//get → プロフィール取得
		//post → プロフィール追加
		//put → プロフィール変更
		v1.GET("/users/:id/profile", controller.GetUserProfileAction)
		v1.POST("/users/:id/profile", controller.CreateUserProfileAction)
		v1.PUT("/users/:id/profile", controller.UpdateUserProfileAction)

		//スペースルーティング
		//get → スペース取得
		//post → スペース作成
		//put → スペース変更
		//delete → スペース削除
		v1.GET("/spaces", controller.GetAllSpaceAction)
		v1.POST("/spaces", controller.CreateSpaceAction)
		v1.GET("/spaces/:id", controller.GetSpaceAction)
		v1.PUT("/spaces/:id", controller.UpdateSpaceAction)
		v1.DELETE("/spaces/:id", controller.DeleteSpaceAction)

		//チャンネルルーティング
		//get → チャンネル取得
		//post → チャンネル追加(スペースからしかできない)
		//put → チャンネル変更
		//delete → チャンネル削除
		v1.GET("/channels", controller.GetAllChannelAction)
		v1.GET("/channels/:id", controller.GetChannelAction)
		v1.PUT("/channels/:id", controller.UpdateChannelAction)
		v1.DELETE("/channels/:id", controller.DeleteChannelAction)

		v1.GET("/spaces/:id/channels", controller.GetChannelBySpaceIdAction) //指定スペースのチャンネル全てを取得
		v1.POST("/spaces/:id/channels", controller.CreateChannelAction)      //指定スペースのチャンネルを作成

		//ポストルーティング
		//get → チャンネル取得
		//post → チャンネル追加(スペースからしかできない)
		//put → チャンネル変更
		//delete → チャンネル削除
		v1.GET("/posts", controller.GetAllPostAction)
		v1.GET("/posts/:id", controller.GetPostAction)
		v1.PUT("/posts/:id", controller.UpdatePostAction)
		v1.DELETE("/posts/:id", controller.DeleteChannelAction)

		v1.GET("/channels/:id/posts", controller.GetPostByChannelIdAction) //指定チャンネルIDの投稿を全て呼び出す
		v1.POST("/channels/:id/posts", controller.CreatePostAction)        //指定チャンネルに投稿を作成する

		v1.GET("/users/:id/posts", controller.GetPostByUserIdAction) //指定ユーザの投稿を検索する

		//ログイン
		v1.POST("/login", controller.LoginAction)

	}

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
