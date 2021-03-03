package model

import (

	//標準ライブラリ

	"errors"
	"time"

	//自作ライブラリ

	"main/config/database"
	//githubライブラリ
)

//UserTokenEntity　Entityを埋め込まれている
type UserToken struct {
	Entity
	UserId int
	Token  string
	Expire time.Time
}

//呼び出し用ユーザトークンモデル
//AppModelを埋め込み
type UserTokenModel struct {
	AppModel
}

func NewUserTokenModel(t string) *UserTokenModel {
	var utm UserTokenModel
	utm.db = database.GetInstance(t)
	utm.nc = t
	return &utm
}

func (utm UserTokenModel) TableName() string {
	return "user_tokens"
}

//作成メソッド
//トークンを作成する
func (utm UserTokenModel) Create(ut UserToken) (UserToken, error) {

	utm.db.AutoMigrate(&ut)

	um := NewUserModel(utm.nc)

	//データベースから取得できなければダメ
	users, err := um.GetById(ut.UserId)
	if len(users) == 0 {
		return UserToken{}, errors.New("ユーザが存在しません。")
	}
	if err != nil {
		return UserToken{}, err
	}

	ut.Created = time.Now()
	ut.Modified = time.Now()
	//バリデーションが通れば作成し、メッセージの中に作成したユーザIDを入れて返す
	utm.db.Create(&ut)
	return ut, nil

}

//トークン文字列でデータベースを検索する
func (utm UserTokenModel) GetByToken(token string) ([]UserToken, error) {

	var r []UserToken

	//dbに問い合わせる。何らかのエラーが発生した場合はここでハンドリング
	if result := utm.db.Model(&UserToken{}).Where("token = ?", token).Find(&r); result.Error != nil {
		//何らかのエラーを返す
		return []UserToken{}, result.Error
	} else {

		if len(r) == 0 {
			return []UserToken{}, errors.New("トークンが見つかりません。")
		} else {
			return r, nil
		}

	}
}

//トークンが有効であるかを調査する
func (utm UserTokenModel) IsValid(token string) (User, error) {

	um := NewUserModel(utm.nc)

	userTokens, err := utm.GetByToken(token)
	//エラーが出ればダメ
	if err != nil {
		return User{}, err
	}

	ut := userTokens[0]
	//有効期限がすぎていてもダメ
	//→現在時刻より有効期限が先じゃないとダメ
	if !ut.Expire.After(time.Now()) {
		//有効期限が切れている場合はそのトークンとトークンに紐づいたユーザを削除する
		if utm.nc == "default" {
			utm.Delete(token)
			um.Delete(ut.UserId)
		}

		return User{}, errors.New("有効期限が切れています。")
	}

	//ユーザの情報を入れて返す
	user, _ := um.GetById(ut.UserId)

	return user[0], nil
}

//削除メソッド
//ユーザを削除する
func (utm UserTokenModel) Delete(token string) error {

	//idで削除を実行する
	userTokens, err := utm.GetByToken(token)
	if err != nil || len(userTokens) == 0 { //削除するトークンがなければダメ
		return errors.New("削除するトークンが存在しません。")
	}
	userToken := userTokens[0]
	utm.db.Delete(&UserToken{}, userToken.Id)
	return nil

}
