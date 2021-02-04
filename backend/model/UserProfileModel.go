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
	UserId    int //MEMO: 本当であればバリデーションを用いたいが、Userの子になっているUserProfileではなぜか独自バリデーションが読み込まれないのでCreate()時に判断する
	Profile   string
	Birthday  time.Time
	From      string //TODO: 出身地は別にテーブルを設けてidで判断する？ → フロントでselectboxで判断すればテーブルを用意する必要はなさそう
	Job       string //TODO: 上に同じ
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
func (upm UserProfileModel) Validate(up UserProfile) ([]string, bool) {

	validate := validator.New()
	_ = validate.Struct(up)
	var messages []string
	//MEMO: 本当であればバリデーションを用いたいが、Userの子になっているUserProfileでは
	//なぜか独自バリデーションが読み込まれないので、Create()の中で判断するように実装する。
	//原因として考えられるのは、バリデーションにデータベースと問い合わせる部分があり、
	// if err != nil {
	// 	for _, err := range err.(validator.ValidationErrors) {
	// 		fieldName := err.Field()
	// 		switch fieldName {
	// 		case "UserId":
	// 			messages = append(messages, "同じユーザIDのプロフィールは追加できません")
	// 		}
	// 	}
	// }

	//プロフィールは存在するユーザのIDのみを使用できる
	um := NewUserModel(upm.nc)

	_, err2 := um.GetById(up.UserId)
	if err2 {
		return []string{"存在しないユーザIDのプロフィールは作成できません。"}, true
	}

	if len(messages) > 0 {
		return messages, true
	} else {
		return []string{}, false
	}

}

//データを作成する
func (upm UserProfileModel) Create(up UserProfile) ([]string, bool) {

	upm.db.AutoMigrate(&up)

	msg, err := upm.Validate(up)

	if !err {

		//MEMO: UserProfileの定義を参照、UserIdの重複チェックを実装する。
		if !upm.ValidateUniqueUserId(up.UserId) {
			//作成できなければエラーメッセージを返す
			msg = append(msg, "既に指定ユーザIdのプロフィールが登録されています。")
			return msg, true
		}

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

//指定プロフィールidの情報を返す
func (upm UserProfileModel) GetAll() ([]UserProfile, bool) {

	var userProfiles []UserProfile
	upm.db.Find(&userProfiles)
	for i := range userProfiles {
		upm.db.Model(userProfiles[i])
	}

	//値が取得できたら
	if len(userProfiles) > 0 {
		return userProfiles, false
	} else {
		return []UserProfile{}, true
	}

}

//指定プロフィールidの情報を返す
func (upm UserProfileModel) GetById(id int) (UserProfile, bool) {

	var up UserProfile
	upm.db.First(&up, id)

	//値が取得できたら
	if up.Id == id {
		return up, false
	} else {
		return UserProfile{}, true
	}

}

//ユーザIDで検索する
//プロフィールは基本ユーザIDで検索する
func (upm UserProfileModel) GetByUserId(userId int) (UserProfile, bool) {
	var up UserProfile
	upm.db.Where("user_id = ?", userId).First(&up)

	//値が取得できたら
	if up.UserId == userId {
		return up, false
	} else {
		return UserProfile{}, true
	}
}

//検索メソッド
//任意の条件に一致するプロフィールを取得する
func (upm UserProfileModel) Find(up UserProfile) ([]UserProfile, bool) {

	var r []UserProfile
	//TODO: like検索の実装
	//→ただ、基本ユーザidで呼び出す為いらないかも？（プロフィールの文言で検索する機会がなさそう）
	upm.db.Where(&up).Find(&r)

	//dbに問い合わせて存在していればプロフィールを返す。なければエラーを返す ←？？
	if len(r) > 0 {
		return r, false
	} else {
		return []UserProfile{}, true
	}
}

//更新メソッド 情報を更新する
//プロフィールの連番は気にせずユーザIDで判断する
func (upm UserProfileModel) Update(userId int, up UserProfile) ([]string, bool) {
	var tup UserProfile
	upm.db.AutoMigrate(&tup)
	tup, err := upm.GetByUserId(userId)

	//存在しなければエラーを返す
	if err {
		return []string{"指定されたユーザが存在しません。"}, true
	}
	//ユーザIDは変更できない
	if tup.UserId != up.UserId {
		return []string{"ユーザIDは変更することはできません。"}, true
	}

	//引数の情報を移す
	//ユーザプロフィールプロフィールはnull許容なのでチェックはなし
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

	//バリデーションが成功していたら
	if !err {
		//セーブした結果がエラーであれば更新失敗
		if result := upm.db.Save(&tup); result.Error != nil {
			return []string{"データベースに保存することができませんでした。"}, true
		} else {
			return []string{}, false
		}
	} else {
		//作成できなければエラーメッセージを返す
		return msg, true
	}

}

//削除メソッド データを削除する
func (upm UserProfileModel) Delete(id int) ([]string, bool) {

	//idで削除を実行する
	_, err := upm.GetById(id)
	if err { //削除するデータがなかったらダメ
		return []string{"削除するプロフィールプロフィールが存在しません。"}, true
	}
	upm.db.Delete(&UserProfile{}, id)
	_, err2 := upm.GetById(id)
	if err2 { //取得できなかったら成功
		return []string{"削除に成功しました。"}, false
	} else {
		return []string{"削除できませんでした。"}, true
	}

}

//func (upm UserProfileModel) ValidateUniqueUserId(fl validator.FieldLevel) bool {
func (upm UserProfileModel) ValidateUniqueUserId(id int) bool {
	// userId := int(fl.Field().Int())
	up := UserProfile{
		UserId: id,
	}
	_, err := upm.Find(up)
	//取得できなかったら1.ユーザが存在しないのかどうかをチェック
	if err {

	}
	return err
}
