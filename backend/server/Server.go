package server

import (
	//標準ライブラリ

	"fmt"
	"io"
	"os"
	"time"

	//自作ライブラリ
	"main/config"
	"main/controller"

	//githubライブラリ

	"github.com/form3tech-oss/jwt-go"
	"github.com/form3tech-oss/jwt-go/request"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Initiate() *gin.Engine {

	server := gin.Default()
	//サーバーのログファイルの出力先を設定
	f, _ := os.Create(config.ToString("rootPath") + "/log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	//環境変数を読み込み
	godotenv.Load(os.Getenv("ENV_PATH"))

	//セッション管理用にredisを設定
	store, _ := redis.NewStore(10, os.Getenv("REDIS_PROTOCOL"), os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"), os.Getenv("REDIS_PASSWORD"), []byte(os.Getenv("REDIS_KEY")))
	//セッションの有効期限一日後を設定
	store.Options(sessions.Options{
		MaxAge: 60 * 60 * 24,
	})
	server.Use(sessions.Sessions("session", store))

	//corsの設定
	server.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:9092",
			"https://localhost:9092",
			"https://yource.space"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	server.LoadHTMLGlob("view/*.html")

	//apiはhttps://hogehoge.com/api/v1以下のルーティングで判断する
	v1 := server.Group("/api/v1")
	{
		//ログインとユーザ作成はセッションなしでもアクセスできる
		v1.POST("/signup", controller.CreateUserAction)
		v1.POST("/login", controller.LoginAction)
		v1.POST("/verify", controller.VerifyUserAction)
		//ログインしている状態の場合のみ以下のルーティングを使用可能
		v1.Use(IsLogin())
		{

			// JWTを元にredisに保存されているuserの情報を取得してくる
			v1.GET("/retrive", controller.RetriveUserByJWTAction)

			//ユーザルーティング
			v1.GET("/users", controller.SearchUserAction)
			v1.POST("/users", controller.CreateUserAction)
			v1.GET("/users/:id", controller.GetUserByIdAction)
			v1.PUT("/users/:id", controller.UpdateUserAction)
			v1.DELETE("/users/:id", controller.DeleteUserAction)

			//プロフィールルーティング
			//get → プロフィール取得
			//post → プロフィール追加
			//put → プロフィール変更
			v1.GET("/users/:id/profile", controller.GetUserByIdAction)
			v1.POST("/users/:id/profile", controller.CreateUserProfileAction)
			v1.PUT("/users/:id/profile", controller.UpdateUserProfileAction)

			//スペースルーティング
			//get → スペース取得
			//post → スペース作成
			//put → スペース変更
			//delete → スペース削除
			v1.GET("/spaces", controller.GetAllSpaceAction)
			v1.POST("/spaces", controller.CreateSpaceAction)
			v1.GET("/spaces/:id", controller.GetSpaceByIdAction)
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
			v1.GET("/posts", controller.SearchUserProfileAction)
			v1.GET("/posts/:id", controller.GetPostAction)
			v1.PUT("/posts/:id", controller.UpdatePostAction)
			v1.DELETE("/posts/:id", controller.DeleteChannelAction)

			v1.GET("/channels/:id/posts", controller.GetPostByChannelIdAction) //指定チャンネルIDの投稿を全て呼び出す
			v1.POST("/channels/:id/posts", controller.CreatePostAction)        //指定チャンネルに投稿を作成する

			v1.GET("/users/:id/posts", controller.GetPostByUserIdAction)  //指定ユーザの投稿を検索する
			v1.GET("/users/:id/space", controller.GetSpaceByUserIdAction) //指定ユーザのスペースを検索する
			v1.POST("/users/:id/space", controller.CreateSpaceAction)     //指定ユーザのスペースを作成する

			// //websocket用のルーティング
			// v1.GET("/ws", func(c *gin.Context) {
			// 	wshandler(c.Writer, c.Request)
			// })
		}

	}
	return server
}

//クッキーのjwtを検証する
func IsLogin() gin.HandlerFunc {
	return func(c *gin.Context) {

		//jwtを検証して存在しなければだめ
		_, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
			b := []byte(os.Getenv("SIGN_KEY"))
			return b, nil
		})
		if err != nil {
			c.JSON(401, gin.H{"error": fmt.Sprint(err)})
			c.Abort()
		}

	}
}
