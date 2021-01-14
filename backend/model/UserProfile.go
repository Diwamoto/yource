package model

import (

	"main/database"
)

type UserProfile struct {
	Id int
	UserId int
	Profile string
	Birthday string
	From string
	Job string
	Twitter string
	Facebook string
	Instagram string
	Other string
	Created string
	Updated string
}

func GetUserProfile(id int) UserProfile{

	db := database.ConnectDB()
	var ret UserProfile
	db.First(&ret, id)

	db.Close()
	
	return ret
}

func GetUserProfileByUserId(userId int) UserProfile{

	db := database.ConnectDB()
	var up UserProfile
	db.AutoMigrate(&up)
	db.First(&up, "user_id = ?", userId)

	db.Close()
	
	return up
}