package config

var config = map[string]interface{}{
	//任意Configを定義できる

	//ルートのパス
	"rootPath": "/go/src/github.com/Diwamoto/yource",

	//新規登録時のメール死活確認のトークン有効期限(日)
	"expireToken": "3",

	//dbのデバッグの有無 (true →全てのsqlがログで流れます)
	"debugDB": false,

	//sendGridのメール送信元
	"sendGridFrom": "info@yource.space",
}

//configの中から特定の値を取得
func Get(key string) interface{} {
	if config[key] != nil {
		return config[key]
	} else {
		return nil
	}
}

//configの中から特定の値を取得
func ToString(key string) string {
	if config[key] != nil {
		str, ok := config[key].(string)
		if ok {
			return str
		} else {
			return ""
		}
	} else {
		return ""
	}
}
