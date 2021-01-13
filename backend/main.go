package main
import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
    gorm.Model
    Id int
    Name string
}

type C struct {
    Id int
    Name string
  }

func main() {

    db, err := gorm.Open("mysql", "user:password@tcp(db:3306)/yource?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        //panic("failed to connect database")
    }
    defer db.Close()

    r := gin.Default()
    r.LoadHTMLGlob("template/*.tmpl")
    r.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "a": "a",
            "b": []string{"b_todo1","b_todo2"},
            "c": []C{{1,"c_mika"},{2,"c_risa"}},
            "d": C{3,"d_mayu"},
            "e": true,
            "f": false,
            "h": false,
          })
    })
    r.Run(":3001")
}