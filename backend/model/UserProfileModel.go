package model

import (
	"main/config/database"
	"strconv"
	"time"

	"github.com/go-playground/validator"
)

type UserProfile struct {
	Entity
	UserID    int
	Profile   string
	Birthday  string
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

//バリデーションをかける
//文字の整形系はフロントで行うので
//最低限の入力チェックのみをgoで行う
func (upm *UserProfileModel) Validate(u UserProfile) ([]string, bool) {

	validate := validator.New()
	err := validate.Struct(u)
	var messages []string
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.Field()
			switch fieldName {
			//ユーザプロフィールは基本的にnull許容
			}
		}
	}

	if len(messages) > 0 {
		return messages, true
	} else {
		return []string{}, false
	}

}

//データを作成する
func (upm *UserProfileModel) Create(u UserProfile) ([]string, bool) {

	upm.db.AutoMigrate(&u)

	msg, err := upm.Validate(u)

	if !err {
		//バリデーションが通れば作成し、メッセージの中に作成したデータのIDを入れて返す
		upm.db.Create(&u)
		msg = append(msg, strconv.Itoa(int(u.Id)))
		upm.db.Close()
		return msg, false
	} else {
		//作成できなければエラーメッセージを返す
		upm.db.Close()
		return msg, true
	}

}

//指定idのデータがあれば返す
func (upm UserProfileModel) GetById(id int) (UserProfile, bool) {

	var up UserProfile

	upm.db.AutoMigrate(&up)
	upm.db.First(&up, id)
	upm.db.Close()

	//値が取得できたら
	if up.Id == id {
		return up, false
	} else {
		return UserProfile{}, true
	}

}

// //検索インターフェース
// //検索文字列はいったんuser構造体に格納してやりとりする
// func WhereUser(u User) User {
// 	upm.db.AutoMigrate(&u)

// 	//id
// 	upm.db.Where("ID = ?", u.ID)
// 	upm.db.Where("Email = ?", )

// 	return u
// }

//更新メソッド 情報を更新する
func (upm UserProfileModel) Update(id int, up UserProfile) (UserProfile, bool) {
	var tup UserProfile
	upm.db.AutoMigrate(&tup)
	upm.db.First(&tup, id)

	//引数の情報を移す
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
	_, err := upm.Validate(tup)

	//バリデーションが成功していたら
	if !err {
		//セーブした結果がエラーであれば更新失敗
		if result := upm.db.Save(&tup); result.Error != nil {
			upm.db.Close()
			return UserProfile{}, false
		} else {
			upm.db.Close()
			return tup, false
		}
	} else {
		//作成できなければエラーメッセージを返す
		upm.db.Close()
		return tup, false
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
	upm.db.Close()
	_, err2 := upm.GetById(id)
	if err2 { //取得できなかったら成功
		return []string{"削除に成功しました。"}, false
	} else {
		return []string{"削除できませんでした。"}, true
	}

}
