/*
Package main is main package.
*/
package main

import (
	//標準ライブラリ
	//自作ライブラリ
	"main/server"
	//githubライブラリ
)

// Main is used to start the gin server.
func main() {
	r := server.Initiate()
	r.Run(":3001")
}

// To exec Mailtest, then use under main().
// And run `go run *.go`
// func main(){
// 	Mailtest()
// }
