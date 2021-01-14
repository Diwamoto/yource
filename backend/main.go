package main
import (
    //標準ライブラリ
    
    //自作ライブラリ
    "main/config"
    
    //githubライブラリ 
)

func main() {

    r := config.GetRouter()
    r.Run(":3001")
}