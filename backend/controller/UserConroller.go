package controller

import (
	//標準ライブラリ

	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	//自作ライブラリ
	"main/model"

	//githubライブラリ
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//ユーザー作成アクション
//POSTされた要素でデータを作成する
func CreateUserAction(c *gin.Context) {

	um := model.NewUserModel("default")

	password, _ := bcrypt.GenerateFromPassword([]byte(c.PostForm("Password")), bcrypt.DefaultCost)

	u := model.User{
		Email:    c.PostForm("Email"),
		Password: string(password),
		Name:     c.PostForm("Name"),
		Phone:    c.PostForm("Phone"),
		Status:   false, //メールアドレス認証ができるまでステータスは有効にならない
		Profile:  model.UserProfile{},
	}
	u.Created = time.Now()
	u.Modified = time.Now()

	msg, err := um.Create(u)
	//エラーじゃなければuserの情報を返す
	if !err {
		userId, _ := strconv.Atoi(msg[0])
		user, _ := um.GetById(userId)
		user.Id = userId

		//ユーザのメールアドレス死活監視トークンを生成する。

		c.JSON(http.StatusCreated, user)
	} else {
		//作成できなければエラーメッセージを返す。
		c.JSON(http.StatusConflict, msg)

	}
}

//ユーザの情報を返すアクション
//GETでパラメータのユーザの情報を取得する
func GetUserAction(c *gin.Context) {

	um := model.NewUserModel("default")

	id, _ := strconv.Atoi(c.Param("id"))
	u, err := um.GetById(id)
	if !err {
		c.JSON(http.StatusOK, u)
	} else {
		c.JSON(http.StatusNotFound, []string{})
	}
}

//全てのユーザの情報を返すアクション
//GETでパラメータのユーザの情報を取得する
func GetAllUserAction(c *gin.Context) {

	um := model.NewUserModel("default")

	users, err := um.GetAll()
	if !err {
		c.JSON(http.StatusOK, users)
	} else {
		c.JSON(http.StatusNotFound, []string{})
	}
}

//ユーザの情報を更新するアクション
//PUTでフォームの情報からユーザの情報を更新する
func UpdateUserAction(c *gin.Context) {

	um := model.NewUserModel("default")

	userId, _ := strconv.Atoi(c.Param("id"))
	//ユーザを取得し、取得できたら更新をかける
	_, err := um.GetById(userId)
	if !err {
		//フォームから更新内容を取得したユーザ構造体を作成
		var u model.User
		u.Email = c.PostForm("Email")
		u.Password = c.PostForm("Password")
		u.Name = c.PostForm("Name")
		u.Phone = c.PostForm("Phone")
		Status, _ := strconv.ParseBool(c.PostForm("Status"))
		u.Status = Status
		msg, err2 := um.Update(userId, u)
		if !err2 {
			r, _ := um.GetById(userId)
			c.JSON(http.StatusOK, r)
		} else {
			c.JSON(http.StatusConflict, msg)
		}
	} else {
		c.JSON(http.StatusNotFound, []string{})
	}
}

//ユーザの削除アクション
func DeleteUserAction(c *gin.Context) {

	um := model.NewUserModel("default")
	userId, _ := strconv.Atoi(c.Param("id"))
	msg, err := um.Delete(userId)
	if !err {
		c.JSON(http.StatusOK, msg)
	} else {
		c.JSON(http.StatusConflict, msg)
	}
}

func LoginAction(c *gin.Context) {

	um := model.NewUserModel("default")
	var user model.User
	user.Email = c.PostForm("Email")
	user.Status = true
	users, err := um.Find(user)

	//ユーザを取得でき、且ハッシュ化されたパスワードが等しければログイン成功
	if !err && len(users) == 1 {
		user := users[0]
		err1 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.PostForm("Password")))
		if err1 == nil {

			//ログインに成功したらjwtを発行しクッキーとredisに保存

			//jwtの作成
			// headerのセット
			token := jwt.New(jwt.SigningMethodHS256)

			// claimsのセット
			claims := token.Claims.(jwt.MapClaims)
			claims["admin"] = true
			claims["sub"] = user.Id
			claims["name"] = "taro"
			claims["iat"] = time.Now()
			claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

			// 電子署名
			tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNKEY")))

			//セッションに保存
			session := sessions.Default(c)
			session.Set(tokenString, user)
			session.Save()

			c.SetCookie("token", tokenString, 3600, "/", "localhost", true, true)
			c.SetCookie("userId", fmt.Sprint(user.Id), 3600, "/", "localhost", true, true)

			c.JSON(http.StatusOK, gin.H{"token": tokenString, "id": fmt.Sprint(user.Id)})
		} else {
			c.JSON(http.StatusUnauthorized, "メールアドレスもしくはパスワードが間違っています。")
		}
	} else {
		c.JSON(http.StatusUnauthorized, "メールアドレスもしくはパスワードが間違っています。")
	}

}
