package model

import (
	"main/model"
	"testing"
	"time"
)

var pm = model.NewPostModel("test")

//PostModel.TableName()のテスト
func TestTableNameForPostModel(t *testing.T) {
	want := "posts"
	tableName := pm.TableName()
	if tableName != want {
		t.Errorf("PostModel.TableName()の値が異常です。TableName()の出力結果: %s", tableName)
	}

}

//ValidatePost()のテスト
func TestValidatePost(t *testing.T) {

	tests := []struct {
		in   model.Post
		want bool
	}{
		{
			//①: 正しい投稿
			model.Post{
				ChannelId: 1,
				UserId:    1,
				Content:   "test content",
				Date:      time.Now(),
			},
			false, //エラーはでないはず
		},
		{
			//②: 存在しないチャンネルの投稿
			model.Post{
				ChannelId: 9999,
				UserId:    1,
				Content:   "test content",
				Date:      time.Now(),
			},
			true, //エラーになるはず
		},
		{
			//③: 存在しないユーザが作成した投稿
			model.Post{
				ChannelId: 1,
				UserId:    9999,
				Content:   "test content",
				Date:      time.Now(),
			},
			true, //エラーになるはず
		},
		{
			//④: チャンネルidが入力されていない投稿
			model.Post{
				ChannelId: 0,
				UserId:    1,
				Content:   "",
				Date:      time.Now(),
			},
			true, //エラーになるはず
		},
		{
			//⑤: ユーザidが入力されていない投稿
			model.Post{
				ChannelId: 1,
				UserId:    0,
				Content:   "",
				Date:      time.Now(),
			},
			true, //エラーになるはず
		},
		{
			//⑥: 内容が入力されていない投稿
			model.Post{
				ChannelId: 1,
				UserId:    1,
				Content:   "",
				Date:      time.Now(),
			},
			true, //エラーになるはず
		},
		{
			//⑦: 投稿日が入力されていない投稿
			model.Post{
				ChannelId: 1,
				UserId:    1,
				Content:   "",
				Date:      time.Time{},
			},
			true, //エラーになるはず
		},
	}
	for i, tt := range tests {
		rs, err := pm.Validate(tt.in)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。ValidatePost()の出力結果: %s", i+1, rs)
		}

	}
}

//CreatePost()のテスト
func TestCreatePost(t *testing.T) {

	tests := []struct {
		in   model.Post
		want bool
	}{
		{
			//①: 正しい投稿
			model.Post{
				ChannelId: 1,
				UserId:    1,
				Content:   "test content",
				Date:      time.Date(2018, 3, 11, 12, 0, 0, 0, time.Local),
			},
			false, //エラーはでないはず
		},
		{
			//②: 作成できない投稿
			model.Post{
				ChannelId: 0,
				UserId:    1,
				Content:   "test content",
				Date:      time.Date(2018, 3, 11, 12, 0, 0, 0, time.Local),
			},
			true, //エラーになるはず
		},
	}
	for i, tt := range tests {
		rs, err := pm.Create(tt.in)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。の出力結果: %s", i+1, rs)
		}
	}

}

//PostModel.GetById()のテスト
//投稿が取得できたらOK,できなければダメ
func TestGetPostById(t *testing.T) {

	tests := []struct {
		in   int //userID
		want bool
	}{
		{
			//①: 先ほど作成した投稿
			1,
			false, //エラーはでないはず
		},
		{
			//②: 存在しない投稿
			9999,
			true, //エラーになるはず
		},
	}
	for _, tt := range tests {
		_, err := pm.GetById(tt.in)
		if err != tt.want {
			t.Errorf("userID:%dの投稿を取得できませんでした。", tt.in)
		}
	}
}

//PostModel.GetByChannelId()のテスト
//投稿が取得できたらOK,できなければダメ
func TestGetPostsByChannelId(t *testing.T) {

	tests := []struct {
		in   int //userID
		want bool
	}{
		{
			//①: テストで作成したチャンネルの投稿
			1,
			false, //エラーはでないはず
		},
		{
			//①: 存在しないチャンネルの投稿
			9999,
			true, //エラーになるはず
		},
	}
	for _, tt := range tests {
		_, err := pm.GetByChannelId(tt.in)
		if err != tt.want {
			t.Errorf("userID:%dの投稿を取得できませんでした。", tt.in)
		}
	}
}

