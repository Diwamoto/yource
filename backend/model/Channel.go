package model

import (

	"main/config/database"
)

type Channel struct {
	Id int
	UserId int
	Name string
	Discription string
	Created string
	Modified string
}

func GetChannel(id int) Channel{

	db := database.ConnectDB()
	var ret Channel
	db.First(&ret, id)

	db.Close()
	
	return ret
}