package model

import (
	//標準ライブラリ
	"testing"
	"time"

	//自作ライブラリ
	"main/model"
	//githubライブラリ
)

var upm = model.NewUserProfileModel("test")

//UserProfileModel.Validate()のテスト
func TestValidateUserProfile(t *testing.T) {

	tests := []struct {
		in   model.UserProfile
		want bool
	}{
		{
			//①: 正しいプロフィール
			model.UserProfile{
				UserId:    1,
				Profile:   "profile test",
				Birthday:  time.Now(),
				From:      "japan",
				Job:       "engineer",
				Twitter:   "@aaa",
				Facebook:  "my awesome facebook",
				Instagram: "@myinsta",
				Other:     "my.awesome.web.com",
			},
			false, //エラーはでないはず
		},
		{
			//②: 存在しないユーザのプロフィール
			model.UserProfile{
				UserId:    9999,
				Profile:   "",
				Birthday:  time.Now(),
				From:      "",
				Job:       "",
				Twitter:   "",
				Facebook:  "",
				Instagram: "",
				Other:     "",
			},
			true, //エラーになるはず
		},
	}
	for i, tt := range tests {
		rs, err := upm.Validate(tt.in)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。ValidateUserProfile()の出力結果: %s", i+1, rs)
		}

	}
}

//UserProfileModel.Create()のテスト
func TestCreateUserProfile(t *testing.T) {

	tests := []struct {
		in   model.UserProfile
		want bool
	}{
		{
			//①: 正しいプロフィール
			model.UserProfile{
				UserId:    1,
				Profile:   "profile test",
				Birthday:  time.Date(2020, 1, 1, 12, 0, 0, 0, time.Local),
				From:      "japan",
				Job:       "engineer",
				Twitter:   "@aaa",
				Facebook:  "my awesome facebook",
				Instagram: "@myinsta",
				Other:     "my.awesome.web.com",
			},
			false, //エラーはでないはず
		},
		{
			//②: 同じユーザidのプロフィール
			model.UserProfile{
				UserId:    2, //既に作成してしまっている
				Profile:   "profile test",
				Birthday:  time.Date(2020, 1, 1, 12, 0, 0, 0, time.Local),
				From:      "japan",
				Job:       "engineer",
				Twitter:   "@aaa",
				Facebook:  "my awesome facebook",
				Instagram: "@myinsta",
				Other:     "my.awesome.web.com",
			},
			true, //エラーになるはず
		},
	}
	for i, tt := range tests {
		rs, err := upm.Create(tt.in)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。出力結果: %s", i+1, rs)
		}
	}

}

//UserProfileModel.GetAll()のテスト
//ユーザが取得できたらOK,できなければダメ
func TestGetAllUserProfile(t *testing.T) {

	tests := []struct {
		want bool
	}{
		{
			//①: 全てのユーザを取得
			false, //取得できるはず
		},
	}
	for _, tt := range tests {
		_, err := um.GetAll()
		if err != tt.want {
			t.Errorf("GetAll()を用いて全プロフィールを取得することができませんでした。")
		}
	}
}

//UserProfileModel.GetById()のテスト
//プロフィールが取得できたらOK,できなければダメ
func TestGetUserProfile(t *testing.T) {

	tests := []struct {
		in   int //UserProfileID
		want bool
	}{
		{
			//①: 先ほど作成したプロフィール
			1,
			false, //エラーはでないはず
		},
	}
	for _, tt := range tests {
		_, err := upm.GetById(tt.in)
		if err != tt.want {
			t.Errorf("UserProfileID:%dのプロフィールを取得できませんでした。", tt.in)
		}
	}
}

//UserProfileModel.GetByUserId()のテスト
//プロフィールが取得できたらOK,できなければダメ
func TestGetUserProfileByUserId(t *testing.T) {

	tests := []struct {
		in   int //userId
		want bool
	}{
		{
			//①: テストで作成したユーザIDのプロフィール
			1,
			false, //エラーはでないはず
		},
		{
			//②: 存在しないユーザIDのプロフィール
			9999,
			true, //エラーになるはず
		},
	}
	for _, tt := range tests {
		_, err := upm.GetById(tt.in)
		if err != tt.want {
			t.Errorf("UserProfileID:%dのプロフィールを取得できませんでした。", tt.in)
		}
	}
}

