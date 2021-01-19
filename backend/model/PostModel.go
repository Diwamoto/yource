package model

import (

	"main/config/database"
)

type Post struct {
	Id int
	UserId int
	ChannelId int
	Content string
	Date string
	Status bool
	Created string
	Modified string
}

func GetPosts(id int) Post{

	db := database.ConnectDB()
	var ret Post
	db.First(&ret, id)

	db.Close()
	
	return ret
}
