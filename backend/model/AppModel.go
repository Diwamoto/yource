package model

import "time"

//AppModel
//モデル構造体は全てこの構造体を埋め込む。
type AppModel struct {
	Id       int `gorm:"primary_key"`
	Created  time.Time
	Modified time.Time
}

//Modelインタフェース
//全てのモデル構造体で同じ名前の関数を使うため
