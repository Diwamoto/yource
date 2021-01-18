package main

import (
	//標準ライブラリ
	//自作ライブラリ
	"main/server"
	//githubライブラリ
)

func main() {
	r := server.GetRouter()
	r.Run(":3001")
}
