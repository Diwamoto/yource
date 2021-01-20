package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

//AppModel
//モデル構造体は全てこの構造体を埋め込む。
type AppModel struct {
	nc        string //現在接続中のデータベース名(now connecting)
	TableName string
	db        *gorm.DB
}

//Entity
//全てのモデルのエンティティ構造体に埋め込む。
type Entity struct {
	Id       int `gorm:"primary_key"`
	Created  time.Time
	Modified time.Time
}

//モデルからdbへの接続を止める
func (m *AppModel) Close() {
	m.db.Close()
}
