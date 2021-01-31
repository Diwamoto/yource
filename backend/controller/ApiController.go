package controller

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func CheckApiKey(c *gin.Context) bool {

	err := godotenv.Load(os.Getenv("ENV_PATH"))
	if err != nil {
		panic(err.Error())
	}

	if c.Request.Header.Get("Apikey") == os.Getenv("APIKEY") {
		return true
	} else {
		if c.Request.Header.Get("Apikey") == "" {
			c.JSON(http.StatusBadRequest, "Auth failed: Apikey not found")
		} else {
			c.JSON(http.StatusBadRequest, "Auth failed: Invalid Apikey")
		}
		return false
	}

}
