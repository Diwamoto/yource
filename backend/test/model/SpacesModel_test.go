package model

import (
	"main/model"
	"testing"
)

var sm = model.NewSpaceModel("test")

//ValidateSpace()のテスト
func TestValidateSpace(t *testing.T) {

	tests := []struct {
		in   model.Space
		want bool
	}{
		{
			//①正しいスペース
			model.Space{
				UserId:      1,
				Name:        "test name",
				Discription: "test disc",
				SubDomain:   "subdomain",
				Status:      true,
				Publish:     true,
			},
			false, //エラーはでないはず
		},
		{
			//②: ユーザIDが存在しないデータのスペース
			model.Space{
				UserId:      0, //ユーザIDが入力されていない
				Name:        "test name",
				Discription: "test disc",
				SubDomain:   "subdomain",
				Status:      true,
				Publish:     true,
			},
			true, //エラーになるはず
		},
		{
			//③: スペース名が存在しないデータのスペース
			model.Space{
				UserId:      1,
				Name:        "", //スペース名が入力されていない
				Discription: "test disc",
				SubDomain:   "subdomain",
				Status:      true,
				Publish:     true,
			},
			true, //エラーになるはず
		},
		{
			//④: サブドメインが入力されていないスペース
			model.Space{
				UserId:      1,
				Name:        "test name",
				Discription: "test disc",
				SubDomain:   "",
				Status:      true,
				Publish:     true,
			},
			true, //エラーになるはず
		},
		{
			//⑤: サブドメインが半角英字以外でスペース
			model.Space{
				UserId:      1,
				Name:        "test name",
				Discription: "test disc",
				SubDomain:   "@wowoj@^~",
				Status:      true,
				Publish:     true,
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

//CreateSpace()のテスト
func TestCreateSpace(t *testing.T) {

	tests := []struct {
		in   model.Space
		want bool
	}{
		{
			//①正しいスペース
			model.Space{
				UserId:      1,
				Name:        "test name",
				Discription: "test disc",
				SubDomain:   "subdomain",
				Status:      true,
				Publish:     true,
			},
			false, //エラーはでないはず
		},
	}
	for i, tt := range tests {
		rs, err := sm.Create(tt.in)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。の出力結果: %s", i+1, rs)
		}
	}

}

//GetSpace()のテスト
//スペースが取得できたらOK,できなければダメ
func TestGetSpace(t *testing.T) {

	tests := []struct {
		in   int //userID
		want bool
	}{
		{
			//①先ほど作成したスペース
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

//UpdateSpace()のテスト
//スペースの情報が更新できなかったらダメ
func TestUpdateSpace(t *testing.T) {

	tests := []struct {
		id    int
		after model.Space
		want  bool
	}{
		{
			2, //先ほどテストで作ったスペース
			model.Space{
				UserId:      1,
				Name:        "Upd Name",
				Discription: "Upd disc",
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

func TestDeleteSpace(t *testing.T) {

	tests := []struct {
		id   int
		want bool
	}{
		{
			2,     //テストで作ったスペース
			false, //エラーはでないはず
		},
		{
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
