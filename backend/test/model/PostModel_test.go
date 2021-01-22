package test_model

import (
	"main/model"
	"testing"
	"time"
)

var pm = model.NewPostModel("test")

//ValidatePost()のテスト
func TestValidatePost(t *testing.T) {

	tests := []struct {
		in   model.Post
		want bool
	}{
		{
			//①正しい投稿
			model.Post{
				ChannelId: 1,
				UserId:    1,
				Content:   "test content",
				Date:      time.Now(),
				Status:    true,
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
				Status:    true,
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
				Status:    true,
			},
			true, //エラーになるはず
		},
		{
			//④: 内容が入力されていない投稿
			model.Post{
				ChannelId: 1,
				UserId:    1,
				Content:   "",
				Date:      time.Now(),
				Status:    true,
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
			//①正しい投稿
			model.Post{
				ChannelId: 1,
				UserId:    1,
				Content:   "test content",
				Date:      time.Now(),
				Status:    true,
			},
			false, //エラーはでないはず
		},
	}
	for i, tt := range tests {
		rs, err := pm.Create(tt.in)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。の出力結果: %s", i+1, rs)
		}
	}

}

//GetPost()のテスト
//投稿が取得できたらOK,できなければダメ
func TestGetPost(t *testing.T) {

	tests := []struct {
		in   int //userID
		want bool
	}{
		{
			//①先ほど作成した投稿
			1,
			false, //エラーはでないはず
		},
	}
	for _, tt := range tests {
		_, err := pm.GetById(tt.in)
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
			1, //先ほどテストで作った投稿
			model.Post{
				ChannelId: 1,
				UserId:    1,
				Content:   "Upd content",
				Date:      time.Now(), //timeは変えられない
				Status:    true,
			},
			false, //エラーはでないはず
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
			1,     //テストで作った投稿
			false, //エラーはでないはず
		},
		{
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
