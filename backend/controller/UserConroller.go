package controller

import (
	//標準ライブラリ
	"net/http"
	"strconv"

	//自作ライブラリ
	"main/model"

	//githubライブラリ
	"github.com/gin-gonic/gin"
)

func CreateUserAction(c *gin.Context) {
	user, err := model.CreateUser(c)
	if err == true {
		c.JSON(http.StatusCreated, user)
	} else {
		c.JSON(http.StatusConflict, user)
	}
}

//@param id User.Id
func GetUserAction(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))
	user := model.GetUser(id)

	c.JSON(http.StatusOK, user)
	c.JSON(http.StatusOK, user.Id)
}
