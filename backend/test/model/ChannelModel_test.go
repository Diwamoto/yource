package test_model

import (
	"main/model"
	"testing"
)

var cm = model.NewChannelModel("test")

//ValidateChannel()のテスト
func TestValidateChannel(t *testing.T) {

	tests := []struct {
		in   model.Channel
		want bool
	}{
		{
			//①正しいチャンネル
			model.Channel{
				SpaceId:     1,
				UserId:      1,
				Name:        "test name",
				Discription: "test disc",
			},
			false, //エラーはでないはず
		},
		{
			//②: 存在しないスペースのチャンネル
			model.Channel{
				SpaceId:     9999,
				UserId:      1,
				Name:        "test name",
				Discription: "test disc",
			},
			true, //エラーになるはず
		},
		{
			//③: 存在しないユーザが作成したチャンネル
			model.Channel{
				SpaceId:     1,
				UserId:      9999,
				Name:        "test name",
				Discription: "test disc",
			},
			true, //エラーになるはず
		},
		{
			//④: 名前が入力されていないチャンネル
			model.Channel{
				SpaceId:     0,
				UserId:      0,
				Name:        "",
				Discription: "test disc",
			},
			true, //エラーになるはず
		},
	}
	for i, tt := range tests {
		rs, err := cm.Validate(tt.in)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。ValidateChannel()の出力結果: %s", i+1, rs)
		}

	}
}

//CreateChannel()のテスト
func TestCreateChannel(t *testing.T) {

	tests := []struct {
		in   model.Channel
		want bool
	}{
		{
			//①正しいチャンネル
			model.Channel{
				SpaceId:     1,
				UserId:      1,
				Name:        "test name",
				Discription: "test disc",
			},
			false, //エラーはでないはず
		},
	}
	for i, tt := range tests {
		rs, err := cm.Create(tt.in)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。の出力結果: %s", i+1, rs)
		}
	}

}

//GetChannel()のテスト
//チャンネルが取得できたらOK,できなければダメ
func TestGetChannel(t *testing.T) {

	tests := []struct {
		in   int //userID
		want bool
	}{
		{
			//①先ほど作成したチャンネル
			2,
			false, //エラーはでないはず
		},
	}
	for _, tt := range tests {
		_, err := cm.GetById(tt.in)
		if err != tt.want {
			t.Errorf("userID:%dのチャンネルを取得できませんでした。", tt.in)
		}
	}
}

//UpdateChannel()のテスト
//チャンネルの情報が更新できなかったらダメ
func TestUpdateChannel(t *testing.T) {

	tests := []struct {
		id    int
		after model.Channel
		want  bool
	}{
		{
			2, //先ほどテストで作ったチャンネル
			model.Channel{
				SpaceId:     1,
				UserId:      1,
				Name:        "test name",
				Discription: "test disc",
			},
			false, //エラーはでないはず
		},
	}
	for i, tt := range tests {
		msg, err := cm.Update(tt.id, tt.after)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。エラーメッセージ:%s", i+1, msg)
		}
	}

}

func TestDeleteChannel(t *testing.T) {

	tests := []struct {
		id   int
		want bool
	}{
		{
			2,     //テストで作ったチャンネル
			false, //エラーはでないはず
		},
		{
			9999999999,
			true, //存在しないチャンネルは削除できない
		},
	}
	for i, tt := range tests {
		msg, err := cm.Delete(tt.id)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。エラーメッセージ:%s", i+1, msg)
		}
	}
}
