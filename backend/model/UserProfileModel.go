package model

import (
	"main/config/database"
)

type UserProfile struct {
	ID        int
	UserID    int
	Profile   string
	Birthday  string
	From      string
	Job       string
	Twitter   string
	Facebook  string
	Instagram string
	Other     string
	Created   string
	Modified  string
}

func GetUserProfile(id int) UserProfile {

	db := database.ConnectDB()
	var ret UserProfile
	db.First(&ret, id)

	db.Close()

	return ret
}

func GetUserProfileByUserId(userId int) UserProfile {

	db := database.ConnectDB()
	var up UserProfile
	db.AutoMigrate(&up)
	db.First(&up, "user_id = ?", userId)

	db.Close()

	return up
}
