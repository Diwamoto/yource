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

//UserEntity　Entityを埋め込まれている
type User struct {
	Entity
	Email    string `validate:"required,email"`
	Password string //フロントで弾いてhash化された物が入るイメージ、不正にデータが作られた場合はログインできない為問題ない
	Name     string
	Nickname string
	Phone    string
	Status   bool
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
func (um UserModel) Validate(u User) ([]string, bool) {

	validate := validator.New()
	err := validate.Struct(u)
	var messages []string

	//独自バリデーション
	//メールアドレスをdbに問い合わせて存在していたらエラーを返す。
	if !um.validateUniqueEmail(u.Email) {
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
		return messages, true
	} else {
		return []string{}, false
	}

}

//ユーザを作成する
func (um UserModel) Create(u User) ([]string, bool) {

	um.db.AutoMigrate(&u)

	msg, err := um.Validate(u)

	if !err {

		u.Created = time.Now()
		u.Modified = time.Now()
		//バリデーションが通れば作成し、メッセージの中に作成したユーザIDを入れて返す
		um.db.Create(&u)
		msg = append(msg, strconv.Itoa(int(u.Id)))
		return msg, false
	} else {
		//作成できなければエラーメッセージを返す
		return msg, true
	}

}

//指定ユーザidの情報を返す
func (um UserModel) GetAll() ([]User, bool) {

	var users []User
	um.db.Find(&users)
	upm := NewUserProfileModel(um.nc)
	for i, u := range users {
		users[i].Profile, _ = upm.GetByUserId(u.Id)
	}

	//値が取得できたら
	if len(users) > 0 {
		return users, false
	} else {
		return []User{}, true
	}

}

//指定ユーザidの情報を返す
func (um UserModel) GetById(id int) (User, bool) {

	var u User
	um.db.First(&u, id)
	upm := NewUserProfileModel(um.nc)
	u.Profile, _ = upm.GetByUserId(u.Id)
	//値が取得できたら
	if u.Id == id {
		return u, false
	} else {
		return User{}, true
	}

}

//検索メソッド
//任意の条件に一致するユーザを取得する
//TODO: 検索に失敗するということの定義を考える
//→指定条件で検索したところ、その条件にあうユーザは
//いなかった。これはエラーなのか？結果が０なだけで
//検索には成功しているのではないか？
//→この場合における「検索の失敗」とはSQLの構文エラーが起こることであり、
//現状の実装だとそこのエラーハンドリングは呼び出し元が請け負っているので
//Find()でエラーが発生することはありえないと思われる
func (um UserModel) Find(u User) ([]User, bool) {

	var r []User
	um.db.Where(&u).Find(&r)

	//dbに問い合わせて存在していればユーザを返す。なければエラーを返す ←？？
	if len(r) > 0 {
		return r, false
	} else {
		return []User{}, true
	}
}

//更新メソッド
//ユーザの情報を更新する
func (um UserModel) Update(id int, u User) ([]string, bool) {
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
		return []string{}, false

		// //セーブした結果がエラーであれば更新失敗
		// if result := um.db.Save(&tu); result.Error != nil {
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
//ユーザを削除する
func (um UserModel) Delete(id int) ([]string, bool) {

	//idで削除を実行する
	_, err := um.GetById(id)
	if err { //削除するユーザがいなかったらダメ
		return []string{"削除するユーザが存在しません。"}, true
	}
	um.db.Delete(&User{}, id)
	// _, err2 := um.GetById(id)
	// if err2 { //ユーザが取得できなかったら成功
	return []string{"削除に成功しました。"}, false
	// } else {
	// 	return []string{"削除できませんでした。"}, true
	// }

}

//独自バリデーション
//同じメールアドレスがdbに存在しないかどうかを検索する。
func (um UserModel) validateUniqueEmail(email string) bool {
	u := User{
		Email: email,
	}
	_, err := um.Find(u)
	return err
}
