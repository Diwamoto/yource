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

//スペース作成アクション
//POSTされた要素でデータを作成する
func CreateSpaceAction(c *gin.Context) {

	sm := model.NewSpaceModel("default")

	//useridをintに変換
	userId, _ := strconv.Atoi(c.PostForm("id"))

	s := model.Space{
		UserId:      userId,
		Name:        c.PostForm("Name"),
		Description: c.PostForm("Description"),
		SubDomain:   c.PostForm("SubDomain"),
		Status:      true,
		Publish:     false, //最初は非公開
	}
	s.Created = time.Now()
	s.Modified = time.Now()

	msg, err := sm.Create(s)
	//エラーじゃなければスペースの情報を返す
	if !err {
		spaceId, _ := strconv.Atoi(msg[0])
		space, _ := sm.GetById(spaceId)

		c.JSON(http.StatusCreated, space)
	} else {
		//作成できなければエラーメッセージを返す。
		c.JSON(http.StatusConflict, msg)

	}
}

//スペースの情報を返すアクション
//GETでパラメータのスペースの情報を取得する
func GetSpaceByIdAction(c *gin.Context) {

	sm := model.NewSpaceModel("default")

	id, _ := strconv.Atoi(c.Param("id"))
	space, err := sm.GetById(id)
	if !err {
		c.JSON(http.StatusOK, space)
	} else {
		c.JSON(http.StatusNotFound, []string{})
	}
}

//スペースの情報を返すアクション
//GETでパラメータのスペースの情報を取得する
func GetSpaceByUserIdAction(c *gin.Context) {

	sm := model.NewSpaceModel("default")

	id, _ := strconv.Atoi(c.Param("id"))
	u, err := sm.GetByUserId(id)
	if !err {
		c.JSON(http.StatusOK, u)
	} else {
		c.JSON(http.StatusNotFound, []string{})
	}
}

//全てのスペースの情報を返すアクション
//GETでパラメータのスペースの情報を取得する
func GetAllSpaceAction(c *gin.Context) {

	sm := model.NewSpaceModel("default")

	users, err := sm.GetAll()
	if !err {
		c.JSON(http.StatusOK, users)
	} else {
		c.JSON(http.StatusNotFound, []string{})
	}
}

//スペースの情報を更新するアクション
//PUTでフォームの情報からスペースの情報を更新する
func UpdateSpaceAction(c *gin.Context) {

	sm := model.NewSpaceModel("default")

	userId, _ := strconv.Atoi(c.Param("id"))
	//スペースを取得し、取得できたら更新をかける
	_, err := sm.GetByUserId(userId)
	if !err {
		//フォームから更新内容を取得したスペース構造体を作成
		var s model.Space
		s.UserId = userId
		s.Name = c.PostForm("Name")
		s.Description = c.PostForm("Description")
		s.SubDomain = c.PostForm("SubDomain")
		Status, _ := strconv.ParseBool(c.PostForm("Status"))
		s.Status = Status
		Publish, _ := strconv.ParseBool(c.PostForm("Publish"))
		s.Publish = Publish
		msg, err2 := sm.Update(userId, s)
		if !err2 {
			r, _ := sm.GetByUserId(userId)
			c.JSON(http.StatusOK, r)
		} else {
			c.JSON(http.StatusConflict, msg)
		}
	} else {
		c.JSON(http.StatusNotFound, []string{})
	}
}

//スペースの削除アクション
func DeleteSpaceAction(c *gin.Context) {

	sm := model.NewSpaceModel("default")
	userId, _ := strconv.Atoi(c.Param("id"))
	msg, err := sm.Delete(userId)
	if !err {
		c.JSON(http.StatusOK, msg)
	} else {
		c.JSON(http.StatusConflict, msg)
	}
}
