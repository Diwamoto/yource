package model

import (
	//標準ライブラリ

	"errors"
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
	Icon      string //CDN先のurlが入る想定
	Birthday  time.Time
	Hometown  string //TODO: 出身地は別にテーブルを設けてidで判断する？ → フロントでselectboxで判断すればテーブルを用意する必要はなさそう
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

func (upm UserProfileModel) TableName() string {
	return "user_profiles"
}

//バリデーションをかける
//文字の整形系はフロントで行うので
//最低限の入力チェックのみをgoで行う
//返す文字列はエラーの文字配列をjson型にしたもの
func (upm UserProfileModel) Validate(up UserProfile) (string, bool) {

	validate := validator.New()
	_ = validate.Struct(up)
	// var messages []string
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

	users, _ := um.GetById(up.UserId)
	if len(users) == 0 {
		return "存在しないユーザIDのプロフィールは作成できません。", true
	} else {
		return "", false
	}
	// if len(messages) > 0 {
	// 	return messages, true
	// } else {
	// 	return []string{}, false
	// }

}

//データを作成する
func (upm UserProfileModel) Create(up UserProfile) (UserProfile, error) {

	upm.db.AutoMigrate(&up)

	msg, err := upm.Validate(up)

	if !err {

		//MEMO: UserProfileの定義を参照、UserIdの重複チェックを実装する。
		if !upm.validateUniqueUserId(up.UserId) {
			return UserProfile{}, errors.New("既に指定ユーザIdのプロフィールが登録されています。")
		}

		up.Created = time.Now()
		up.Modified = time.Now()
		//バリデーションが通れば作成し、メッセージの中に作成したデータのIDを入れて返す
		upm.db.Create(&up)
		return up, nil
	} else {
		//作成できなければエラーメッセージを返す
		return UserProfile{}, errors.New(msg)
	}

}

//指定プロフィールidの情報を返す
//id検索findのラッパー
func (upm UserProfileModel) GetById(id int) (UserProfile, error) {

	var up UserProfile
	up.Id = id
	result, err := upm.Find(up)
	if err != nil {
		return UserProfile{}, err
	} else {
		if len(result) > 0 {
			return result[0], nil
		} else {
			return UserProfile{}, errors.New("指定ユーザIDのプロフィールは存在しません。")
		}

	}
}

//ユーザIDで検索する
//プロフィールは基本ユーザIDで検索する
//ユーザID検索findのラッパー
func (upm UserProfileModel) GetByUserId(userId int) (UserProfile, error) {

	var up UserProfile
	up.UserId = userId
	result, err := upm.Find(up)
	if err != nil {
		return UserProfile{}, err
	} else {
		if len(result) > 0 {
			return result[0], nil
		} else {
			return UserProfile{}, errors.New("指定ユーザIDのプロフィールは存在しません。")
		}

	}

}

//検索メソッド
func (um UserProfileModel) Find(up UserProfile) ([]UserProfile, error) {

	var r []UserProfile
	//受け取った検索パラメータに応じてSQLをビルドする

	builder := um.db.Model(&UserProfile{})

	//検索パラメータを解析し、検索条件が存在していればwhere文を追加する

	if up.Id != 0 {
		builder = builder.Where("id = ?", up.Id)
	}
	if up.Profile != "" {
		builder = builder.Where("profile LIKE ?", "%"+up.Profile+"%")
	}
	var nilTime time.Time
	if up.Birthday != nilTime { //入力されていない設定
		builder = builder.Where("birthday = ?", up.Birthday.Format("2006-01-02 03:04:05"))
	}
	if up.UserId != 0 {
		builder = builder.Where("user_id = ?", up.UserId)
	}
	if up.Hometown != "" {
		builder = builder.Where("hometown LIKE ?", "%"+up.Hometown+"%")
	}
	if up.Job != "" {
		builder = builder.Where("job = ?", up.Job)
	}
	if up.Twitter != "" {
		builder = builder.Where("twitter LIKE ?", "%"+up.Twitter+"%")
	}
	if up.Facebook != "" {
		builder = builder.Where("facebook LIKE ?", "%"+up.Facebook+"%")
	}
	if up.Instagram != "" {
		builder = builder.Where("instagram LIKE ?", "%"+up.Instagram+"%")
	}
	if up.Other != "" {
		builder = builder.Where("other LIKE ?", "%"+up.Other+"%")
	}

	//dbに問い合わせる。何らかのエラーが発生した場合はここでハンドリング
	if result := builder.Find(&r); result.Error != nil {
		//何らかのエラーを返す
		//テストで発生させることができずカバレッジが取れない。。。
		return []UserProfile{}, result.Error
	} else {
		return r, nil
	}
}

//更新メソッド 情報を更新する
//プロフィールの連番は気にせずユーザIDで判断する
func (upm UserProfileModel) Update(userId int, up UserProfile) (UserProfile, error) {

	var tup UserProfile
	upm.db.AutoMigrate(&tup)

	um := NewUserModel(upm.nc)
	users, _ := um.GetById(userId)

	//存在しなければエラーを返す
	if len(users) == 0 {
		return UserProfile{}, errors.New("指定されたユーザが存在しません。")
	}

	tup, _ = upm.GetByUserId(userId)

	//ユーザIDは変更できない
	if tup.UserId != up.UserId {
		return UserProfile{}, errors.New("ユーザIDは変更することはできません。")
	}

	//引数の情報を移す
	//ユーザプロフィールプロフィールはnull許容なのでチェックはなし
	tup.Profile = up.Profile
	tup.Icon = up.Icon
	tup.Birthday = up.Birthday
	tup.Hometown = up.Hometown
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
			return UserProfile{}, result.Error
		} else {
			return tup, nil
		}
	} else {
		//作成できなければエラーメッセージを返す
		return UserProfile{}, errors.New(msg)
	}

}

//削除メソッド データを削除する
func (upm UserProfileModel) Delete(id int) error {

	//idで削除を実行する
	up, _ := upm.GetById(id)
	if up.Id != id { //削除するデータがなかったらダメ
		return errors.New("削除するプロフィールが存在しません。")
	}
	upm.db.Delete(&UserProfile{}, id)
	return nil
	// _, err2 := upm.GetById(id)
	// if err2 { //取得できなかったら成功
	// 	return []string{"削除に成功しました。"}, false
	// } else {
	// 	return []string{"削除できませんでした。"}, true
	// }

}

//func (upm UserProfileModel) ValidateUniqueUserId(fl validator.FieldLevel) bool {
func (upm UserProfileModel) validateUniqueUserId(id int) bool {

	up, _ := upm.GetByUserId(id)

	//まだ持っていなければ
	if up.Id == 0 {
		return true
	} else {
		return false
	}

	// // userId := int(fl.Field().Int())
	// up := UserProfile{
	// 	UserId: id,
	// }
	// _, err := upm.Find(up)
	// //取得できなかったら1.ユーザが存在しないのかどうかをチェック
	// if err {

	// }
	// return err
}
