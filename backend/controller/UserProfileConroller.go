package controller

import (
	//標準ライブラリ
	"net/http"
	"strconv"
	"time"

	//自作ライブラリ
	"main/model"

	//githubライブラリ
	"github.com/gin-gonic/gin"
)

//ユーザープロフィール作成アクション
//POSTされた要素でデータを作成する
func CreateUserProfileAction(c *gin.Context) {

	upm := model.NewUserProfileModel("default")
	//誕生日を時間型に変換
	birth, _ := time.Parse("2006/01/02 15:04:05", c.PostForm("Birthday"))
	//useridをintに変換
	userId, _ := strconv.Atoi(c.Param("id"))

	up := model.UserProfile{
		UserId:    userId,
		Profile:   c.PostForm("Profile"),
		Birthday:  birth,
		From:      c.PostForm("From"),
		Job:       c.PostForm("Job"),
		Twitter:   c.PostForm("Twitter"),
		Facebook:  c.PostForm("Facebook"),
		Instagram: c.PostForm("Instagram"),
		Other:     c.PostForm("Other"),
	}
	up.Created = time.Now()
	up.Modified = time.Now()

	msg, err := upm.Create(up)
	//エラーじゃなければユーザプロフィールの情報を返す
	if !err {
		userID, _ := strconv.Atoi(msg[0])
		a, _ := upm.GetById(userID)
		a.Id = userID

		//ユーザのメールアドレス死活監視トークンを生成する。

		c.JSON(http.StatusCreated, a)
	} else {
		//作成できなければエラーメッセージを返す。
		c.JSON(http.StatusConflict, msg)

	}
}

//ユーザの情報を返すアクション
//GETでパラメータのユーザの情報を取得する
func GetUserProfileAction(c *gin.Context) {

	upm := model.NewUserProfileModel("")

	id, _ := strconv.Atoi(c.Param("id"))
	up, err := upm.GetById(id)
	if !err {
		c.JSON(http.StatusOK, up)
	} else {
		c.JSON(http.StatusNotFound, []string{})
	}
}

//全てのユーザの情報を返すアクション
//GETでパラメータのユーザの情報を取得する
func GetAllUserProfileAction(c *gin.Context) {

	um := model.NewUserProfileModel("")

	users, err := um.GetAll()
	if !err {
		c.JSON(http.StatusOK, users)
	} else {
		c.JSON(http.StatusNotFound, []string{})
	}
}

//ユーザのプロフィールを更新するアクション
//PUTでフォームの情報からユーザの情報を更新する
func UpdateUserProfileAction(c *gin.Context) {

	upm := model.NewUserProfileModel("default")

	userId, _ := strconv.Atoi(c.Param("Id"))
	//ユーザを取得し、取得できたら更新をかける
	_, err := upm.GetByUserId(userId)
	if !err {
		//フォームから更新内容を取得したユーザ構造体を作成
		var up model.UserProfile

		//誕生日を時間型に変換
		birth, _ := time.Parse("2006/01/02 15:04:05", c.PostForm("Birthday"))
		up.Profile = c.PostForm("Profile")
		up.Birthday = birth
		up.From = c.PostForm("From")
		up.Job = c.PostForm("Job")
		up.Twitter = c.PostForm("Twitter")
		up.Facebook = c.PostForm("Facebook")
		up.Instagram = c.PostForm("Instagram")
		up.Other = c.PostForm("Other")

		msg, err2 := upm.Update(userId, up)
		if !err2 {
			up, _ = upm.GetByUserId(userId)
			c.JSON(http.StatusOK, up)
		} else {
			c.JSON(http.StatusConflict, msg)
		}
	} else {
		c.JSON(http.StatusNotFound, []string{})
	}
}

//ユーザの削除アクション
func DeleteUserProfileAction(c *gin.Context) {

	upm := model.NewUserProfileModel("default")
	upId, _ := strconv.Atoi(c.Param("Id"))
	msg, err := upm.Delete(upId)
	if !err {
		c.JSON(http.StatusOK, msg)
	} else {
		c.JSON(http.StatusConflict, msg)
	}
}
