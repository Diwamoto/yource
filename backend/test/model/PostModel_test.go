package test_model

// var pm *model.PostModel

// //ValidatePost()のテスト
// func TestValidatePost(t *testing.T) {

// 	tests := []struct {
// 		in   model.Post
// 		want bool
// 	}{
// 		{
// 			//①正しい投稿
// 			model.Post{

// 			},
// 			false, //エラーはでないはず
// 		},
// 		{
// 			//②: メールアドレスがないデータの投稿
// 			model.Post{
// 				Email:    "", //メールアドレスが空欄
// 				Password: "4AeNkWVisJ",
// 				Name:     "test name",
// 				Phone:    "000-0000-0000",
// 				Status:   true,
// 				Profiles: model.PostProfile{},
// 			},
// 			true, //エラーになるはず
// 		},
// 		{
// 			//③: メールアドレスがおかしいデータの投稿
// 			model.Post{
// 				Email:    "testexample.com", //メールアドレスが正しくない
// 				Password: "4AeNkWVisJ",
// 				Name:     "test name",
// 				Phone:    "000-0000-0000",
// 				Status:   true,
// 				Profiles: model.PostProfile{},
// 			},
// 			true, //エラーになるはず
// 		},
// 		{
// 			//④: 名前が入力されていない投稿
// 			model.Post{
// 				Email:    "test@example.com",
// 				Password: "4AeNkWVisJ",
// 				Name:     "", //名前が入力されていない
// 				Phone:    "000-0000-0000",
// 				Status:   true,
// 				Profiles: model.PostProfile{},
// 			},
// 			true, //エラーになるはず
// 		},
// 		{
// 			//⑤: 電話番号が入力されていない投稿
// 			model.Post{
// 				Email:    "test@example.com",
// 				Password: "4AeNkWVisJ",
// 				Name:     "test name",
// 				Phone:    "", //電話番号が入力されていない
// 				Status:   true,
// 				Profiles: model.PostProfile{},
// 			},
// 			true, //エラーになるはず
// 		},
// 	}
// 	for i, tt := range tests {
// 		rs, err := pm.Validate(tt.in)
// 		if err != tt.want {
// 			t.Errorf("%d番目のテストが失敗しました。ValidatePost()の出力結果: %s", i+1, rs)
// 		}

// 	}
// }

// //CreatePost()のテスト
// func TestCreatePost(t *testing.T) {

// 	tests := []struct {
// 		in   model.Post
// 		want bool
// 	}{
// 		{
// 			//①正しい投稿
// 			model.Post{
// 				Email:    "CreateTest@example.com",
// 				Password: "CrtTestPsw",
// 				Name:     "Crt Test",
// 				Phone:    "029-8475-1109",
// 				Status:   true,
// 				Profiles: model.PostProfile{},
// 			},
// 			false, //エラーはでないはず
// 		},
// 	}
// 	for i, tt := range tests {
// 		rs, err := pm.Create(tt.in)
// 		if err != tt.want {
// 			t.Errorf("%d番目のテストが失敗しました。の出力結果: %s", i+1, rs)
// 		}
// 	}

// }

// //GetPost()のテスト
// //投稿が取得できたらOK,できなければダメ
// func TestGetPost(t *testing.T) {

// 	tests := []struct {
// 		in   int //userID
// 		want bool
// 	}{
// 		{
// 			//①先ほど作成した投稿
// 			1,
// 			false, //エラーはでないはず
// 		},
// 	}
// 	for _, tt := range tests {
// 		_, err := pm.GetById(tt.in)
// 		if err != tt.want {
// 			t.Errorf("userID:%dの投稿を取得できませんでした。", tt.in)
// 		}
// 	}
// }

// //UpdatePost()のテスト
// //投稿の情報が更新できなかったらダメ
// func TestUpdatePost(t *testing.T) {

// 	tests := []struct {
// 		id    int
// 		after model.Post
// 		want  bool
// 	}{
// 		{
// 			1, //先ほどテストで作った投稿
// 			model.Post{
// 				Email:    "Upd@example.com",
// 				Password: "UpdTestPsw",
// 				Name:     "Upd Test",
// 				Phone:    "048-8476-8173",
// 				Status:   true,
// 				Profiles: model.PostProfile{},
// 			},
// 			false, //エラーはでないはず
// 		},
// 	}
// 	for i, tt := range tests {
// 		msg, err := pm.Update(tt.id, tt.after)
// 		if err != tt.want {
// 			t.Errorf("%d番目のテストが失敗しました。エラーメッセージ:%s", i+1, msg)
// 		}
// 	}

// }

// func TestDeletePost(t *testing.T) {

// 	tests := []struct {
// 		id   int
// 		want bool
// 	}{
// 		{
// 			1,     //テストで作った投稿
// 			false, //エラーはでないはず
// 		},
// 		{
// 			9999999999,
// 			true, //投稿IDの最大値を持つ投稿はまだ存在していないという設定
// 		},
// 	}
// 	for i, tt := range tests {
// 		msg, err := pm.Delete(tt.id)
// 		if err != tt.want {
// 			t.Errorf("%d番目のテストが失敗しました。エラーメッセージ:%s", i+1, msg)
// 		}
// 	}
// }