//PostModel.Find()のテスト
//投稿が取得できたらOK,できなければダメ
func TestFindPosts(t *testing.T) {

	tests := []struct {
		in   model.Post
		want bool
	}{
		{
			//①: チャンネルIDで検索
			model.Post{
				ChannelId: 1,
			},
			false, //エラーはでないはず
		},
		{
			//②: ユーザIDで検索
			model.Post{
				UserId: 1,
			},
			false, //エラーはでないはず
		},
		{
			//③: 内容で検索
			model.Post{
				Content: "test content",
			},
			false, //エラーはでないはず
		},
		{
			//④: 投稿日で検索
			model.Post{
				Date: time.Date(2018, 3, 11, 12, 0, 0, 0, time.Local),
			},
			false, //エラーはでないはず
		},
		{
			//⑤: 存在しないチャンネルIDで検索
			model.Post{
				ChannelId: 9999,
			},
			true, //エラーになるはず
		},
		{
			//⑥: 存在しないユーザIDで検索
			model.Post{
				UserId: 9999,
			},
			true, //エラーになるはず
		},
		{
			//⑦: 存在しない内容で検索
			model.Post{
				Content: "内容",
			},
			true, //エラーになるはず
		},
		{
			//⑧: dbに存在しない投稿日で検索
			model.Post{
				Date: time.Date(2021, 3, 11, 12, 0, 0, 0, time.Local),
			},
			true, //エラーになるはず
		},
	}
	for i, tt := range tests {
		posts, err := pm.Find(tt.in)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。Find()の結果:%#v", i, posts)
		}
	}
}

//PostModel.GetAll()のテスト
//投稿が取得できたらOK,できなければダメ
func TestGetAllPosts(t *testing.T) {

	tests := []struct {
		want bool
	}{
		{
			//①: 投稿を全取得
			false, //テストで作成しているため、エラーはでないはず
		},
	}
	for _, tt := range tests {
		_, err := pm.GetAll()
		if err != tt.want {
			t.Errorf("投稿を取得できませんでした。")
		}
	}
}

//PostModel.GetByUserId()のテスト
//投稿が取得できたらOK,できなければダメ
func TestGetPostByUserId(t *testing.T) {

	tests := []struct {
		in   int //userID
		want bool
	}{
		{
			//①: 先ほど作成した投稿
			1,
			false, //エラーはでないはず
		},
		{
			//②: 存在しないユーザIDで検索
			9999,
			true, //エラーになるはず
		},
	}
	for _, tt := range tests {
		_, err := pm.GetByUserId(tt.in)
		if err != tt.want {
			t.Errorf("userID:%dの投稿を取得できませんでした。", tt.in)
		}
	}
}

//UpdatePost()のテスト
//投稿の情報が更新できなかったらダメ
func TestUpdatePost(t *testing.T) {

	tests := []struct {
		id    int
		after model.Post
		want  bool
	}{
		{
			//①: 正常に変更できる
			1, //先ほどテストで作った投稿
			model.Post{
				ChannelId: 1,
				UserId:    1,
				Content:   "Upd content",
				Date:      time.Now(), //timeは変えられない
			},
			false, //エラーはでないはず
		},
		{
			//②: 異常なデータ
			9999, //存在しない投稿
			model.Post{
				ChannelId: 1,
				UserId:    1,
				Content:   "Upd content",
				Date:      time.Now(), //timeは変えられない
			},
			true, //エラーになるはず
		},
	}
	for i, tt := range tests {
		msg, err := pm.Update(tt.id, tt.after)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。エラーメッセージ:%s", i+1, msg)
		}
	}

}

func TestDeletePost(t *testing.T) {

	tests := []struct {
		id   int
		want bool
	}{
		{
			//①: 存在する投稿
			1,     //テストで作った投稿
			false, //エラーはでないはず
		},
		{
			//②: 存在しない投稿
			9999999999,
			true, //存在しない投稿は削除できない
		},
	}
	for i, tt := range tests {
		msg, err := pm.Delete(tt.id)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。エラーメッセージ:%s", i+1, msg)
		}
	}
}
