package controller

import (
	//標準ライブラリ

	"main/model"
	"net/http"
	"strconv"

	//自作ライブラリ

	//githubライブラリ
	"github.com/gin-gonic/gin"
)

//ユーザー作成アクション
//POSTされた要素でデータを作成する
func CreateUserAction(c *gin.Context) {

	u := model.User{
		Email:    c.PostForm("Email"),
		Password: c.PostForm("Password"),
		Name:     c.PostForm("Name"),
		Phone:    c.PostForm("Phone"),
		Status:   true,
		Profiles: model.UserProfile{},
	}

	msg, err := model.CreateUser(u)
	//エラーじゃなければuserの情報を返す
	if err == false {
		userID, _ := strconv.Atoi(msg[0])
		user, _ := model.GetUser(userID)
		c.JSON(http.StatusCreated, user)
	} else {
		//作成できなければエラーメッセージを返す。
		c.JSON(http.StatusConflict, msg)

	}
}

//ユーザの情報を返すアクション
func GetUserAction(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	user, _ := model.GetUser(id)

	c.JSON(http.StatusOK, user)
}
