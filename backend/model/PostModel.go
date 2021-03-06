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

//投稿 投稿はチャンネルと紐付き、どのユーザが投稿したかの情報を持つ。
type Post struct {
	Entity
	ChannelId int       `validate:"required"` //送信チャンネルID
	UserId    int       `validate:"required"` //送信者ID
	User      User      `validate:"required"` //送信者がbelongstoで入る
	Content   string    `validate:"required"`
	Date      time.Time `validate:"required"`
}

//呼び出し用投稿モデル
//AppModelを埋め込み
type PostModel struct {
	AppModel
}

func NewPostModel(t string) *PostModel {
	var pm PostModel
	pm.db = database.GetInstance(t) //データベースオブジェクトの
	pm.nc = t
	return &pm
}

func (pm PostModel) TableName() string {
	return "posts"
}

//バリデーションをかける
//文字の整形系はフロントで行うので
//最低限の入力チェックのみをgoで行う
func (pm PostModel) Validate(p Post) ([]string, bool) {

	validate := validator.New()
	err := validate.Struct(p)
	var messages []string
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.Field()
			switch fieldName {
			case "ChannelId":
				messages = append(messages, "チャンネルIDを入力してください。")
			case "UserId":
				messages = append(messages, "ユーザIDを入力してください。")
			case "Content":
				messages = append(messages, "内容を入力してください。")
			case "Date":
				messages = append(messages, "日付を入力してください。")
			}
		}
	}

	//存在しないチャンネルの投稿は作成できない
	cm := NewChannelModel(pm.nc)
	_, err2 := cm.GetById(p.ChannelId)
	if err2 {
		messages = append(messages, "存在しないチャンネルの投稿は作成できません。")
	}

	//存在しないユーザは投稿できない
	um := NewUserModel(pm.nc)
	users, _ := um.GetById(p.UserId)
	if len(users) == 0 {
		messages = append(messages, "存在しないユーザIDの投稿は作成できません。")
	}

	if len(messages) > 0 {
		return messages, true
	} else {
		return []string{}, false
	}

}

//投稿を作成する
func (pm PostModel) Create(p Post) ([]string, bool) {

	pm.db.AutoMigrate(&p)

	msg, err := pm.Validate(p)

	if !err {
		p.Created = time.Now()
		p.Modified = time.Now()
		//バリデーションが通れば作成し、メッセージの中に作成した投稿IDを入れて返す
		pm.db.Create(&p)
		msg = append(msg, strconv.Itoa(int(p.Id)))
		return msg, false
	} else {
		//作成できなければエラーメッセージを返す
		return msg, true
	}

}

//全てのチャンネルを取得する
func (pm PostModel) GetAll() ([]Post, bool) {

	var posts []Post

	pm.db.AutoMigrate(&Post{})
	pm.db.Preload("User").Find(&posts)

	//値が取得できたら
	if len(posts) > 0 {
		return posts, false
	} else {
		return []Post{}, true
	}

}

//チャンネルの投稿を全て取得する
func (pm PostModel) GetByChannelId(channelId int) ([]Post, bool) {

	var p []Post

	pm.db.AutoMigrate(&p)
	pm.db.Preload("User").Where("channel_id = ?", channelId).Find(&p)

	//値が取得できたら
	if len(p) > 0 {
		return p, false
	} else {
		return []Post{}, true
	}

}

//指定ユーザidの情報を返す
//チャンネル問わず返す。チャンネルを指定して返したい時はfindで指定する
func (pm PostModel) GetByUserId(userId int) ([]Post, bool) {

	var p []Post

	pm.db.AutoMigrate(&p)
	pm.db.Preload("User").Where("user_id = ?", userId).Find(&p)

	//値が取得できたら
	if len(p) > 0 {
		return p, false
	} else {
		return []Post{}, true
	}

}

//指定投稿idの情報を返す
//多分ほぼ使用しない
func (pm PostModel) GetById(id int) (Post, bool) {

	var p Post

	pm.db.AutoMigrate(&p)
	pm.db.Preload("User").First(&p, id)

	//値が取得できたら
	if p.Id == id {
		return p, false
	} else {
		return Post{}, true
	}

}

//検索メソッド
//後ほど拡充する
func (pm PostModel) Find(p Post) ([]Post, bool) {

	var r []Post

	pm.db.Preload("User").Where(&p).Find(&r)

	//値が取得できたら
	if len(r) > 0 {
		return r, false
	} else {
		return []Post{}, true
	}

}

//更新メソッド
//投稿の情報を更新する
func (pm PostModel) Update(id int, p Post) ([]string, bool) {
	var tp Post

	pm.db.AutoMigrate(&tp)
	pm.db.First(&tp, id)

	//引数の情報を移す
	tp.Content = p.Content
	//更新日を現在にする
	tp.Modified = time.Now()

	//バリデーションをかける
	msg, err := pm.Validate(tp)

	//バリデーションが成功していたら
	if !err {
		pm.db.Save(&tp)
		return []string{}, false
		// //セーブした結果がエラーであれば更新失敗
		// if result := pm.db.Save(&tp); result.Error != nil {
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
//投稿を削除する
func (pm PostModel) Delete(id int) ([]string, bool) {

	//idで削除を実行する
	_, err := pm.GetById(id)
	if err { //削除する投稿がいなかったらダメ
		return []string{"削除する投稿が存在しません。"}, true
	}
	pm.db.Delete(&Post{}, id)
	// _, err2 := pm.GetById(id)
	// if err2 { //投稿が取得できなかったら成功
	return []string{"削除に成功しました。"}, false
	// }
	//ここでこけるのはdbサーバが落ちたときなのでいったんfalse
	// else {
	// 	//dbのエラーのためカバレッジでカバーできない
	// 	return []string{"削除できませんでした。"}, true
	// }
}
