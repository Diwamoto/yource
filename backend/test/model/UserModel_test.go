package test_model

import (
	"main/model"
	"testing"
)

var um = model.NewUserModel("test")

//ValidateUser()のテスト
func TestValidateUser(t *testing.T) {

	tests := []struct {
		in   model.User
		want bool
	}{
		{
			//①正しいユーザ
			model.User{
				Email:    "test@example.com",
				Password: "4AeNkWVisJ",
				Name:     "test name",
				Phone:    "000-0000-0000",
				Status:   true,
				Profile:  model.UserProfile{},
			},
			false, //エラーはでないはず
		},
		{
			//②: メールアドレスがないデータのユーザ
			model.User{
				Email:    "", //メールアドレスが空欄
				Password: "4AeNkWVisJ",
				Name:     "test name",
				Phone:    "000-0000-0000",
				Status:   true,
				Profile:  model.UserProfile{},
			},
			true, //エラーになるはず
		},
		{
			//③: メールアドレスがおかしいデータのユーザ
			model.User{
				Email:    "testexample.com", //メールアドレスが正しくない
				Password: "4AeNkWVisJ",
				Name:     "test name",
				Phone:    "000-0000-0000",
				Status:   true,
				Profile:  model.UserProfile{},
			},
			true, //エラーになるはず
		},
		{
			//④: 名前が入力されていないユーザ
			model.User{
				Email:    "test@example.com",
				Password: "4AeNkWVisJ",
				Name:     "", //名前が入力されていない
				Phone:    "000-0000-0000",
				Status:   true,
				Profile:  model.UserProfile{},
			},
			true, //エラーになるはず
		},
		{
			//⑤: 電話番号が入力されていないユーザ
			model.User{
				Email:    "test@example.com",
				Password: "4AeNkWVisJ",
				Name:     "test name",
				Phone:    "", //電話番号が入力されていない
				Status:   true,
				Profile:  model.UserProfile{},
			},
			true, //エラーになるはず
		},
	}
	for i, tt := range tests {
		rs, err := um.Validate(tt.in)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。ValidateUser()の出力結果: %s", i+1, rs)
		}

	}
}

//CreateUser()のテスト
func TestCreateUser(t *testing.T) {

	tests := []struct {
		in   model.User
		want bool
	}{
		{
			//①正しいユーザ
			model.User{
				Email:    "CreateTest@example.com",
				Password: "CrtTestPsw",
				Name:     "Crt Test",
				Phone:    "029-8475-1109",
				Status:   true,
				Profile:  model.UserProfile{},
			},
			false, //エラーはでないはず
		},
	}
	for i, tt := range tests {
		rs, err := um.Create(tt.in)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。の出力結果: %s", i+1, rs)
		}
	}

}

//GetUser()のテスト
//ユーザが取得できたらOK,できなければダメ
func TestGetUser(t *testing.T) {

	tests := []struct {
		in   int //userID
		want bool
	}{
		{
			//①先ほど作成したユーザ
			2,
			false, //エラーはでないはず
		},
	}
	for _, tt := range tests {
		_, err := um.GetById(tt.in)
		if err != tt.want {
			t.Errorf("userID:%dのユーザを取得できませんでした。", tt.in)
		}
	}
}

//UpdateUser()のテスト
//ユーザの情報が更新できなかったらダメ
func TestUpdateUser(t *testing.T) {

	tests := []struct {
		id    int
		after model.User
		want  bool
	}{
		{
			2, //先ほどテストで作ったユーザ
			model.User{
				Email:    "Upd@example.com",
				Password: "UpdTestPsw",
				Name:     "Upd Test",
				Phone:    "048-8476-8173",
				Status:   true,
				Profile:  model.UserProfile{},
			},
			false, //エラーはでないはず
		},
	}
	for i, tt := range tests {
		msg, err := um.Update(tt.id, tt.after)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。エラーメッセージ:%s", i+1, msg)
		}
	}

}

func TestDeleteUser(t *testing.T) {

	tests := []struct {
		id   int
		want bool
	}{
		{
			2,     //テストで作ったユーザ
			false, //エラーはでないはず
		},
		{
			9999999999,
			true, //存在しないユーザは削除できない
		},
	}
	for i, tt := range tests {
		msg, err := um.Delete(tt.id)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。エラーメッセージ:%s", i+1, msg)
		}
	}
}
