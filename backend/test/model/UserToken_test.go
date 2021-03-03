package model

import (
	"errors"
	"main/model"
	"testing"
	"time"
)

var utm = model.NewUserTokenModel("test")

//UserModel.TableName()のテスト
func TestTableNameForUserTokenModel(t *testing.T) {
	want := "user_tokens"
	tableName := utm.TableName()
	if tableName != want {
		t.Errorf("UserModel.TableName()の値が異常です。TableName()の出力結果: %s", tableName)
	}

}

//UserTokenModel.Create()のテスト
func TestCreateUserToken(t *testing.T) {

	tests := []struct {
		in   model.UserToken
		want error
	}{
		{
			//①: 存在するユーザのトークン
			model.UserToken{
				UserId: 1,
				Token:  "test token",
				Expire: time.Now().Add(24 * time.Hour),
			},
			nil, //エラーにはならないはず
		},
		{
			//②: 有効期限が現在時刻までのトークン(後ほどテストで使用する)
			model.UserToken{
				UserId: 2,
				Token:  "test token2",
				Expire: time.Now(),
			},
			nil, //エラーにはならないはず
		},
		{
			//②: 存在しないユーザのトークン
			model.UserToken{
				UserId: 9999,
				Token:  "test token",
				Expire: time.Now().Add(24 * time.Hour),
			},
			errors.New("ユーザが存在しません。"), //エラーになるはず
		},
	}
	for i, tt := range tests {
		_, err := utm.Create(tt.in)
		if err != tt.want {
			if err.Error() != tt.want.Error() {
				t.Errorf("%d番目のテストが失敗しました。期待した結果:%s 実際の結果:%s", i+1, tt.want.Error(), err.Error())
			}
		}
	}
	//有効期限を意図的に切れさせるために1秒待つ
	time.Sleep(time.Second * 1)
}

//UserTokenModel.GetByToken()のテスト
func TestGetUserTokenByToken(t *testing.T) {

	tests := []struct {
		in   string
		want error
	}{
		{
			//①: 存在するトークン
			"test token",
			nil, //エラーにはならないはず
		},
		{
			//②: 存在しないトークン
			"Invalid Token",
			errors.New("トークンが見つかりません。"), //エラーになるはず
		},
	}
	for i, tt := range tests {
		_, err := utm.GetByToken(tt.in)
		if err != tt.want {
			if err.Error() != tt.want.Error() {
				t.Errorf("%d番目のテストが失敗しました。期待した結果:%s 実際の結果:%s", i+1, tt.want.Error(), err.Error())
			}
		}
	}
}

//UserTokenModel.IsValid()のテスト
func TestIsValidUserToken(t *testing.T) {

	tests := []struct {
		in   string
		want error
	}{
		{
			//①: 存在するトークン
			"test token",
			nil, //エラーにはならないはず
		},
		{
			//②: 存在しないトークン
			"Invalid Token",
			errors.New("トークンが見つかりません。"), //エラーになるはず
		},
		{
			//③: 有効期限切れのトークン
			"test token2",
			errors.New("有効期限が切れています。"), //エラーになるはず
		},
	}
	for i, tt := range tests {
		_, err := utm.IsValid(tt.in)
		if err != tt.want {
			if err.Error() != tt.want.Error() {
				t.Errorf("%d番目のテストが失敗しました。期待した結果:%s 実際の結果:%s", i+1, tt.want.Error(), err.Error())
			}
		}
	}
}

//UserTokenModel.Delete()のテスト
func TestDeleteUserToken(t *testing.T) {

	tests := []struct {
		in   string
		want error
	}{
		{
			//①: 存在するトークン
			"test token",
			nil, //エラーにはならないはず
		},
		{
			//②: 存在しないトークン
			"Invalid Token",
			errors.New("削除するトークンが存在しません。"), //エラーになるはず
		},
	}
	for i, tt := range tests {
		err := utm.Delete(tt.in)
		if err != tt.want {
			if err.Error() != tt.want.Error() {
				t.Errorf("%d番目のテストが失敗しました。期待した結果:%s 実際の結果:%s", i, tt.want.Error(), err.Error())
			}
		}
	}
}
