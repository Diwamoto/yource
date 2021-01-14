package config

//標準ライブラリ

//自作ライブラリ

//githubライブラリ
type Config struct {
	Key   string
	Value string
}

var config = []Config{
	//任意Configを定義できる
	{"test", "TestValue"},
}

//configを取得
func GetConfig() []Config {
	return config
}
