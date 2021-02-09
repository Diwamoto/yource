package model

import (
	//標準ライブラリ

	"strconv"
	"time"

	//自作ライブラリ
	"main/config/database"

	//githubライブラリ
	"github.com/go-playground/validator"
)

//チャンネルモデル 一つのスペースは複数のチャンネルを持つ
type Channel struct {
	Entity
	SpaceId int `validate:"required"`
	//UserIdはスペースが持っているので保持する必要なし。
	Name        string `validate:"required"`
	Description string
}

//呼び出し用チャンネルモデル
//AppModelを埋め込み
type ChannelModel struct {
	AppModel
}

//チャンネルモデルを取得する
func NewChannelModel(t string) *ChannelModel {
	var cm ChannelModel
	cm.db = database.GetInstance(t)
	cm.nc = t
	return &cm
}

func (ChannelModel) TableName() string {
	return "channels"
}

//バリデーションをかける
//文字の整形系はフロントで行うので
//最低限の入力チェックのみをgoで行う
func (cm ChannelModel) Validate(c Channel) ([]string, bool) {

	validate := validator.New()
	err := validate.Struct(c)
	var messages []string
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.Field()
			switch fieldName {
			case "SpaceId":
				messages = append(messages, "スペースIDを入力してください")
			case "Name":
				messages = append(messages, "名前を入力してください")
			case "Description":
				messages = append(messages, "説明を入力してください")
			}
		}
	}
	//存在するスペースIDのみを使用できる
	sm := NewSpaceModel(cm.nc)
	_, err2 := sm.GetById(c.SpaceId)
	if err2 {
		messages = append(messages, "存在しないスペースIDのチャンネルは作成できません。")
	}

	if len(messages) > 0 {
		return messages, true
	} else {
		return []string{}, false
	}

}

//チャンネルを作成する
func (cm ChannelModel) Create(c Channel) ([]string, bool) {

	cm.db.AutoMigrate(&c)

	msg, err := cm.Validate(c)

	if !err {
		c.Created = time.Now()
		c.Modified = time.Now()
		//バリデーションが通れば作成し、メッセージの中に作成したチャンネルIDを入れて返す
		cm.db.Create(&c)
		msg = append(msg, strconv.Itoa(int(c.Id)))
		return msg, false
	} else {
		//作成できなければエラーメッセージを返す
		return msg, true
	}

}

//全てのチャンネルを取得する
func (cm ChannelModel) GetAll() ([]Channel, bool) {

	var c []Channel

	cm.db.AutoMigrate(&c)
	cm.db.Find(&c)

	//値が取得できたら
	if len(c) > 0 {
		return c, false
	} else {
		return []Channel{}, true
	}

}

//指定チャンネルidの情報を返す
func (cm ChannelModel) GetById(id int) (Channel, bool) {

	var c Channel
	cm.db.First(&c, id)

	//値が取得できたら
	if c.Id == id {
		return c, false
	} else {
		return Channel{}, true
	}

}

//チャンネルをスペースIDで検索する
//スペースに対してチャンネルは複数存在する
func (sm ChannelModel) GetBySpaceId(spaceId int) ([]Channel, bool) {

	var c []Channel

	sm.db.AutoMigrate(&c)
	sm.db.Where("space_id = ?", spaceId).Find(&c)

	//値が取得できたら
	if len(c) > 0 {
		return c, false
	} else {
		return []Channel{}, true
	}

}

//更新メソッド
//チャンネルの情報を更新する
func (cm ChannelModel) Update(id int, c Channel) ([]string, bool) {
	var tc Channel

	cm.db.AutoMigrate(&tc)
	cm.db.First(&tc, id)

	//引数のチャンネルの情報を移す
	//ここでは変更の検知のみ
	if c.Name != "" {
		tc.Name = c.Name
	}
	if c.Description != "" {
		tc.Description = c.Description
	}

	//更新日を現在にする
	tc.Modified = time.Now()

	//バリデーションをかける
	msg, err := cm.Validate(tc)

	//バリデーションが成功していたら
	if !err {
		//セーブした結果がエラーであれば更新失敗
		if result := cm.db.Save(&tc); result.Error != nil {
			return []string{"データベースに保存することができませんでした。"}, true
		} else {
			return []string{}, false
		}
	} else {
		//バリデーションが失敗していたらそのエラーメッセージを返す
		return msg, true
	}

}

//削除メソッド
//チャンネルを削除する
func (cm ChannelModel) Delete(id int) ([]string, bool) {

	//idで削除を実行する
	_, err := cm.GetById(id)
	if err { //削除するチャンネルがいなかったらダメ
		return []string{"削除するチャンネルが存在しません。"}, true
	}
	cm.db.Delete(&Channel{}, id)
	_, err2 := cm.GetById(id)
	if err2 { //チャンネルが取得できなかったら成功
		return []string{"削除に成功しました。"}, false
	} else {
		return []string{"削除できませんでした。"}, true
	}

}
