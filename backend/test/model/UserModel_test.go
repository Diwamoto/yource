package test

import (
	"main/model"
	"os"
	"testing"
)

var um *model.UserModel

//テストメイン関数
//全てのテストはここから呼ばれる
func TestMain(m *testing.M) {

	//テストに使う共通テスト用モデルを呼び出す
	um = model.NewUserModel("test")
	code := m.Run()

	//テスト用データベースの連番をリセット
	um.ExecRawSQL("ALTER TABLE `users` auto_increment = 1;")
	os.Exit(code)
}

//ValidateUser()のテスト
func TestValidateUser(t *testing.T) {

	//テスト用のデータベースからユーザモデルを呼び出す。
	um := model.NewUserModel("test")

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
				Profiles: model.UserProfile{},
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
				Profiles: model.UserProfile{},
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
				Profiles: model.UserProfile{},
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
				Profiles: model.UserProfile{},
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
				Profiles: model.UserProfile{},
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

	//テスト用のデータベースからユーザモデルを呼び出す。
	um := model.NewUserModel("test")

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
				Profiles: model.UserProfile{},
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

	//テスト用のデータベースからユーザモデルを呼び出す。
	um := model.NewUserModel("test")

	tests := []struct {
		in   int //userID
		want bool
	}{
		{
			//①先ほど作成したユーザ
			1,
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

	//テスト用のデータベースからユーザモデルを呼び出す。
	um := model.NewUserModel("test")

	tests := []struct {
		id    int
		after model.User
		want  bool
	}{
		{
			1, //先ほどテストで作ったユーザ
			model.User{
				Email:    "Upd@example.com",
				Password: "UpdTestPsw",
				Name:     "Upd Test",
				Phone:    "048-8476-8173",
				Status:   true,
				Profiles: model.UserProfile{},
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

	//テスト用のデータベースからユーザモデルを呼び出す。
	um := model.NewUserModel("test")

	tests := []struct {
		id   int
		want bool
	}{
		{
			1,     //テストで作ったユーザ
			false, //エラーはでないはず
		},
		{
			9999999999,
			true, //ユーザIDの最大値を持つユーザはまだ存在していないという設定
		},
	}
	for i, tt := range tests {
		msg, err := um.Delete(tt.id)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。エラーメッセージ:%s", i+1, msg)
		}
	}
}
