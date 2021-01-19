package model

import (
	"main/config/database"
)

//
type Channel struct {
	ID          int    `json:"id,omitempty"`
	UserID      int    `json:"user_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Discription string `json:"discription,omitempty"`
	Created     string `json:"created,omitempty"`
	Modified    string `json:"modified,omitempty"`
}

//チャンネルを取得する
func GetChannel(id int) Channel {

	db := database.ConnectDB()
	var ret Channel
	db.First(&ret, id)

	db.Close()

	return ret
}
