package model

import (

	//標準ライブラリ
	"strconv"

	//自作ライブラリ
	"main/config/database"

	//githubライブラリ
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type User struct {
	Id       int
	Email    string `validate:"email"`
	Password string //hash()
	Name     string
	Phone    string
	Status   bool
	Profiles UserProfile
	Created  string
	Modified string
}

var db = database.ConnectDB()

//プロフィールを引っ張ってきて返す
func JoinUserProfile(u *User) {

	up := GetUserProfileByUserId(1)
	//up変数に値が入っていれば追加
	if up.Id > 0 {
		u.Profiles = up
	}

}

//バリデーションをかける
func ValidateUser(u User) (string, bool) {

	validate := validator.New()
	err := validate.Struct(u)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {

			fieldName := err.Field()
			switch fieldName {
			case "Email":
				return "Please enter a valid email address.", true
			}
		}
	}

	return "", false
}

//ユーザを作成する
func CreateUser(c *gin.Context) (User, bool) {

	var user User
	if c.PostForm("id") != "" {
		user.Id, _ = strconv.Atoi(c.PostForm("id"))
		user.Email = c.PostForm("Email")

		return user, false
	} else {
		return User{}, true
	}

}

//指定ユーザidの情報を返す
func GetUser(id int) User {

	//var ret User
	var u User
	db.AutoMigrate(&u)
	db.First(&u, id)
	db.Close()

	//値が取得できたら
	if u.Id == id {
		JoinUserProfile(&u)

	}

	return u

}
