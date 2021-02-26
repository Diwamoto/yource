package model

import (
	//標準ライブラリ

	"errors"
	"testing"
	"time"

	//自作ライブラリ
	"main/model"
	//githubライブラリ
)

var upm = model.NewUserProfileModel("test")

//UserProfileModel.TableName()のテスト
func TestTableNameForUserProfileModel(t *testing.T) {
	want := "user_profiles"
	tableName := upm.TableName()
	if tableName != want {
		t.Errorf("UserProfileModel.TableName()の値が異常です。TableName()の出力結果: %s", tableName)
	}

}

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
				Icon:      "test url",
				Birthday:  time.Now(),
				Hometown:  "japan",
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
				Icon:      "",
				Birthday:  time.Now(),
				Hometown:  "",
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
		want error
	}{
		{
			//①: 正しいプロフィール
			model.UserProfile{
				UserId:    1,
				Profile:   "profile test",
				Icon:      "test url",
				Birthday:  time.Date(2020, 1, 1, 12, 0, 0, 0, time.Local),
				Hometown:  "japan",
				Job:       "engineer",
				Twitter:   "@aaa",
				Facebook:  "my awesome facebook",
				Instagram: "@myinsta",
				Other:     "my.awesome.web.com",
			},
			nil, //エラーはでないはず
		},
		{
			//②: 同じユーザidのプロフィール
			model.UserProfile{
				UserId:    1, //既に作成してしまっている
				Profile:   "profile test",
				Icon:      "test url",
				Birthday:  time.Date(2020, 1, 1, 12, 0, 0, 0, time.Local),
				Hometown:  "japan",
				Job:       "engineer",
				Twitter:   "@aaa",
				Facebook:  "my awesome facebook",
				Instagram: "@myinsta",
				Other:     "my.awesome.web.com",
			},
			errors.New("既に指定ユーザIdのプロフィールが登録されています。"), //エラーになるはず
		},
		{
			//②: 存在しないユーザIDのプロフィール
			model.UserProfile{
				UserId:    9999, //存在しない
				Profile:   "profile test",
				Icon:      "test url",
				Birthday:  time.Date(2020, 1, 1, 12, 0, 0, 0, time.Local),
				Hometown:  "japan",
				Job:       "engineer",
				Twitter:   "@aaa",
				Facebook:  "my awesome facebook",
				Instagram: "@myinsta",
				Other:     "my.awesome.web.com",
			},
			errors.New("存在しないユーザIDのプロフィールは作成できません。"), //エラーになるはず
		},
	}
	for i, tt := range tests {
		_, err := upm.Create(tt.in)
		if err != tt.want {
			if err.Error() != tt.want.Error() {
				t.Errorf("%d番目のテストが失敗しました。エラーメッセージ:%#v", i+1, err)
			}
		}
	}

}

