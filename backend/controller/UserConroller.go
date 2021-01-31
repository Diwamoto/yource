package controller

import (
	//標準ライブラリ

	"main/model"
	"net/http"
	"strconv"
	"time"

	//自作ライブラリ

	//githubライブラリ
	"github.com/gin-gonic/gin"
)

//ユーザー作成アクション
//POSTされた要素でデータを作成する
func CreateUserAction(c *gin.Context) {

	um := model.NewUserModel("default")

	u := model.User{
		Email:    c.PostForm("Email"),
		Password: c.PostForm("Password"),
		Name:     c.PostForm("Name"),
		Phone:    c.PostForm("Phone"),
		Status:   true,
		Profile:  model.UserProfile{},
	}
	u.Created = time.Now()
	u.Modified = time.Now()

	msg, err := um.Create(u)
	//エラーじゃなければuserの情報を返す
	if !err {
		userID, _ := strconv.Atoi(msg[0])
		a, _ := um.GetById(userID)
		a.Id = userID
		c.JSON(http.StatusCreated, a)
	} else {
		//作成できなければエラーメッセージを返す。
		c.JSON(http.StatusConflict, msg)

	}
}

//ユーザの情報を返すアクション
//GETでパラメータのユーザの情報を取得する
func GetUserAction(c *gin.Context) {

	um := model.NewUserModel("")

	id, _ := strconv.Atoi(c.Param("id"))
	user, err := um.GetById(id)
	if !err {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusNotFound, []string{})
	}
}

//全てのユーザの情報を返すアクション
//GETでパラメータのユーザの情報を取得する
func GetAllUserAction(c *gin.Context) {

	um := model.NewUserModel("")

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
		var user model.User
		user.Email = c.PostForm("Email")
		user.Password = c.PostForm("Password")
		user.Name = c.PostForm("Name")
		user.Phone = c.PostForm("Phone")
		Status, _ := strconv.ParseBool(c.PostForm("Status"))
		user.Status = Status
		msg, err2 := um.Update(userId, user)
		if !err2 {
			c.JSON(http.StatusOK, user)
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
