package test

import (
	"main/model"
	"strconv"
	"testing"
	"time"
)

///かなーーーりブラックボックスなテストになっているのであとで直す

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

	//テストユーザ
	u := model.User{
		Email:    "CreateTest@example.com",
		Password: "CrtTestPsw",
		Name:     "Crt Test",
		Phone:    "029-8475-1109",
		Status:   true,
		Profiles: model.UserProfile{},
	}
	u.Created = time.Now()
	u.Modified = time.Now()

	msg, err := model.CreateUser(u)

	//ユーザが作られなかったら失敗
	if err {
		t.Error(msg)
	}
	userID, _ := strconv.Atoi(msg[0])

	//作成したユーザを削除
	_, _ = model.DeleteUser(userID - 1)

}

//GetUser()のテスト
//ユーザが取得できたらOK,できなければダメ
func TestGetUser(t *testing.T) {

	_, err := model.GetUser(1)
	//エラーフラグがtrueなら失敗
	if err {
		t.Error("ユーザを取得できませんでした。")
	}
}

//UpdateUser()のテスト
//ユーザの情報が更新できなかったらダメ
func TestUpdateUser(t *testing.T) {

	//テストユーザを作成
	//ユーザの作成はブラックボックス
	u := model.User{
		Email:    "Upd@example.com",
		Password: "UpdTestPsw",
		Name:     "Upd Test",
		Phone:    "048-8476-8173",
		Status:   true,
		Profiles: model.UserProfile{},
	}
	u.Created = time.Now()
	u.Modified = time.Now()
	msg, _ := model.CreateUser(u)
	userID, _ := strconv.Atoi(msg[0])

	//ユーザの情報をアップデート
	testu := model.User{
		Email:    "Upd2@example.com",
		Password: "UpdTestPsw2",
		Name:     "Upd Test2",
		Phone:    "087-9898-0283",
		Status:   true,
		Profiles: model.UserProfile{},
	}
	model.UpdateUser(userID, testu)

	//ユーザの情報を改めて取得し、変更点が反映されていればOKとする。
	//一つでも変更できていなければ失敗
	aftu, _ := model.GetUser(userID)
	if aftu.Email != testu.Email {
		t.Error("メールアドレスを変更することができませんでした。")
	}
	if aftu.Password != testu.Password {
		t.Error("パスワードを変更することができませんでした。")
	}
	if aftu.Name != testu.Name {
		t.Error("ユーザ名を変更することができませんでした。")
	}
	if aftu.Phone != testu.Phone {
		t.Error("電話番号を変更することができませんでした。")
	}

}

func TestDeleteUser(t *testing.T) {

	//テストユーザを作成
	//ユーザの作成はブラックボックス
	u := model.User{
		Email:    "DeleteTest@example.com",
		Password: "DelTestPsw",
		Name:     "Del Test",
		Phone:    "010-0293-4739",
		Status:   true,
		Profiles: model.UserProfile{},
	}
	u.Created = time.Now()
	u.Modified = time.Now()
	msg, _ := model.CreateUser(u)
	userID, _ := strconv.Atoi(msg[0])

	//作成したユーザを削除
	msg2, err := model.DeleteUser(userID - 1)
	if err {
		//削除失敗
		t.Errorf("ユーザを削除できませんでした。userID = %d %s", userID, msg2)
	}
}
