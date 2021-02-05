package model

import (

	//標準パッケージ
	"regexp"
	"strconv"
	"time"

	//自作パッケージ
	"main/config/database"

	//githubパッケージ
	"github.com/go-playground/validator"
)

//スペース（ブログ）構造体 ユーザそれぞれが作ることのできる空間。
type Space struct {
	Entity
	UserId      int    `validate:"required"`
	Name        string `validate:"required"`
	Discription string
	SubDomain   string `validate:"required"`
	Status      bool
	Publish     bool //boolなので初期値はfalse(非公開)→バリデーション不要
}
type SpaceModel struct {
	AppModel
}

func NewSpaceModel(t string) *SpaceModel {
	var sm SpaceModel
	sm.db = database.GetInstance(t)
	sm.nc = t
	return &sm
}

func (Space) TableName() string {
	return "spaces"
}

//バリデーションをかける
//文字の整形系はフロントで行うので
//最低限の入力チェックのみをgoで行う
func (sm SpaceModel) Validate(s Space) ([]string, bool) {

	validate := validator.New()
	err := validate.Struct(s)
	var messages []string
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.Field()
			switch fieldName {
			case "UserId":
				messages = append(messages, "ユーザidを入力してください。")
			case "Name":
				messages = append(messages, "スペース名を入力してください。")
			case "SubDomain":
				messages = append(messages, "サブドメインを入力してください。")
			}

		}
	}

	//ユーザIDは存在するユーザのIDのみを使用できる
	um := NewUserModel(sm.nc)
	_, err2 := um.GetById(s.UserId)
	if err2 {
		messages = append(messages, "存在しないユーザIDを持つスペースは作成できません。")
	}

	//正規表現チェックを追加
	//go-playground/validatorには正規表現チェックがないため。
	//参考：https://godoc.org/gopkg.in/go-playground/validator.v9
	if !regexp.MustCompile(`^[0-9a-zA-Z]+$`).Match([]byte(s.SubDomain)) {
		messages = append(messages, "サブドメインに半角英字以外の文字は使えません。")
	}

	if len(messages) > 0 {
		return messages, true
	} else {
		return []string{}, false
	}

}

//スペースを作成する
func (sm SpaceModel) Create(s Space) ([]string, bool) {

	sm.db.AutoMigrate(&s)

	//バリデーションが通れば作成し、メッセージの中に作成したスペースIDを入れて返す
	msg, err := sm.Validate(s)
	if !err {
		s.Created = time.Now()
		s.Modified = time.Now()

		sm.db.Create(&s)
		msg = append(msg, strconv.Itoa(int(s.Id)))
		return msg, false
	} else {
		//作成できなければエラーメッセージを返す
		return msg, true
	}

}

//指定スペースidの情報を返す
func (sm SpaceModel) GetAll() ([]Space, bool) {

	var s []Space

	sm.db.AutoMigrate(&s)
	sm.db.Find(&s)

	//値が取得できたら
	if len(s) > 0 {
		return s, false
	} else {
		return []Space{}, true
	}

}

//指定スペースidの情報を返す
func (sm SpaceModel) GetById(id int) (Space, bool) {

	var s Space

	sm.db.AutoMigrate(&s)
	sm.db.First(&s, id)

	//値が取得できたら
	if s.Id == id {
		return s, false
	} else {
		return Space{}, true
	}

}

//指定スペースidの情報を返す
func (sm SpaceModel) GetByUserId(userId int) (Space, bool) {

	var s Space

	sm.db.AutoMigrate(&s)
	sm.db.Where("user_id = ?", userId).First(&s)

	//値が取得できたら
	if s.UserId == userId {
		return s, false
	} else {
		return Space{}, true
	}

}

//指定スペースidの情報を返す
func (sm SpaceModel) Find(s Space) ([]Space, bool) {

	var r []Space

	sm.db.Where(&s).Find(&r)

	//値が取得できたら
	if len(r) > 0 {
		return r, false
	} else {
		return []Space{}, true
	}

}

//更新メソッド
//スペースの情報を更新する
func (sm SpaceModel) Update(id int, s Space) ([]string, bool) {
	var ts Space

	sm.db.AutoMigrate(&ts)
	sm.db.First(&ts, id)

	//引数のスペースの情報を移す
	//ここでは変更の検知のみ
	//ユーザIDは変更することができない
	ts.Name = s.Name
	ts.Discription = s.Discription
	ts.SubDomain = s.SubDomain
	ts.Status = s.Status
	ts.Publish = s.Publish
	//更新日を現在にする
	ts.Modified = time.Now()

	//バリデーションをかける
	msg, err := sm.Validate(ts)

	//バリデーションが成功していたら
	if !err {
		//セーブした結果がエラーであれば更新失敗
		if result := sm.db.Save(&ts); result.Error != nil {
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
//スペースを削除する
func (sm SpaceModel) Delete(id int) ([]string, bool) {

	//idで削除を実行する
	_, err := sm.GetById(id)
	if err { //削除するスペースがいなかったらダメ
		return []string{"削除するスペースが存在しません。"}, true
	}
	sm.db.Delete(&Space{}, id)
	_, err2 := sm.GetById(id)
	if err2 { //スペースが取得できなかったら成功
		return []string{"削除に成功しました。"}, false
	} else {
		return []string{"削除できませんでした。"}, true
	}

}
