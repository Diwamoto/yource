package test

import (
	"main/model"
	"testing"
)

func TestValidateUser(t *testing.T) {

	//①: 正しいデータのユーザ
	validUser := model.User{
		Id:       1,
		Email:    "test@example.com",
		Password: "4AeNkWVisJ",
		Name:     "test name",
		Phone:    "000-0000-0000",
		Status:   true,
		Profiles: model.UserProfile{},
		Created:  "2021-01-01 00:00:00",
		Modified: "2021-01-01 01:00:00",
	}
	result, err := model.ValidateUser(validUser)

	if err == true {
		t.Error(result)
	}

	//②: メールアドレスがおかしいデータのユーザ
	inValidUser := model.User{
		Id:       1,
		Email:    "testexample.com", //メールアドレスが正しくない
		Password: "4AeNkWVisJ",
		Name:     "test name",
		Phone:    "000-0000-0000",
		Status:   true,
		Profiles: model.UserProfile{},
		Created:  "2021-01-01 00:00:00",
		Modified: "2021-01-01 01:00:00",
	}
	result, err = model.ValidateUser(inValidUser)

	if err == true {
		t.Error(result)
	}

}

func TestGetUser(t *testing.T) {

	result := model.GetUser(1)
	//ユーザidに0が入っているかつemailが空欄であればnullであると判定
	if result.Id == 0 {
		t.Error("func GetUser() failed.")
	} else {
		t.Logf("ok, userId = %d", result.Id)
	}
}
