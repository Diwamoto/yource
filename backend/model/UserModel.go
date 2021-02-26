package model

import (

	//標準ライブラリ

	"encoding/json"
	"errors"
	"time"

	//自作ライブラリ
	"main/config/database"

	//githubライブラリ
	"github.com/go-playground/validator"
)

//UserEntity　Entityを埋め込まれている
type User struct {
	Entity
	Email    string `validate:"required,email"`
	Password string //フロントで弾いてhash化された物が入るイメージ、不正にデータが作られた場合はログインできない為問題ない
	Name     string
	Nickname string
	Phone    string
	Status   int //1: 有効, 2: 無効
	Profile  UserProfile
}

//呼び出し用ユーザモデル
//AppModelを埋め込み
type UserModel struct {
	AppModel
}

func NewUserModel(t string) *UserModel {
	var um UserModel
	um.db = database.GetInstance(t)
	um.nc = t
	return &um
}

func (um UserModel) TableName() string {
	return "users"
}

//バリデーションをかける
//文字の整形系はフロントで行うので
//最低限の入力チェックのみをgoで行う
//返す文字列はエラーの文字配列をjson型にしたもの
func (um UserModel) Validate(u User) (string, bool) {

	validate := validator.New()
	err := validate.Struct(u)
	var messages []string

	//独自バリデーション
	//メールアドレスをdbに問い合わせて存在していたらエラーを返す。
	if !um.validateIsUniqueEmail(u.Email) {
		//作成できなければエラーメッセージを返す
		messages = append(messages, "入力されたメールアドレスは既に登録されています。")
	}

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fieldName := err.Field()
			switch fieldName {
			case "Email":
				var typ = err.Tag() //バリデーションでNGになったタグ名を取得
				switch typ {
				case "required":
					messages = append(messages, "メールアドレスを入力してください。")
				case "email":
					messages = append(messages, "正しいメールアドレスを入力してください。")
				}
				//emailとパスワードで登録させるためにいったん名前と電話番号のvalidationを外す
				// case "Name":
				// 	messages = append(messages, "名前を入力してください。")
				// case "Phone":
				// 	messages = append(messages, "電話番号を入力してください。"
			}
		}
	}

	if len(messages) > 0 {
		msgJson, _ := json.Marshal(messages)
		return string(msgJson), true
	} else {
		return "[]", false
	}

}

//ユーザを作成する
func (um UserModel) Create(u User) (User, error) {

	um.db.AutoMigrate(&u)

	msg, err := um.Validate(u)

	if !err {

		u.Created = time.Now()
		u.Modified = time.Now()
		//バリデーションが通れば作成し、メッセージの中に作成したユーザIDを入れて返す
		um.db.Create(&u)
		return u, nil
	} else {
		//作成できなければエラーメッセージを返す
		return User{}, errors.New(msg)
	}

}

//指定ユーザidの情報を返す
//Idで検索するFind()のラッパー
func (um UserModel) GetById(id int) ([]User, error) {

	var u User
	u.Id = id
	result, err := um.Find(u)
	return result, err

}

//検索メソッド
func (um UserModel) Find(u User) ([]User, error) {

	var r []User
	//受け取った検索パラメータに応じてSQLをビルドする

	builder := um.db.Model(&User{})

	//検索パラメータを解析し、検索条件が存在していればwhere文を追加する
	if u.Id != 0 {
		builder = builder.Where("id = ?", u.Id)
	}
	if u.Email != "" {
		builder = builder.Where("email = ?", u.Email)
	}
	if u.Name != "" {
		builder = builder.Where("name LIKE ?", "%"+u.Name+"%")
	}
	if u.Nickname != "" {
		builder = builder.Where("nickname LIKE ?", "%"+u.Nickname+"%")
	}
	if u.Phone != "" {
		builder = builder.Where("phone LIKE ?", u.Phone)
	}
	if u.Status != 0 {
		builder = builder.Where("status = ?", u.Status)
	}

	//dbに問い合わせる。何らかのエラーが発生した場合はここでハンドリング
	if result := builder.Find(&r); result.Error != nil {
		//何らかのエラーを返す
		//テストで発生させることができずカバレッジが取れない。。。
		return []User{}, result.Error
	} else {
		upm := NewUserProfileModel(um.nc)
		for index, user := range r {
			r[index].Profile, _ = upm.GetByUserId(user.Id)
		}
		return r, nil
	}
}

//更新メソッド
//ユーザの情報を更新する
func (um UserModel) Update(id int, u User) (User, error) {
	var tu User

	um.db.AutoMigrate(&tu)
	um.db.First(&tu, id)

	//引数のユーザの情報を移す
	//ここでは変更の検知のみ
	if u.Email != "" {
		tu.Email = u.Email
	}
	if u.Name != "" {
		tu.Name = u.Name
	}
	if u.Nickname != "" {
		tu.Nickname = u.Nickname
	}
	if u.Password != "" {
		tu.Password = u.Password
	}
	if u.Phone != "" {
		tu.Phone = u.Phone
	}
	if u.Status != tu.Status {
		tu.Status = u.Status
	}

	//更新日を現在にする
	tu.Modified = time.Now()

	//バリデーションをかける
	msg, err := um.Validate(tu)

	//バリデーションが成功していたら
	if !err {
		um.db.Save(&tu)
		return tu, nil

		// //セーブした結果がエラーであれば更新失敗
		// if result := um.db.Save(&tu); result.Error != nil {
		// 	return []string{"データベースに保存することができませんでした。"}, true
		// } else {
		// 	return []string{}, false
		// }
	} else {
		//バリデーションが失敗していたらそのエラーメッセージを返す
		return User{}, errors.New(msg)
	}

}

//削除メソッド
//ユーザを削除する
func (um UserModel) Delete(id int) error {

	//idで削除を実行する
	users, _ := um.GetById(id)
	if len(users) == 0 { //削除するユーザがいなかったらダメ
		return errors.New("削除するユーザが存在しません。")
	}
	um.db.Delete(&User{}, id)
	return nil

}

//独自バリデーション
//同じメールアドレスがdbに存在しないかどうかを検索する。
func (um UserModel) validateIsUniqueEmail(email string) bool {
	u := User{
		Email: email,
	}
	users, _ := um.Find(u)
	if len(users) > 0 {
		return false
	} else {
		return true
	}
}
