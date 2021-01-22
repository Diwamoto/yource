package test_model

import (
	"main/config/database"
	"main/model"
	"os"
	"testing"
	"time"
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
	//ユーザ以外のテストに使用するテストユーザを作成
	mtestuser := model.User{
		Email:    "master@example.com",
		Password: "4AeNkWVisJ",
		Name:     "master name",
		Phone:    "028-0728-9727",
		Status:   true,
		Profile:  model.UserProfile{},
	}
	mtestuser.Created = time.Now()
	mtestuser.Modified = time.Now()
	db.Create(&mtestuser)
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
