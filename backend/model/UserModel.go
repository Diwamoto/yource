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

//User
type User struct {
	AppModel
	Email    string `validate:"required,email"`
	Password string //フロントで弾いてhash化された物が入るイメージ、不正にデータが作られた場合はログインできない為問題ない
	Name     string `validate:"required"`
	Phone    string `validate:"required"`
	Status   bool
	Profiles UserProfile
}

type UserModel struct {
	AppModel
}

//プロフィールを引っ張ってきて返す
func (um UserModel) Join(u *User) {

	up := GetUserProfileByUserId(1)
	//up変数に値が入っていれば追加
	if up.ID > 0 {
		u.Profiles = up
	}

}

//バリデーションをかける
//文字の整形系はフロントで行うので
//最低限の入力チェックのみをgoで行う
func (um *UserModel) Validate(u User) ([]string, bool) {

	validate := validator.New()
	err := validate.Struct(u)
	var messages []string
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.Field()
			switch fieldName {
			case "Email":
				var typ = err.Tag() //バリデーションでNGになったタグ名を取得
				switch typ {
				case "required":
					messages = append(messages, "メールアドレスを入力してください")
				case "email":
					messages = append(messages, "正しいメールアドレスを入力してください")
				}
			case "Name":
				messages = append(messages, "名前を入力してください")
			case "Phone":
				messages = append(messages, "電話番号を入力してください")
			}
		}
	}

	if len(messages) > 0 {
		return messages, true
	} else {
		return []string{}, false
	}

}

//ユーザを作成する
func (um *UserModel) Create(u User) ([]string, bool) {

	var user User
	var db = database.ConnectDB()
	db.AutoMigrate(&user)

	msg, err := um.Validate(u)

	if !err {
		//バリデーションが通れば作成し、メッセージの中に作成したユーザIDを入れて返す
		db.Create(&u)
		msg = append(msg, strconv.Itoa(int(u.Id)))
		db.Close()
		return msg, false
	} else {
		//作成できなければエラーメッセージを返す
		db.Close()
		return msg, true
	}

}

//指定ユーザidの情報を返す
func (um UserModel) GetById(id int) (User, bool) {

	//var ret User
	var u User
	var db = database.ConnectDB()

	db.AutoMigrate(&u)
	db.First(&u, id)
	db.Close()

	//値が取得できたら
	if u.Id == id {
		um.Join(&u)
		return u, false
	} else {
		return User{}, true
	}

}

// //検索インターフェース
// //検索文字列はいったんuser構造体に格納してやりとりする
// func WhereUser(u User) User {
// 	db.AutoMigrate(&u)

// 	//id
// 	db.Where("ID = ?", u.ID)
// 	db.Where("Email = ?", )

// 	return u
// }

//更新メソッド
//ユーザの情報を更新する
func (um UserModel) Update(id int, u User) (User, bool) {
	var tu User
	var db = database.ConnectDB()
	db.AutoMigrate(&tu)
	db.First(&tu, id)

	//引数のユーザの情報を移す
	tu.Email = u.Email
	tu.Name = u.Name
	tu.Password = u.Password
	tu.Phone = u.Phone
	//更新日を現在にする
	tu.Modified = time.Now()

	//バリデーションをかける
	_, err := um.Validate(u)

	//バリデーションが成功していたら
	if !err {
		//セーブした結果がエラーであれば更新失敗
		if result := db.Save(&tu); result.Error != nil {
			db.Close()
			return User{}, false
		} else {
			db.Close()
			return tu, false
		}
	} else {
		//作成できなければエラーメッセージを返す
		db.Close()
		return tu, false
	}

}

//削除メソッド
//ユーザを削除する
func (um *UserModel) Delete(id int) ([]string, bool) {

	var db = database.ConnectDB()
	//idで削除を実行する
	_, err := um.GetById(id)
	if err { //削除するユーザがいなかったらダメ
		return []string{"削除するユーザが存在しません。"}, true
	}
	db.Delete(&User{}, id)
	db.Close()
	_, err2 := um.GetById(id)
	if err2 { //ユーザが取得できなかったら成功
		return []string{"削除に成功しました。"}, false
	} else {
		return []string{"削除できませんでした。"}, true
	}

}
