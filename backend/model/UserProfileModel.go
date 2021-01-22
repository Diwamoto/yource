package model

import (
	//標準ライブラリ
	"strconv"
	"time"

	//自作ライブラリ
	"main/config/database"

	//githubライブラリ
	"github.com/go-playground/validator"
)

type UserProfile struct {
	Entity
	UserId    int
	Profile   string
	Birthday  time.Time
	From      string
	Job       string
	Twitter   string
	Facebook  string
	Instagram string
	Other     string
}

type UserProfileModel struct {
	AppModel
}

func NewUserProfileModel(t string) *UserProfileModel {
	var upm UserProfileModel
	upm.db = database.ConnectDB(t)
	upm.nc = t
	return &upm
}

func (UserProfileModel) TableName() string {
	return "user_profiles"
}

//バリデーションをかける
//文字の整形系はフロントで行うので
//最低限の入力チェックのみをgoで行う
func (upm *UserProfileModel) Validate(up UserProfile) ([]string, bool) {

	validate := validator.New()
	err := validate.Struct(up)
	var messages []string
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.Field()
			switch fieldName {
			//ユーザプロフィールは基本的にnull許容
			}
		}
	}

	//ユーザIDは存在するユーザのIDのみを使用できる
	um := NewUserModel(upm.nc)
	_, err2 := um.GetById(up.UserId)
	if err2 {
		messages = append(messages, "存在しないユーザIDのプロフィールは作成できません。")
	}

	if len(messages) > 0 {
		return messages, true
	} else {
		return []string{}, false
	}

}

//データを作成する
func (upm *UserProfileModel) Create(up UserProfile) ([]string, bool) {

	upm.db.AutoMigrate(&up)

	msg, err := upm.Validate(up)

	if !err {
		up.Created = time.Now()
		up.Modified = time.Now()
		//バリデーションが通れば作成し、メッセージの中に作成したデータのIDを入れて返す
		upm.db.Create(&up)
		msg = append(msg, strconv.Itoa(int(up.Id)))
		return msg, false
	} else {
		//作成できなければエラーメッセージを返す
		return msg, true
	}

}

//指定idのデータがあれば返す
func (upm UserProfileModel) GetById(id int) (UserProfile, bool) {

	var up UserProfile

	upm.db.AutoMigrate(&up)
	upm.db.First(&up, id)

	//値が取得できたら
	if up.Id == id {
		return up, false
	} else {
		return UserProfile{}, true
	}

}

//更新メソッド 情報を更新する
func (upm UserProfileModel) Update(id int, up UserProfile) ([]string, bool) {
	var tup UserProfile
	upm.db.AutoMigrate(&tup)
	upm.db.First(&tup, id)

	//引数の情報を移す
	//ユーザプロフィールはnull許容なのでチェックはなし
	tup.Profile = up.Profile
	tup.Birthday = up.Birthday
	tup.From = up.From
	tup.Job = up.Job
	tup.Twitter = up.Twitter
	tup.Facebook = up.Facebook
	tup.Instagram = up.Instagram
	tup.Other = up.Other
	//更新日を現在にする
	tup.Modified = time.Now()

	//バリデーションをかける
	msg, err := upm.Validate(tup)

	//ユーザIDは変更できない
	if tup.UserId != up.UserId {
		msg = append(msg, "ユーザIDは変更することはできません。")
		err = true
	}

	//バリデーションが成功していたら
	if !err {
		//セーブした結果がエラーであれば更新失敗
		if result := upm.db.Save(&tup); result.Error != nil {
			return []string{"データベースに保存することができませんでした。"}, false
		} else {
			return []string{}, false
		}
	} else {
		//作成できなければエラーメッセージを返す
		return msg, false
	}

}

//削除メソッド データを削除する
func (upm *UserProfileModel) Delete(id int) ([]string, bool) {

	//idで削除を実行する
	_, err := upm.GetById(id)
	if err { //削除するデータがなかったらダメ
		return []string{"削除するユーザプロフィールが存在しません。"}, true
	}
	upm.db.Delete(&UserProfile{}, id)
	_, err2 := upm.GetById(id)
	if err2 { //取得できなかったら成功
		return []string{"削除に成功しました。"}, false
	} else {
		return []string{"削除できませんでした。"}, true
	}

}
