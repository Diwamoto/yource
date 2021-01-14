package controller

import (
	//標準ライブラリ
	"net/http"

	//自作ライブラリ
	"main/model"

	//githubライブラリ 
	"github.com/gin-gonic/gin"

)

func Get(c *gin.Context) {

	User := model.GetUser(1)

	c.JSON(http.StatusOK, User)
	c.JSON(http.StatusOK, User.Id)
}
