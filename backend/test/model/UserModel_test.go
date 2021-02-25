package model

import (
	"main/model"
	"testing"
)

var um = model.NewUserModel("test")

//UserModel.TableName()のテスト
func TestTableNameForUserModel(t *testing.T) {
	want := "users"
	tableName := um.TableName()
	if tableName != want {
		t.Errorf("UserModel.TableName()の値が異常です。TableName()の出力結果: %s", tableName)
	}

}

//UserModel.Validate()のテスト
func TestValidateUser(t *testing.T) {

	tests := []struct {
		in   model.User
		want bool
	}{
		{
			//①: 正しいユーザ
			model.User{
				Email:    "test@example.com",
				Password: "4AeNkWVisJ",
				Name:     "test name",
				Phone:    "000-0000-0000",
				Nickname: "Crt Nick name",
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
				Nickname: "Crt Nick name",
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
				Nickname: "Crt Nick name",
				Status:   true,
				Profile:  model.UserProfile{},
			},
			true, //エラーになるはず
		},
		// 2021/2/2 メールアドレスとパスワードで登録できる様にするためにいったんテストを止める
		// {
		// 	//④: 名前が入力されていないユーザ
		// 	model.User{
		// 		Email:    "test@example.com",
		// 		Password: "4AeNkWVisJ",
		// 		Name:     "", //名前が入力されていない
		// 		Phone:    "000-0000-0000",
		// 		Nickname: "Crt Nick name",
		// 		Status:   true,
		// 		Profile:  model.UserProfile{},
		// 	},
		// 	true, //エラーになるはず
		// },
		// {
		// 	//⑤: 電話番号が入力されていないユーザ
		// 	model.User{
		// 		Email:    "test@example.com",
		// 		Password: "4AeNkWVisJ",
		// 		Name:     "test name",
		// 		Phone:    "", //電話番号が入力されていない
		// 		Nickname: "Crt Nick name",
		// 		Status:   true,
		// 		Profile:  model.UserProfile{},
		// 	},
		// 	true, //エラーになるはず
		// },
		{
			//⑥(4): メールアドレスが既にデータベースに存在しているユーザ
			model.User{
				Email:    "master@example.com",
				Password: "4AeNkWVisJ",
				Name:     "test name",
				Phone:    "000-0000-0000",
				Nickname: "Crt Nick name",
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

//UserModel.Create()のテスト
func TestCreateUser(t *testing.T) {

	tests := []struct {
		in   model.User
		want bool
	}{
		{
			//①: 正しいユーザ
			model.User{
				Email:    "CreateTest@example.com",
				Password: "CrtTestPsw",
				Name:     "Crt Test",
				Phone:    "000-0000-0000",
				Nickname: "Crt Nick name",
				Status:   true,
				Profile:  model.UserProfile{},
			},
			false, //エラーはでないはず
		},
		{
			//②: データがおかしいユーザ
			model.User{
				Email:    "",
				Password: "CrtTestPsw",
				Name:     "Crt Test",
				Phone:    "000-0000-0000",
				Nickname: "Crt Nick name",
				Status:   true,
				Profile:  model.UserProfile{},
			},
			true, //エラーになるはず
		},
	}
	for i, tt := range tests {
		rs, err := um.Create(tt.in)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。の出力結果: %s", i+1, rs)
		}
	}

}

//UserModel.GetAll()のテスト
//ユーザが取得できたらOK,できなければダメ
func TestGetAllUser(t *testing.T) {

	tests := []struct {
		want bool
	}{
		{
			//①: 全て取得できる
			false, //取得できるはず
		},
	}
	for _, tt := range tests {
		_, err := um.GetAll()
		if err != tt.want {
			t.Errorf("GetAll()を用いてユーザを取得することができませんでした。")
		}
	}
}

//UserModel.GetById()のテスト
//ユーザが取得できたらOK,できなければダメ
func TestGetUserById(t *testing.T) {

	tests := []struct {
		in   int //userID
		want bool
	}{
		{
			//①: 先ほど作成したユーザ
			2,
			false, //エラーはでないはず
		},
		{
			//②: 存在しないidのユーザ
			9999999,
			true, //エラーになるはず
		},
	}
	for _, tt := range tests {
		_, err := um.GetById(tt.in)
		if err != tt.want {
			t.Errorf("userID:%dのユーザを取得できませんでした。", tt.in)
		}
	}
}

//UserModel.Find()のテスト
//ユーザを検索する
//検索の失敗についての定義は議論中
func TestFindUser(t *testing.T) {
	tests := []struct {
		in   model.User
		t    string //検索の種類
		want bool
	}{
		{
			//①: メールアドレスで検索
			model.User{
				Email: "CreateTest@example.com",
			},
			"メールアドレス",
			false, //検索は成功するはず
		},
		{
			//②: ユーザ名で検索
			model.User{
				Name: "Crt Test",
			},
			"ユーザ名",
			false, //検索は成功するはず
		},
		{
			//③: 電話番号で検索
			model.User{
				Phone: "000-0000-0000",
			},
			"電話番号",
			false, //検索は成功するはず
		},
		{
			//④: ステータスで検索
			model.User{
				Status: true,
			},
			"有効状態のユーザ",
			false, //検索は成功するはず
		},
		{
			//⑤: ニックネームで検索
			model.User{
				Nickname: "Crt Nick name",
			},
			"ニックネーム",
			false, //検索は成功するはず
		},
	}
	for _, tt := range tests {
		_, err := um.Find(tt.in)
		if err != tt.want {
			t.Errorf("「%s」での検索が失敗しました。", tt.t)
		}
	}
}

//UserModel.Update()のテスト
//ユーザの情報が更新できなかったらダメ
func TestUpdateUser(t *testing.T) {

	tests := []struct {
		id    int
		after model.User
		want  bool
	}{
		{
			//①: 存在するユーザ
			2, //先ほどテストで作ったユーザ
			model.User{
				Email:    "Upd@example.com",
				Password: "UpdTestPsw",
				Name:     "Upd Test",
				Phone:    "048-8476-8173",
				Nickname: "Upd nickname",
				Status:   true,
				Profile:  model.UserProfile{},
			},
			false, //エラーはでないはず
		},
		{
			//②: 存在しないユーザ
			99999, //存在しないユーザ
			model.User{
				Email:    "Upd@example.com",
				Password: "UpdTestPsw",
				Name:     "Upd Test",
				Phone:    "048-8476-8173",
				Nickname: "Upd nickname",
				Status:   true,
				Profile:  model.UserProfile{},
			},
			true, //更新はできないはず
		},
	}
	for i, tt := range tests {
		msg, err := um.Update(tt.id, tt.after)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。エラーメッセージ:%s", i+1, msg)
		}
	}

}

//UserModel.Delete()のテスト
func TestDeleteUser(t *testing.T) {

	tests := []struct {
		id   int
		want bool
	}{
		{
			//①: 存在するユーザ
			2,     //テストで作ったユーザ
			false, //エラーはでないはず
		},
		{
			//②: 存在しないユーザ
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
