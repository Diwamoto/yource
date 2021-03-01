package config

var config = map[string]interface{}{
	//任意Configを定義できる

	//ルートのパス
	"rootPath": "/go/src/github.com/Diwamoto/yource",

	//新規登録時のメール死活確認のトークン有効期限(日)
	"expireToken": "3",

	//dbのデバッグの有無 (true →全てのsqlがログで流れます)
	"debugDB": false,
}

//configの中から特定の値を取得
func Get(key string) interface{} {
	if config[key] != nil {
		return config[key]
	} else {
		return nil
	}
}
