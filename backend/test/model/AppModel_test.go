package test_model

import (
	"main/config/database"
	"main/model"
	"os"
	"testing"
)

//テストメイン関数
//全てのテストはここから呼ばれる
func TestMain(m *testing.M) {

	//まずデータベース接続オブジェクトを作成する
	db := database.GetInstance("test")
	//テストに必要なテーブルを全て作成する
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.UserProfile{})
	db.AutoMigrate(&model.Space{})
	db.AutoMigrate(&model.Channel{})
	db.AutoMigrate(&model.Post{})

	code := m.Run()

	//テスト用のデータベースの全てのテーブルを破棄
	db.DropTable(&model.User{})
	db.DropTable(&model.UserProfile{})
	db.DropTable(&model.Space{})
	db.DropTable(&model.Channel{})
	db.DropTable(&model.Post{})
	db.Close()
	os.Exit(code)
}
