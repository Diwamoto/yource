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

//投稿作成アクション
//POSTされた要素でデータを作成する
func CreatePostAction(c *gin.Context) {

	pm := model.NewPostModel("default")

	//spaceidをintに変換
	channelId, _ := strconv.Atoi(c.Param("id"))
	//useridをintに変換
	userId, _ := strconv.Atoi(c.PostForm("UserId"))

	s := model.Post{
		Entity:    model.Entity{},
		ChannelId: channelId,
		UserId:    userId,
		Content:   c.PostForm("Content"),
		Date:      time.Now(),
	}
	s.Created = time.Now()
	s.Modified = time.Now()

	msg, err := pm.Create(s)
	//エラーじゃなければuserの情報を返す
	if !err {
		postId, _ := strconv.Atoi(msg[0])
		post, _ := pm.GetById(postId)

		//ユーザのメールアドレス死活監視トークンを生成する。

		c.JSON(http.StatusCreated, post)
	} else {
		//作成できなければエラーメッセージを返す。
		c.JSON(http.StatusConflict, msg)

	}
}

//投稿の情報を返すアクション
//GETで指定IDの投稿を返す
func GetPostAction(c *gin.Context) {

	pm := model.NewPostModel("default")

	id, _ := strconv.Atoi(c.Param("id"))
	post, err := pm.GetById(id)
	if !err {
		c.JSON(http.StatusOK, post)
	} else {
		c.JSON(http.StatusNotFound, []string{})
	}
}

//投稿の情報を返すアクション
//GETで指定投稿の投稿を全て返す
func GetPostByChannelIdAction(c *gin.Context) {

	pm := model.NewPostModel("default")

	channelId, _ := strconv.Atoi(c.Param("id"))
	posts, err := pm.GetByChannelId(channelId)
	if !err {
		c.JSON(http.StatusOK, posts)
	} else {
		c.JSON(http.StatusNotFound, []string{})
	}
}

//投稿の情報を返すアクション
//GETで指定ユーザの投稿を全て返す
func GetPostByUserIdAction(c *gin.Context) {

	pm := model.NewPostModel("default")

	userId, _ := strconv.Atoi(c.Param("id"))
	posts, err := pm.GetByUserId(userId)
	if !err {
		c.JSON(http.StatusOK, posts)
	} else {
		c.JSON(http.StatusNotFound, []string{})
	}
}

//全ての投稿の情報を返すアクション
//GETで全ての投稿の情報を取得する
func GetAllPostAction(c *gin.Context) {

	pm := model.NewPostModel("default")

	posts, err := pm.GetAll()
	if !err {
		c.JSON(http.StatusOK, posts)
	} else {
		c.JSON(http.StatusNotFound, []string{})
	}
}

//投稿の情報を更新するアクション
//PUTでフォームの情報から投稿の情報を更新する
func UpdatePostAction(c *gin.Context) {

	pm := model.NewPostModel("default")

	id, _ := strconv.Atoi(c.Param("id"))
	//を取得し、取得できたら更新をかける
	_, err := pm.GetById(id)
	if !err {
		//フォームから更新内容を取得し投稿構造体を作成
		var post model.Post
		post.Content = c.PostForm("Content")
		msg, err2 := pm.Update(id, post)
		if !err2 {
			r, _ := pm.GetById(id)
			c.JSON(http.StatusOK, r)
		} else {
			c.JSON(http.StatusConflict, msg)
		}
	} else {
		c.JSON(http.StatusNotFound, []string{})
	}
}

//投稿の削除アクション
func DeletePostAction(c *gin.Context) {

	pm := model.NewPostModel("default")
	postId, _ := strconv.Atoi(c.Param("id"))
	msg, err := pm.Delete(postId)
	if !err {
		c.JSON(http.StatusOK, msg)
	} else {
		c.JSON(http.StatusConflict, msg)
	}
}
