package config

import (
	"main/config"
	"testing"
)

//config.Get()のテスト
func TestGetConfig(t *testing.T) {
	tests := []struct {
		in   string
		want interface{}
	}{
		{
			//①: rootPathを取得
			"rootPath",
			"/go/src/github.com/Diwamoto/yource", //取得できるはず
		},
		{
			//②: expireTokenを取得
			"expireToken",
			"3", //取得できるはず
		},
		{
			//③: debugモードの取得
			"debugDB",
			false, //取得できるはず
		},
		{
			//④: 未定義configの取得
			"SomethingNotInConfig",
			nil, //取得できないはず
		},
		{
			//⑤: 未定義configの取得その二
			"CantFindString",
			nil, //取得できないはず
		},
		{
			//⑥: sendGridの送信者名を取得
			"sendGridFrom",
			"info@yource.space", //取得できるはず
		},
	}
	for i, tt := range tests {
		result := config.Get(tt.in)
		if result != tt.want {
			t.Errorf("%d番目のテストが失敗しました。Get(%s)の想定：%s, 結果：%#v ", i+1, tt.in, tt.want, result)
		}
	}
}
