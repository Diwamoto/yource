package test

import (
	"main/model"
	"testing"
)

//ユーザバリデーションのテスト
func TestValidateUser(t *testing.T) {

	var validUser model.User
	var inValidUser model.User

	//①: 正しいデータのユーザ
	validUser = model.User{
		Email:    "test@example.com",
		Password: "4AeNkWVisJ",
		Name:     "test name",
		Phone:    "000-0000-0000",
		Status:   true,
		Profiles: model.UserProfile{},
	}
	result, err := model.ValidateUser(validUser)
	//正しいユーザのデータならバリデーションは通るはず
	if err == true {
		t.Error(result)
	}

	//②: メールアドレスがないデータのユーザ
	inValidUser = model.User{
		Email:    "", //メールアドレスが空欄
		Password: "4AeNkWVisJ",
		Name:     "test name",
		Phone:    "000-0000-0000",
		Status:   true,
		Profiles: model.UserProfile{},
	}
	result, err = model.ValidateUser(inValidUser)

	//エラーが出なければテスト失敗
	if err == false {
		t.Error(result)
	}

	//③: メールアドレスがおかしいデータのユーザ
	inValidUser = model.User{
		Email:    "testexample.com", //メールアドレスが正しくない
		Password: "4AeNkWVisJ",
		Name:     "test name",
		Phone:    "000-0000-0000",
		Status:   true,
		Profiles: model.UserProfile{},
	}
	result, err = model.ValidateUser(inValidUser)

	//エラーが出なければテスト失敗
	if err == false {
		t.Error(result)
	}

	//④: 名前が入力されていないユーザ
	inValidUser = model.User{
		Email:    "test@example.com",
		Password: "4AeNkWVisJ",
		Name:     "", //名前が入力されていない
		Phone:    "000-0000-0000",
		Status:   true,
		Profiles: model.UserProfile{},
	}
	result, err = model.ValidateUser(inValidUser)

	//エラーが出なければテスト失敗
	if err == false {
		t.Error(result)
	}

	//⑤: 電話番号が入力されていないユーザ
	inValidUser = model.User{
		Email:    "test@example.com",
		Password: "4AeNkWVisJ",
		Name:     "test name",
		Phone:    "", //電話番号が入力されていない
		Status:   true,
		Profiles: model.UserProfile{},
	}
	result, err = model.ValidateUser(inValidUser)

	//エラーが出なければテスト失敗
	if err == false {
		t.Error(result)
	}

}

func TestCreateUser(t *testing.T) {
	u := model.User{
		Email:    "test@example.com",
		Password: "TestPassword",
		Name:     "test name",
		Phone:    "080-0299-8293",
		Status:   true,
		Profiles: model.UserProfile{},
	}
	msg, err := model.CreateUser(u)

	//ユーザが作られなかったら失敗
	if err == true {
		t.Error(msg)
	}
}

func TestGetUser(t *testing.T) {

	result, err := model.GetUser(1)
	//エラーフラグがtrueなら失敗
	if err == true {
		t.Error("func GetUser() failed.")
	} else {
		t.Logf("ok, userId = %d", result.ID)
	}
}

func TestUpdateUser(t *testing.T) {
}

func TestDeleteUser(t *testing.T) {

}
