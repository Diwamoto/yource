package model

import (
	"main/model"
	"testing"
)

var cm = model.NewChannelModel("test")

//ChannelModel.TableName()のテスト
func TestTableNameForChannelModel(t *testing.T) {
	want := "channels"
	tableName := cm.TableName()
	if tableName != want {
		t.Errorf("ChannelModel.TableName()の値が異常です。TableName()の出力結果: %s", tableName)
	}

}

//ChannelModel.Validate()のテスト
func TestValidateChannel(t *testing.T) {
	tests := []struct {
		in   model.Channel
		want bool
	}{
		{
			//①: 正しいチャンネル
			model.Channel{
				SpaceId:     1,
				Name:        "test name",
				Description: "test desc",
			},
			false, //エラーはでないはず
		},
		{
			//②: 存在しないスペースのチャンネル
			model.Channel{
				SpaceId:     9999,
				Name:        "test name",
				Description: "test desc",
			},
			true, //エラーになるはず
		},
		{
			//③: 名前が入力されていないチャンネル
			model.Channel{
				SpaceId:     0,
				Name:        "",
				Description: "test desc",
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

//ChannelModel.Create()のテスト
func TestCreateChannel(t *testing.T) {

	tests := []struct {
		in   model.Channel
		want bool
	}{
		{
			//①: 正しいチャンネル
			model.Channel{
				SpaceId:     1,
				Name:        "test name",
				Description: "test desc",
			},
			false, //エラーはでないはず
		},
		{
			//②: 存在しないスペースIDのチャンネル
			model.Channel{
				SpaceId:     9999,
				Name:        "test name",
				Description: "test desc",
			},
			true, //エラーになるはず
		},
		{
			//③: 同名のチャンネル
			model.Channel{
				SpaceId:     1,
				Name:        "test name",
				Description: "test desc",
			},
			true, //エラーになるはず
		},
	}
	for i, tt := range tests {
		rs, err := cm.Create(tt.in)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。の出力結果: %s", i+1, rs)
		}
	}

}

//ChannelModel.GetByAll()のテスト
//レコードの数が想定通りならオーケー、違うならダメ
func TestGetAllChannel(t *testing.T) {

	tests := []struct {
		want int //テストテーブルではチャンネルは2個存在する
	}{
		{
			//①: 全探索
			2, //探索した結果は2個なはず
		},
	}
	for _, tt := range tests {
		channels, err := cm.GetAll()
		if err || len(channels) != tt.want {
			t.Errorf("GetAll()でチャンネルを取得できませんでした。GetAll()の結果の個数: %d", len(channels))
		}
	}
}

//ChannelModel.GetById()のテスト
//チャンネルが取得できたらOK,できなければダメ
func TestGetChannelById(t *testing.T) {

	tests := []struct {
		in   int //channelId
		want bool
	}{
		{
			//①: 存在するチャンネル
			1,
			false, //取得できるはず
		},
		{
			//②: 存在しないチャンネル
			9999,
			true, //取得出来ないはず
		},
	}
	for i, tt := range tests {
		ret, err := cm.GetById(tt.in)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。GetById()の返り値:%#v", i+1, ret)
		}
	}
}

//ChannelModel.Find()のテスト
//チャンネルが取得できたらOK,できなければダメ
func TestFindChannel(t *testing.T) {

	tests := []struct {
		in   model.Channel
		t    string //検索する項目
		want bool
	}{
		{
			//①: 名前で検索
			model.Channel{
				Name: "test name",
			},
			"名前",
			false, //取得できるはず
		},
		{
			//②: 説明で検索
			model.Channel{
				Description: "test desc",
			},
			"説明",
			false, //取得できるはず
		},
		{
			//③: 存在しないスペースIDでの検索
			model.Channel{
				SpaceId: 9999,
			},
			"スペースID",
			true, //取得できないはず
		},
	}
	for i, tt := range tests {
		ret, err := cm.Find(tt.in)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。Find()の返り値:%#v", i+1, ret)
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
			//①: 正常に変更できる
			2, //先ほどテストで作ったチャンネル
			model.Channel{
				SpaceId:     1,
				Name:        "upd name",
				Description: "upd desc",
			},
			false, //エラーはでないはず
		},
		{
			//②: スペースIDは変更できない
			2, //先ほどテストで作ったチャンネル
			model.Channel{
				SpaceId:     2,
				Name:        "upd name",
				Description: "upd desc",
			},
			true, //エラーになるはず
		},
		{
			//③: 存在しないチャンネル
			3, //存在しない
			model.Channel{
				SpaceId:     1,
				Name:        "upd name",
				Description: "upd desc",
			},
			true, //エラーになるはず
		},
		{
			//④: 名前は空にできない
			2, //先ほどテストで作ったチャンネル
			model.Channel{
				SpaceId:     1,
				Name:        "",
				Description: "upd desc",
			},
			true, //エラーになるはず
		},
		{
			//⑤: 説明文は空欄にできる
			2, //先ほどテストで作ったチャンネル
			model.Channel{
				SpaceId:     1,
				Name:        "upd name",
				Description: "",
			},
			false, //エラーにならないはず
		},
		{
			//⑥: 既に存在するチャンネル名に変更はできない
			2, //先ほどテストで作ったチャンネル
			model.Channel{
				SpaceId:     1,
				Name:        "master name",
				Description: "master desc",
			},
			true, //エラーになるはず
		},
	}
	for i, tt := range tests {
		msg, err := cm.Update(tt.id, tt.after)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。エラーメッセージ:%s", i+1, msg)
		}
	}

}

//ChannelModel.Delete()のテスト
func TestDeleteChannel(t *testing.T) {

	tests := []struct {
		id   int
		want bool
	}{
		{
			//①: 存在するチャンネル
			2,     //テストで作ったチャンネル
			false, //エラーはでないはず
		},
		{
			//②: 存在しないチャンネル
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
