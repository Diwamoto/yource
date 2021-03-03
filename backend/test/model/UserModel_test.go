package model

import (
	"errors"
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
				Status:   1,
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
				Status:   1,
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
				Status:   1,
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
		//
		//
		// メールアドレスのユニークチェックはCreate()とUpdate()で処理が違う為ここではバリデーションしない
		//
		// {
		// 	//⑥(4): メールアドレスが既にデータベースに存在しているユーザ
		// 	model.User{
		// 		Email:    "master@example.com",
		// 		Password: "4AeNkWVisJ",
		// 		Name:     "test name",
		// 		Phone:    "000-0000-0000",
		// 		Nickname: "Crt Nick name",
		// 		Status:   1,
		// 		Profile:  model.UserProfile{},
		// 	},
		// 	true, //エラーになるはず
		// },
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
		want error
	}{
		{
			//①: 正しいユーザ
			model.User{
				Email:    "CreateTest@example.com",
				Password: "CrtTestPsw",
				Name:     "Crt Test",
				Phone:    "000-0000-0000",
				Nickname: "Crt Nick name",
				Status:   1,
				Profile:  model.UserProfile{},
			},
			nil, //エラーはでないはず
		},
		{
			//②: データがおかしいユーザ
			model.User{
				Email:    "",
				Password: "CrtTestPsw",
				Name:     "Crt Test",
				Phone:    "000-0000-0000",
				Nickname: "Crt Nick name",
				Status:   1,
				Profile:  model.UserProfile{},
			},
			errors.New("[\"メールアドレスを入力してください。\"]"), //エラーになるはず
		},
		{
			//②: 既にメールアドレスがデータベースに存在するユーザ
			model.User{
				Email:    "CreateTest@example.com",
				Password: "CrtTestPsw",
				Name:     "Crt Test",
				Phone:    "000-0000-0000",
				Nickname: "Crt Nick name",
				Status:   1,
				Profile:  model.UserProfile{},
			},
			errors.New("入力されたメールアドレスは既に登録されています。"), //エラーはでないはず
		},
	}
	for i, tt := range tests {
		_, err := um.Create(tt.in)
		if err != tt.want {
			if err.Error() != tt.want.Error() {
				t.Errorf("%d番目のテストが失敗しました。エラー内容: %#v", i+1, err.Error())
			}
		}
	}

}

//UserModel.GetById()のテスト
//ユーザが取得できたらOK,できなければダメ
func TestGetUserById(t *testing.T) {

	tests := []struct {
		in   int //userID
		want int //帰ってくる数
	}{
		{
			//①: 先ほど作成したユーザ
			2,
			1, //一つのみ帰ってくるはず
		},
		{
			//②: 存在しないidのユーザ
			9999999,
			0, //データは帰ってこないはず
		},
	}
	for i, tt := range tests {
		users, _ := um.GetById(tt.in)
		if len(users) != tt.want {
			t.Errorf("%d番目のテストが失敗しました。期待した結果:%d 実際の結果:%d", i, tt.want, len(users))
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
		want error
	}{
		{
			//①: 条件無しで全探索(GetAll()をここでカバー)
			model.User{},
			"全探索",
			nil, //検索は成功するはず
		},
		{
			//②: メールアドレスで検索
			model.User{
				Email: "CreateTest@example.com",
			},
			"メールアドレス",
			nil, //検索は成功するはず
		},
		{
			//③: ユーザ名で検索
			model.User{
				Name: "Crt Test",
			},
			"ユーザ名",
			nil, //検索は成功するはず
		},
		{
			//④: 電話番号で検索
			model.User{
				Phone: "000-0000-0000",
			},
			"電話番号",
			nil, //検索は成功するはず
		},
		{
			//⑤: ステータスで検索
			model.User{
				Status: 1,
			},
			"有効状態のユーザ",
			nil, //検索は成功するはず
		},
		{
			//⑥: ニックネームで検索
			model.User{
				Nickname: "Crt Nick name",
			},
			"ニックネーム",
			nil, //検索は成功するはず
		},
		// {
		// 	//⑦: TODO: カバレッジ100%の為にエラーを吐かせる
		// 	model.User{
		// 		Email:    "What is Email?",
		// 		Password: "Password id 1234!",
		// 		Name:     "I dont know...",
		// 		Nickname: "Mike stands for McDonald's.",
		// 		Phone:    "I have PocketBell!",
		// 		Status:   8,
		// 	},
		// 	"ニックネーム",
		// 	errors.New(""), //意図的にエラーを発生させる まだできてないです
		// },
	}
	for _, tt := range tests {
		_, err := um.Find(tt.in)
		if err != tt.want {
			t.Errorf("「%s」での検索が失敗しました。エラー内容:%#v", tt.t, err)
		}
	}
}

//UserModel.Update()のテスト
//ユーザの情報が更新できなかったらダメ
func TestUpdateUser(t *testing.T) {

	tests := []struct {
		id    int
		after model.User
		want  error
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
				Status:   0,
				Profile:  model.UserProfile{},
			},
			nil, //エラーはでないはず
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
				Status:   1,
				Profile:  model.UserProfile{},
			},
			errors.New("ユーザが存在しません。"), //更新はできないはず
		},
		{
			//③: データがおかしいユーザ
			2, //先ほどテストで作ったユーザ
			model.User{
				Email:    "test",
				Password: "UpdTestPsw",
				Name:     "Upd Test",
				Phone:    "",
				Nickname: "",
				Status:   0,
			},
			errors.New("[\"正しいメールアドレスを入力してください。\"]"), //エラーになるはず
		},
	}
	for i, tt := range tests {
		_, err := um.Update(tt.id, tt.after)
		if err != tt.want {
			if err.Error() != tt.want.Error() {
				t.Errorf("%d番目のテストが失敗しました。エラーメッセージ:%#v", i+1, err.Error())
			}
		}
	}

}

//UserModel.Delete()のテスト
func TestDeleteUser(t *testing.T) {

	tests := []struct {
		id   int
		want error
	}{
		{
			//①: 存在するユーザ
			3,   //テストで作ったユーザ
			nil, //エラーはでないはず
		},
		{
			//②: 存在しないユーザ
			9999999999,
			errors.New("削除するユーザが存在しません。"), //存在しないユーザは削除できない
		},
	}
	for i, tt := range tests {
		err := um.Delete(tt.id)
		if err != tt.want {
			if err.Error() != tt.want.Error() {
				t.Errorf("%d番目のテストが失敗しました。エラーメッセージ:%#v", i+1, err)
			}
		}
	}
}
