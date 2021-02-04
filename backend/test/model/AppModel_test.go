package model

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

	//テスト用のユーザを作成

	//ユーザ以外のテストに使用するテストユーザを作成
	mtu := model.User{
		Email:    "master@example.com",
		Password: "4AeNkWVisJ",
		Name:     "master name",
		Phone:    "028-0728-9727",
		Status:   true,
		Profile:  model.UserProfile{},
	}
	mtu.Created = time.Now()
	mtu.Modified = time.Now()
	db.Create(&mtu)

	//ユーザ以外のテストに使用するテストユーザを作成その２
	mtu2 := model.User{
		Email:    "master2@example.com",
		Password: "4AeNkWVisJ",
		Name:     "master name",
		Phone:    "028-0728-9727",
		Status:   true,
		Profile:  model.UserProfile{},
	}
	mtu2.Created = time.Now()
	mtu2.Modified = time.Now()
	db.Create(&mtu2)

	//スペース以外のテストに使用するテストスペースを作成
	mts := model.Space{
		UserId:      1,
		Name:        "master name",
		Discription: "master disc",
		SubDomain:   "master",
		Status:      true,
		Publish:     true,
	}
	mts.Created = time.Now()
	mts.Modified = time.Now()
	db.Create(&mts)

	//チャンネル以外のテストに使用するテストチャンネルを作成
	mtc := model.Channel{
		SpaceId:     1,
		UserId:      1,
		Name:        "master name",
		Discription: "master disc",
	}
	mtc.Created = time.Now()
	mtc.Modified = time.Now()
	db.Create(&mtc)

	//テストを実行
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
