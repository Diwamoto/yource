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
		Icon:      c.PostForm("Icon"),
		Birthday:  birth,
		Hometown:  c.PostForm("Hometown"),
		Job:       c.PostForm("Job"),
		Twitter:   c.PostForm("Twitter"),
		Facebook:  c.PostForm("Facebook"),
		Instagram: c.PostForm("Instagram"),
		Other:     c.PostForm("Other"),
	}
	up.Created = time.Now()
	up.Modified = time.Now()

	userProfile, err := upm.Create(up)
	//エラーじゃなければユーザプロフィールの情報を返す
	if err == nil && userProfile.Id != 0 {

		c.JSON(http.StatusCreated, userProfile)
	} else {
		//作成できなければエラーメッセージを返す。
		c.JSON(http.StatusConflict, gin.H{"message": err.Error()})

	}
}

//検索用アクション
//パラメータを解析して、検索用のオブジェクトに挿入してモデルにて検索する
func SearchUserProfileAction(c *gin.Context) {

	upm := model.NewUserProfileModel("default")

	//クエリより検索文字列を取得して、構造体に入れる
	userId, _ := strconv.Atoi(c.Query("Id"))
	u := model.UserProfile{
		UserId:    userId,
		Profile:   c.Query("Profile"),
		Birthday:  time.Time{},
		Hometown:  c.Query("Hometown"),
		Job:       c.Query("Job"),
		Twitter:   c.Query("Twitter"),
		Facebook:  c.Query("Facebook"),
		Instagram: c.Query("Instagram"),
		Other:     c.Query("Other"),
	}
	userProfiles, err := upm.Find(u)
	//検索した結果が0件でもエラーにはならない。
	//検索した条件が間違えていればエラーに入る
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	} else {
		if len(userProfiles) > 0 {
			c.JSON(http.StatusOK, userProfiles)
		} else {
			c.JSON(http.StatusNotFound, gin.H{})
		}

	}

}

//ユーザのプロフィールを更新するアクション
//PUTでフォームの情報からユーザの情報を更新する
func UpdateUserProfileAction(c *gin.Context) {

	upm := model.NewUserProfileModel("default")

	userId, _ := strconv.Atoi(c.Param("id"))

	//ユーザを取得し、取得できたら更新をかける
	userProfile, err := upm.GetByUserId(userId)
	if err == nil && userProfile.UserId == userId {
		//フォームから更新内容を取得したユーザ構造体を作成
		var up model.UserProfile
		up.UserId = userId

		//誕生日を時間型に変換
		birth, _ := time.Parse("2006/01/02 15:04:05", c.PostForm("Birthday"))
		up.Profile = c.PostForm("Profile")
		up.Icon = c.PostForm("Icon")
		up.Birthday = birth
		up.Hometown = c.PostForm("Hometown")
		up.Job = c.PostForm("Job")
		up.Twitter = c.PostForm("Twitter")
		up.Facebook = c.PostForm("Facebook")
		up.Instagram = c.PostForm("Instagram")
		up.Other = c.PostForm("Other")

		UserProfile, err2 := upm.Update(userId, up)
		if err2 == nil && UserProfile.UserId == userId {
			up, _ = upm.GetByUserId(userId)
			c.JSON(http.StatusOK, up)
		} else {
			c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{})
	}
}

//ユーザの削除アクション
func DeleteUserProfileAction(c *gin.Context) {

	upm := model.NewUserProfileModel("default")
	upId, _ := strconv.Atoi(c.Param("id"))
	err := upm.Delete(upId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": "削除に成功しました。"})
	} else {
		c.JSON(http.StatusConflict, gin.H{"message": err.Error()})
	}
}
