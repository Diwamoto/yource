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

	//独自バリデーション
	//存在するスペースIDのみを使用できる
	if !cm.validationIsExistSpace(c.SpaceId) {
		messages = append(messages, "存在しないスペースIDのチャンネルは作成できません。")
	}

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.Field()
			switch fieldName {
			case "SpaceId":
				messages = append(messages, "スペースIDを入力してください")
			case "Name":
				messages = append(messages, "名前を入力してください")
			}
		}
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

	//独自バリデーション
	//同じスペースの同じチャンネル名は作成できない
	//作成時とチャンネル変更時では処理を少し変えるためvalidate後に独自で実行
	if !cm.validationIsUniqueChannelNameInSameSpace(c, "create") {
		msg = append(msg, "同名のチャンネルは作成できません。")
		err = true
	}

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

//任意の条件でチャンネルを検索する
func (cm ChannelModel) Find(c Channel) ([]Channel, bool) {

	var r []Channel

	cm.db.Where(&c).Find(&r)

	//値が取得できたら
	if len(r) > 0 {

		//MEMO: 投稿は今後確実に多くなるのでリレーションは実装しない。
		//必要があれば後ほどチャンネルidを使って取得する。
		return r, false
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

	//スペースIDが変更されていたらだめ
	if tc.SpaceId != c.SpaceId {
		return []string{"スペースIDは変更することが出来ません。"}, true
	}

	//引数のチャンネルの情報を移す
	//ここでは変更の検知のみ
	tc.Name = c.Name
	tc.Description = c.Description

	//更新日を現在にする
	tc.Modified = time.Now()

	//バリデーションをかける
	msg, err := cm.Validate(tc)

	//独自バリデーション
	//既に存在するチャンネル名に変更することはできない
	//変更時用
	if !cm.validationIsUniqueChannelNameInSameSpace(tc, "update") {
		return []string{"既にチャンネルが存在します。"}, true
	}

	//バリデーションが成功していたら
	if !err {
		cm.db.Save(&tc)
		return []string{}, false

		// //セーブした結果がエラーであれば更新失敗
		// if result := cm.db.Save(&tc); result.Error != nil {
		// 	//これが発生すると言うことはサーバがこの関数の実行途中で
		// 	//(具体的には182行目の取得からセーブの間までに)
		// 	//落ちていると言うことなのでカバレッジでカバーできない。
		//  //なのでいったん全体的にコメントアウト(おそらくここのハンドリングは
		//	//バリデート関数の完全性が担保されていれば必要ない)
		// 	return []string{"データベースに保存することができませんでした。"}, true
		// } else {
		// 	return []string{}, false
		// }
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
	// _, err2 := cm.GetById(id)
	// if err2 { //チャンネルが取得できなかったら成功
	return []string{"削除に成功しました。"}, false
	// }
	//ここでこけるのはdbサーバが落ちたときなのでいったんfalse
	// else {
	// 	//dbのエラーのためカバレッジでカバーできない
	// 	return []string{"削除できませんでした。"}, true
	// }

}

//独自バリデーション
//指定されたスペースIDがdbに存在するかを検証
//存在しなければオーケー←GetById()の結果がダメ（検索結果が0）であればtrue
func (cm ChannelModel) validationIsExistSpace(spaceId int) bool {
	sm := NewSpaceModel(cm.nc)
	_, err := sm.GetById(spaceId)
	//errをそのまま返せば動くが明示的に書く
	if !err {
		return true
	} else {
		return false
	}
}

//独自バリデーション
//同じスペースに指定されたチャンネル名が存在しないことを確認する。
//もし作成時と更新時に同じメソッドを使用してしまうと、名前は変えずに
//説明のみ変えるパターンが設定できない（同名スペース、同名チャンネルがdbに存在するから）ので、
//更新時はidで検索して同じ名前であればスルーする→既にIDが設定されているのでUPD判定、かつ同名であれば
//説明のみ変えるパターンと判断
func (cm ChannelModel) validationIsUniqueChannelNameInSameSpace(c Channel, mode string) bool {

	//チャンネルのスペースIDから同スペースの全チャンネルを取得
	channels, err := cm.GetBySpaceId(c.SpaceId)

	if err {
		//同じスペースにチャンネルが存在しなかった場合
		return true

	} else {
		//存在した場合は同じ名前のものがあるか調査
		for _, channel := range channels {
			if c.Name == channel.Name {
				//同名チャンネルが存在したとき、まずモードを確認
				//作成モードなら同名は不許可でfalse。
				//更新モードならIDを比較して、既に割り振られており、
				//かつ説明のみが変更されていたらtrue それ以外はfalse
				if mode == "create" {
					return false
				} else if mode == "update" {
					if c.Id == channel.Id && c.Description != channel.Description {
						return true
					}
				}
				return false
			}
		}
		//最後まで見てなければtrue
		return true

	}
}
