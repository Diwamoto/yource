package test_model

import (
	"main/model"
	"testing"
	"time"
)

var upm = model.NewUserProfileModel("test")

//UserProfileModel.Validate()のテスト
func TestValidateUserProfile(t *testing.T) {

	tests := []struct {
		in   model.UserProfile
		want bool
	}{
		{
			//①正しいプロフィール
			model.UserProfile{
				UserId:    1,
				Profile:   "",
				Birthday:  time.Now(),
				From:      "",
				Job:       "",
				Twitter:   "",
				Facebook:  "",
				Instagram: "",
				Other:     "",
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
			//①正しいプロフィール
			model.UserProfile{
				UserId:    1,
				Profile:   "",
				Birthday:  time.Now(),
				From:      "",
				Job:       "",
				Twitter:   "",
				Facebook:  "",
				Instagram: "",
				Other:     "",
			},
			false, //エラーはでないはず
		},
	}
	for i, tt := range tests {
		rs, err := upm.Create(tt.in)
		if err != tt.want {
			t.Errorf("%d番目のテストが失敗しました。の出力結果: %s", i+1, rs)
		}
	}

}

//UserProfileModel.Get()のテスト
//プロフィールが取得できたらOK,できなければダメ
func TestGetUserProfile(t *testing.T) {

	tests := []struct {
		in   int //UserProfileID
		want bool
	}{
		{
			//①先ほど作成したプロフィール
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

//UserProfileModel.Update()のテスト
//プロフィールの情報が更新できなかったらダメ
func TestUpdateUserProfile(t *testing.T) {

	tests := []struct {
		id    int
		after model.UserProfile
		want  bool
	}{
		//①:正しい変更内容
		{
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
		//②:ユーザIDを変更してしまっているユーザ
		{
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
			false, //エラーはでないはず
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
			1,     //テストで作ったプロフィール
			false, //エラーはでないはず
		},
		{
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
