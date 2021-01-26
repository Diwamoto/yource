package config

var config = map[string]string{
	//任意Configを定義できる
	"rootPath": "/go/src/github.com/Diwamoto/yource",
}

//configの中から特定の値を取得
func Get(key string) string {
	if config[key] != "" {
		return config[key]
	} else {
		return ""
	}
}
