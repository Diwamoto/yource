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

//チャンネル作成アクション
//POSTされた要素でデータを作成する
func CreateChannelAction(c *gin.Context) {

	cm := model.NewChannelModel("default")

	//spaceidをintに変換
	spaceId, _ := strconv.Atoi(c.Param("id"))

	s := model.Channel{
		Entity:      model.Entity{},
		SpaceId:     spaceId,
		Name:        c.PostForm("Name"),
		Description: c.PostForm("Description"),
	}
	s.Created = time.Now()
	s.Modified = time.Now()

	msg, err := cm.Create(s)
	//エラーじゃなければチャンネルの情報を返す
	if !err {
		channelId, _ := strconv.Atoi(msg[0])
		channel, _ := cm.GetById(channelId)

		c.JSON(http.StatusCreated, channel)
	} else {
		//作成できなければエラーメッセージを返す。
		c.JSON(http.StatusConflict, msg)

	}
}

//チャンネルの情報を返すアクション
//GETで指定IDのチャンネルを返す
func GetChannelAction(c *gin.Context) {

	cm := model.NewChannelModel("default")

	id, _ := strconv.Atoi(c.Param("id"))
	ch, err := cm.GetById(id)
	if !err {
		c.JSON(http.StatusOK, ch)
	} else {
		c.JSON(http.StatusNotFound, []string{})
	}
}

//チャンネルの情報を返すアクション
//GETで指定スペースのチャンネルを全て返す
func GetChannelBySpaceIdAction(c *gin.Context) {

	cm := model.NewChannelModel("default")

	spaceId, _ := strconv.Atoi(c.Param("id"))
	channels, err := cm.GetBySpaceId(spaceId)
	if !err {
		c.JSON(http.StatusOK, channels)
	} else {
		c.JSON(http.StatusNotFound, []string{})
	}
}

//チャンネルの情報を返すアクション
//GETで指定idのチャンネルを返す
func GetChannelByIdAction(c *gin.Context) {

	cm := model.NewChannelModel("default")

	channelId, _ := strconv.Atoi(c.Param("id"))
	channel, err := cm.GetById(channelId)
	if !err {
		c.JSON(http.StatusOK, channel)
	} else {
		c.JSON(http.StatusNotFound, []string{})
	}
}

//全てのチャンネルの情報を返すアクション
//GETで全てのチャンネルの情報を取得する
func GetAllChannelAction(c *gin.Context) {

	cm := model.NewChannelModel("default")

	users, err := cm.GetAll()
	if !err {
		c.JSON(http.StatusOK, users)
	} else {
		c.JSON(http.StatusNotFound, []string{})
	}
}

//チャンネルの情報を更新するアクション
//PUTでフォームの情報からチャンネルの情報を更新する
func UpdateChannelAction(c *gin.Context) {

	cm := model.NewChannelModel("default")

	id, _ := strconv.Atoi(c.Param("id"))
	//パラメータからチャンネルを取得し、取得できたら更新をかける
	_, err := cm.GetById(id)
	if !err {
		//フォームから更新内容を取得したチャンネル構造体を作成
		var ch model.Channel
		ch.Name = c.PostForm("Name")
		ch.Description = c.PostForm("Description")
		msg, err2 := cm.Update(id, ch)
		if !err2 {
			r, _ := cm.GetById(id)
			c.JSON(http.StatusOK, r)
		} else {
			c.JSON(http.StatusConflict, msg)
		}
	} else {
		c.JSON(http.StatusNotFound, []string{})
	}
}

//チャンネルの削除アクション
func DeleteChannelAction(c *gin.Context) {

	cm := model.NewChannelModel("default")
	channelId, _ := strconv.Atoi(c.Param("id"))
	msg, err := cm.Delete(channelId)
	if !err {
		c.JSON(http.StatusOK, msg)
	} else {
		c.JSON(http.StatusConflict, msg)
	}
}
