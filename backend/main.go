package main
import (
    //標準ライブラリ
    "net/http"
    
    //自作ライブラリ
    "main/model"
    
    //githubライブラリ 
    "github.com/gin-gonic/gin"
    // jsoniter "github.com/json-iterator/go"
)

func main() {

    // var json = jsoniter.ConfigCompatibleWithStandardLibrary

    User := model.GetUser(1)

    r := gin.Default()
    //htmlファイルを読み込み
    r.LoadHTMLGlob("template/*.tmpl")
    
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, User)
        c.JSON(http.StatusOK, User.Id)
    })
    r.Run(":3001")
}