//UserProfileModel.GetById()のテスト
//プロフィールが取得できたらOK,できなければダメ
func TestGetUserProfile(t *testing.T) {

	tests := []struct {
		in   int //UserProfileID
		want error
	}{
		{
			//①: 先ほど作成したプロフィール
			1,
			nil, //エラーはでないはず
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
		want error
	}{
		{
			//①: テストで作成したユーザIDのプロフィール
			1,
			nil, //エラーはでないはず
		},
		{
			//②: 存在しないユーザIDのプロフィール
			9999,
			errors.New("指定ユーザIDのプロフィールは存在しません。"), //エラーになるはず
		},
	}
	for i, tt := range tests {
		result, err := upm.GetById(tt.in)
		if err != tt.want && result.Id != 0 {
			if err.Error() != tt.want.Error() {
				t.Errorf("%d番目のテストが失敗しました。エラーメッセージ:%#v", i+1, err)
			}
		} else {

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
		want error
	}{
		{
			//①: プロフィールの文言で検索
			//TODO: like検索の追加
			model.UserProfile{
				Profile: "profile test",
			},
			"プロフィール",
			nil, //検索は成功するはず
		},
		{
			//②: Iconで検索
			model.UserProfile{
				Icon: "test url",
			},
			"アイコン",
			nil, //検索は成功するはず
		},
		{
			//③: 誕生日で検索
			model.UserProfile{
				Birthday: time.Date(2020, 1, 1, 12, 0, 0, 0, time.Local),
			},
			"誕生日",
			nil, //検索は成功するはず
		},
		{
			//④: 出身地で検索
			//TODO: 出身地の多様性をどうするか
			model.UserProfile{
				Hometown: "japan",
			},
			"出身地",
			nil, //検索は成功するはず
		},
		{
			//⑤: 仕事で検索
			model.UserProfile{
				Job: "engineer",
			},
			"仕事",
			nil, //検索は成功するはず
		},
		{
			//⑥: ツイッターで検索
			model.UserProfile{
				Twitter: "@aaa",
			},
			"ツイッター",
			nil, //検索は成功するはず
		},
		{
			//⑦: facebookで検索
			model.UserProfile{
				Facebook: "my awesome facebook",
			},
			"フェイスブック",
			nil, //検索は成功するはず
		},
		{
			//⑧: インスタグラムで検索
			model.UserProfile{
				Instagram: "@myinsta",
			},
			"インスタグラム",
			nil, //検索は成功するはず
		},
		{
			//⑨: 他のwebサイトで検索
			model.UserProfile{
				Other: "my.awesome.web.com",
			},
			"Other",
			nil, //検索は成功するはず
		},
	}
	for _, tt := range tests {
		_, err := upm.Find(tt.in)
		if err != tt.want {
			t.Errorf("「%s」での検索が失敗しました。%#v", tt.t, err)
		}
	}
}

//UserProfileModel.Update()のテスト
//プロフィールの情報が更新できなかったらダメ
func TestUpdateUserProfile(t *testing.T) {

	tests := []struct {
		id    int
		after model.UserProfile
		want  error
	}{
		{
			//①: 正しい変更内容
			1, //先ほどテストで作ったプロフィール
			model.UserProfile{
				UserId:    1,
				Profile:   "Update",
				Icon:      "Update",
				Birthday:  time.Now(),
				Hometown:  "Update",
				Job:       "Update",
				Twitter:   "Update",
				Facebook:  "Update",
				Instagram: "Update",
				Other:     "Update",
			},
			nil, //エラーはでないはず
		},
		{
			//②: ユーザIDを変更してしまっているユーザ
			1, //先ほどテストで作ったプロフィール
			model.UserProfile{
				UserId:    9999,
				Profile:   "Update",
				Icon:      "Update",
				Birthday:  time.Now(),
				Hometown:  "Update",
				Job:       "Update",
				Twitter:   "Update",
				Facebook:  "Update",
				Instagram: "Update",
				Other:     "Update",
			},
			errors.New("ユーザIDは変更することはできません。"), //ユーザidは変更できない
		},
		{
			//③: 存在しないユーザIDのプロフィール
			9999, //存在しないユーザ
			model.UserProfile{
				UserId:    1,
				Profile:   "Update",
				Icon:      "Update",
				Birthday:  time.Now(),
				Hometown:  "Update",
				Job:       "Update",
				Twitter:   "Update",
				Facebook:  "Update",
				Instagram: "Update",
				Other:     "Update",
			},
			errors.New("指定されたユーザが存在しません。"), //存在しないユーザのプロフィールは変更できない
		},
	}
	for i, tt := range tests {
		_, err := upm.Update(tt.id, tt.after)
		if err != tt.want {
			if err.Error() != tt.want.Error() {
				t.Errorf("%d番目のテストが失敗しました。エラーメッセージ:%#v", i+1, err)
			}
		}
	}

}

//UserProfileModel.Delete()のテスト
func TestDeleteUserProfile(t *testing.T) {

	tests := []struct {
		id   int
		want error
	}{
		{
			//①: 存在するユーザ
			1,   //テストで作ったプロフィール
			nil, //エラーはでないはず
		},
		{
			//②: 存在しないユーザ
			9999999999,
			errors.New("削除するプロフィールが存在しません。"), //存在しないプロフィールは削除できない
		},
	}
	for i, tt := range tests {
		err := upm.Delete(tt.id)
		if err != tt.want {
			if err.Error() != tt.want.Error() {
				t.Errorf("%d番目のテストが失敗しました。エラーメッセージ:%#v", i+1, err)
			}
		}
	}
}
