package database

import (
	"os"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectDB() *gorm.DB {

	//環境変数を読み込む
	err := godotenv.Load(fmt.Sprintf("/go/app/%s.env", os.Getenv("GO_ENV")))
    if err != nil {
        panic(err.Error())
    }

	DBMS     := "mysql"
	USER     := os.Getenv("DB_USER")
	PASS     := os.Getenv("DB_PASSWORD")
	PROTOCOL := os.Getenv("DB_HOST")
	DBNAME   := os.Getenv("DB_NAME")
  
	CONNECT := USER+":"+PASS+"@"+PROTOCOL+"/"+DBNAME
	db,err := gorm.Open(DBMS, CONNECT)
  
	if err != nil {
	  panic(err.Error())
	}
	return db
}
