package model

//チャンネルモデル 一つのスペースは複数のチャンネルを持つ
type Channel struct {
	Entity
	SpaceId     int
	Name        string
	Discription string
}
