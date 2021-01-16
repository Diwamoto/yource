package model

import (

	//標準ライブラリ
	"fmt"
	"strconv"

	//自作ライブラリ
	"main/config/database"

	//githubライブラリ

	"github.com/go-playground/validator"
	"github.com/jinzhu/gorm"
)

//User
type User struct {
	gorm.Model
	Email    string `validate:"required,email"`
	Password string //フロントで弾いてhash化された物が入るイメージ、不正にデータが作られた場合はログインできない為問題ない
	Name     string `validate:"required"`
	Phone    string `validate:"required"`
	Status   bool
	Profiles UserProfile
}

var db = database.ConnectDB()

//プロフィールを引っ張ってきて返す
func JoinUserProfile(u *User) {

	up := GetUserProfileByUserId(1)
	//up変数に値が入っていれば追加
	if up.ID > 0 {
		u.Profiles = up
	}

}

//バリデーションをかける
//文字の整形系はフロントで行うので
//最低限の入力チェックのみをgoで行う
func ValidateUser(u User) ([]string, bool) {

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
func CreateUser(u User) ([]string, bool) {

	var user User
	db.AutoMigrate(&user)

	msg, err := ValidateUser(u)

	if err == false {
		//バリデーションが通れば作成し、メッセージの中に作成したユーザIDを入れて返す

		db.Create(&u)
		fmt.Println(u.ID)
		msg = append(msg, strconv.Itoa(int(u.ID)))
		return msg, false
	} else {
		//作成できなければエラーメッセージを返す
		return msg, err
	}

}

//指定ユーザidの情報を返す
func GetUser(id int) (User, bool) {

	//var ret User
	var u User
	db.AutoMigrate(&u)
	db.First(&u, id)
	db.Close()

	//値が取得できたら
	if u.ID == uint(id) {
		JoinUserProfile(&u)
		return u, true
	} else {
		return User{}, false
	}

}

func UpdateUser(u User) User {

	return u
}

func DeleteUser(id int) bool {
	u, err := GetUser(id)
	if err == false {
		db.Delete(&u)
		return true
	} else {
		return false
	}
}
