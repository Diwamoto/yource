package model

import (
	"main/model"
	"testing"
)

var sm = model.NewSpaceModel("test")

//SpaceModel.Validate()のテスト
func TestValidateSpace(t *testing.T) {

	tests := []struct {
		in   model.Space
		want bool
	}{
		{
			//①: 正しいスペース
			model.Space{
				UserId:      1,
				Name:        "test name",
				Description: "test desc",
				SubDomain:   "subdomain",
			},
			false, //エラーはでないはず
		},
		{
			//②: ユーザIDが存在しないデータのスペース
			model.Space{
				UserId:      0, //ユーザIDが入力されていない
				Name:        "test name",
				Description: "test desc",
				SubDomain:   "subdomain",
			},
			true, //エラーになるはず
		},
		{
			//③: スペース名が存在しないデータのスペース
			model.Space{
				UserId:      1,
				Name:        "", //スペース名が入力されていない
				Description: "test desc",
				SubDomain:   "subdomain",
			},
			true, //エラーになるはず
		},
		{
			//④: サブドメインが入力されていないスペース
			model.Space{
				UserId:      1,
				Name:        "test name",
				Description: "test desc",
				SubDomain:   "",
			},
			true, //エラーになるはず
		},
		{
			//⑤: サブドメインが半角英字以外でスペース
			model.Space{
				UserId:      1,
				Name:        "test name",
				Description: "test desc",
				SubDomain:   "@wowoj@^~",
			},
			true, //エラーになるはず
		},
	}
	for i, tt := range tests {
		rs, err := sm.Validate(tt.in)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。ValidateSpace()の出力結果: %s", i+1, rs)
		}

	}
}

//SpaceModel.Create()のテスト
//スペースが作成できたらok, できなければだめ
func TestCreateSpace(t *testing.T) {

	tests := []struct {
		in   model.Space
		want bool
	}{
		{
			//①: 正しいスペース
			model.Space{
				UserId:      1,
				Name:        "test name",
				Description: "test desc",
				SubDomain:   "subdomain",
			},
			false, //エラーはでないはず
		},
		{
			//②: 名前が入力されていない
			model.Space{
				UserId:      1,
				Name:        "",
				Description: "test desc",
				SubDomain:   "subdomain",
			},
			true, //エラーになるはず
		},
		{
			//③: サブドメインが入力されていない
			model.Space{
				UserId:      1,
				Name:        "test name",
				Description: "test desc",
				SubDomain:   "",
			},
			true, //エラーになるはず
		},
	}
	for i, tt := range tests {
		rs, err := sm.Create(tt.in)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。の出力結果: %s", i+1, rs)
		}
	}

}

//SpaceModel.GetAll()のテスト
//スペースが取得できたらOK,できなければダメ
func TestGetAllSpace(t *testing.T) {

	tests := []struct {
		want bool
	}{
		{
			//①: 全てのスペースを取得
			false, //エラーはでないはず
		},
	}
	for _, tt := range tests {
		_, err := sm.GetAll()
		if err != tt.want {
			t.Errorf("全てのスペースを取得できませんでした。")
		}
	}
}

//SpaceModel.GetById()のテスト
//スペースが取得できたらOK,できなければダメ
func TestGetSpaceById(t *testing.T) {

	tests := []struct {
		in   int //userID
		want bool
	}{
		{
			//①: 先ほど作成したスペース
			2,
			false, //エラーはでないはず
		},
	}
	for _, tt := range tests {
		_, err := sm.GetById(tt.in)
		if err != tt.want {
			t.Errorf("userID:%dのスペースを取得できませんでした。", tt.in)
		}
	}
}

//SpaceModel.Find()のテスト
//スペースが取得できたらOK,できなければダメ
func TestFindSpace(t *testing.T) {

	tests := []struct {
		in   model.Space //検索するスペース
		t    string      //検索の種類
		want bool
	}{
		{
			//①: 正しいスペース
			model.Space{
				UserId:      1,
				Name:        "test name",
				Description: "test desc",
				SubDomain:   "subdomain",
			},
			"名前",
			false, //エラーはでないはず
		},
		{
			//①: 正しいスペース
			model.Space{
				UserId:      1,
				Name:        "test name",
				Description: "test desc",
				SubDomain:   "subdomain",
			},
			"説明",
			false, //エラーはでないはず
		},
		{
			//①: 正しいスペース
			model.Space{
				UserId:      1,
				Name:        "test name",
				Description: "test desc",
				SubDomain:   "subdomain",
			},
			"サブドメイン",
			false, //エラーはでないはず
		},
	}
	for _, tt := range tests {
		_, err := sm.Find(tt.in)
		if err != tt.want {
			t.Errorf("%sでスペースを検索できませんでした。", tt.t)
		}
	}
}

//SpaceModel.GetByUserId()のテスト
//スペースが取得できたらOK,できなければダメ
func TestGetSpaceByUserId(t *testing.T) {

	tests := []struct {
		in   int //userID
		want bool
	}{
		{
			//①: 先ほど作成したスペース
			2,
			false, //エラーはでないはず
		},
	}
	for _, tt := range tests {
		_, err := sm.GetById(tt.in)
		if err != tt.want {
			t.Errorf("userID:%dのスペースを取得できませんでした。", tt.in)
		}
	}
}

//SpaceModel.Update()のテスト
//スペースの情報が更新できなかったらダメ
func TestUpdateSpace(t *testing.T) {

	tests := []struct {
		id    int
		after model.Space
		want  bool
	}{
		{
			//①: 正常に変更できる
			2, //先ほどテストで作ったスペース
			model.Space{
				UserId:      1,
				Name:        "Upd Name",
				Description: "Upd desc",
				SubDomain:   "upd",
				Status:      false,
				Publish:     false,
			},
			false, //エラーはでないはず
		},
	}
	for i, tt := range tests {
		msg, err := sm.Update(tt.id, tt.after)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。エラーメッセージ:%s", i+1, msg)
		}
	}

}

//SpaceModel.Delete()のテスト
//正しいデータが削除できればオーケー、できなければダメ
func TestDeleteSpace(t *testing.T) {

	tests := []struct {
		id   int
		want bool
	}{
		{
			//①: 存在するスペース
			2,     //テストで作ったスペース
			false, //エラーはでないはず
		},
		{
			//②: 存在しないスペース
			9999999999,
			true, //存在しないスペースidは削除できない
		},
	}
	for i, tt := range tests {
		msg, err := sm.Delete(tt.id)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。エラーメッセージ:%s", i+1, msg)
		}
	}
}