//UserProfileModel.Find()のテスト
//プロフィールを検索する
//検索の失敗についての定義は議論中
func TestFindUserProfile(t *testing.T) {
	tests := []struct {
		in   model.UserProfile
		t    string //検索の種類
		want bool
	}{
		{
			//①: プロフィールの文言で検索
			//TODO: like検索の追加
			model.UserProfile{
				Profile: "profile test",
			},
			"プロフィール",
			false, //検索は成功するはず
		},
		{
			//②: 誕生日で検索
			model.UserProfile{
				Birthday: time.Date(2020, 1, 1, 12, 0, 0, 0, time.Local),
			},
			"誕生日",
			false, //検索は成功するはず
		},
		{
			//③: 出身地で検索
			//TODO: 出身地の多様性をどうするか
			model.UserProfile{
				From: "japan",
			},
			"出身地",
			false, //検索は成功するはず
		},
		{
			//④: 仕事で検索
			model.UserProfile{
				Job: "engineer",
			},
			"仕事",
			false, //検索は成功するはず
		},
		{
			//⑤: ツイッターで検索
			model.UserProfile{
				Twitter: "@aaa",
			},
			"ツイッター",
			false, //検索は成功するはず
		},
		{
			//⑥: facebookで検索
			model.UserProfile{
				Facebook: "my awesome facebook",
			},
			"フェイスブック",
			false, //検索は成功するはず
		},
		{
			//⑦: インスタグラムで検索
			model.UserProfile{
				Instagram: "@myinsta",
			},
			"インスタグラム",
			false, //検索は成功するはず
		},
		{
			//⑧: 他のwebサイトで検索
			model.UserProfile{
				Other: "my.awesome.web.com",
			},
			"Other",
			false, //検索は成功するはず
		},
		{
			//⑨: プロフィールの文言で検索
			model.UserProfile{
				Profile: "fail test",
			},
			"Profile",
			true, //検索は失敗するはず
		},
	}
	for _, tt := range tests {
		_, err := upm.Find(tt.in)
		if err != tt.want {
			t.Errorf("「%s」での検索が失敗しました。", tt.t)
		}
	}
}

//UserProfileModel.Update()のテスト
//プロフィールの情報が更新できなかったらダメ
func TestUpdateUserProfile(t *testing.T) {

	tests := []struct {
		id    int
		after model.UserProfile
		want  bool
	}{
		{
			//①: 正しい変更内容
			1, //先ほどテストで作ったプロフィール
			model.UserProfile{
				UserId:    1,
				Profile:   "Update",
				Birthday:  time.Now(),
				From:      "Update",
				Job:       "Update",
				Twitter:   "Update",
				Facebook:  "Update",
				Instagram: "Update",
				Other:     "Update",
			},
			false, //エラーはでないはず
		},
		{
			//②: ユーザIDを変更してしまっているユーザ
			1, //先ほどテストで作ったプロフィール
			model.UserProfile{
				UserId:    9999,
				Profile:   "Update",
				Birthday:  time.Now(),
				From:      "Update",
				Job:       "Update",
				Twitter:   "Update",
				Facebook:  "Update",
				Instagram: "Update",
				Other:     "Update",
			},
			true, //ユーザidは変更できない
		},
	}
	for i, tt := range tests {
		msg, err := upm.Update(tt.id, tt.after)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。エラーメッセージ:%s", i+1, msg)
		}
	}

}

//UserProfileModel.Delete()のテスト
func TestDeleteUserProfile(t *testing.T) {

	tests := []struct {
		id   int
		want bool
	}{
		{
			//①: 存在するユーザ
			1,     //テストで作ったプロフィール
			false, //エラーはでないはず
		},
		{
			//②: 存在しないユーザ
			9999999999,
			true, //存在しないプロフィールは削除できない
		},
	}
	for i, tt := range tests {
		msg, err := upm.Delete(tt.id)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。エラーメッセージ:%s", i+1, msg)
		}
	}
}
