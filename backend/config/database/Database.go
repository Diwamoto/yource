package database

import (
	"main/config"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func ConnectDB(t string) *gorm.DB {

	//環境変数を読み込む
	err := godotenv.Load(config.Get("envPath"))
	if err != nil {
		panic(err.Error())
	}
	var DBMS, USER, PASS, PROTOCOL, DBNAME string
	DBMS = "mysql"
	//テストと明示的に指定された場合のみテスト用のデータベースに接続する。
	if t == "test" {
		USER = os.Getenv("DB_TEST_USER")
		PASS = os.Getenv("DB_TEST_PASSWORD")
		PROTOCOL = os.Getenv("DB_TEST_HOST")
		DBNAME = os.Getenv("DB_TEST_NAME")
	} else {
		USER = os.Getenv("DB_USER")
		PASS = os.Getenv("DB_PASSWORD")
		PROTOCOL = os.Getenv("DB_HOST")
		DBNAME = os.Getenv("DB_NAME")
	}

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true&loc=Asia%2FTokyo"
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}
	return db
}
