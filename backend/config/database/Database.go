package database

import (
	"main/config"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

//データベースはモデルの数にかかわらず一つになるようにする（dbに接続しようとするとDB側からエラーが帰ってくる）
//↓自分なりのシングルトン実装
var dataBase *gorm.DB
var singleton_flg = false

func GetInstance(t string) *gorm.DB {
	if !singleton_flg {
		singleton_flg = true
		dataBase = ConnectDB(t)
		if config.Get("debugDB") == true {
			dataBase = dataBase.Debug()
		}
	}

	return dataBase
}

func ConnectDB(t string) *gorm.DB {

	//環境変数を読み込む
	err := godotenv.Load(os.Getenv("ENV_PATH"))
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
