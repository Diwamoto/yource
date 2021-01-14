package model

import (

	"main/database"

)

type User struct {
	Id int 
	Email string 
	Name string 
	Phone string 
	Status bool 
	Profiles UserProfile 
	Created string 
	Updated string 
}

//プロフィールを引っ張ってきて返す
func JoinUserProfile(u *User) {


	up := GetUserProfileByUserId(1)
	//up変数に値が入っていれば追加
	if up.Id > 0 {
		u.Profiles = up
	}

}

func GetUser(id int) User{

	db := database.ConnectDB()
	//var ret User
	var u User
	db.AutoMigrate(&u)
	db.First(&u, id)
	db.Close()

	//値が取得できたら
	if u.Id == id {
		JoinUserProfile(&u)

	}

	return u

}