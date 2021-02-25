package model

// import (

// 	//標準ライブラリ

// 	"strconv"
// 	"time"

// 	//自作ライブラリ
// 	"main/config"
// 	"main/config/database"

// 	//githubライブラリ
// 	"github.com/go-playground/validator"
// )

// //UserTokenEntity　Entityを埋め込まれている
// type UserToken struct {
// 	Entity
// 	UserId int
// 	Token  string
// 	Expire time.Time
// }

// //呼び出し用ユーザトークンモデル
// //AppModelを埋め込み
// type UserTokenModel struct {
// 	AppModel
// }

// func NewUserTokenModel(t string) *UserTokenModel {
// 	var utm UserTokenModel
// 	utm.db = database.GetInstance(t)
// 	utm.nc = t
// 	return &utm
// }

// func (UserTokenModel) TableName() string {
// 	return "users"
// }

// //バリデーションをかける
// func (utm UserTokenModel) Validate(ut UserToken) ([]string, bool) {

// 	validate := validator.New()
// 	err := validate.Struct(ut)
// 	var messages []string
// 	if err != nil {
// 		for _, err := range err.(validator.ValidationErrors) {
// 			fieldName := err.Field()
// 			switch fieldName {
// 			//ユーザトークンは基本的にnull許容
// 			}
// 		}
// 	}

// 	if len(messages) > 0 {
// 		return messages, true
// 	} else {
// 		return []string{}, false
// 	}

// }

// //ユーザトークンを作成する
// func (utm UserTokenModel) Create(ut UserToken) ([]string, bool) {

// 	utm.db.AutoMigrate(&ut)

// 	msg, err := utm.Validate(ut)

// 	if !err {
// 		//有効期限をconfigから持ってきて生成する
// 		expire := config.Get("expireToken")
// 		expireDay, _ := strconv.Atoi(expire)
// 		ut.Expire = time.Now().AddDate(0, 0, expireDay)
// 		ut.Created = time.Now()
// 		ut.Modified = time.Now()
// 		//バリデーションが通れば作成し、メッセージの中に作成したユーザトークンIDを入れて返す
// 		utm.db.Create(&ut)
// 		msg = append(msg, strconv.Itoa(int(ut.Id)))
// 		return msg, false
// 	} else {
// 		//作成できなければエラーメッセージを返す
// 		return msg, true
// 	}

// }

// //指定ユーザトークンidの情報を返す
// func (utm UserTokenModel) GetAll() ([]UserToken, bool) {

// 	var userTokens []UserToken
// 	utm.db.Find(&userTokens)

// 	//値が取得できたら
// 	if len(userTokens) > 0 {
// 		return userTokens, false
// 	} else {
// 		return []UserToken{}, true
// 	}

// }

// //指定ユーザトークンidの情報を返す
// func (utm UserTokenModel) GetById(id int) (UserToken, bool) {

// 	var ut UserToken
// 	utm.db.First(&ut, id)

// 	//値が取得できたら
// 	if ut.Id == id {
// 		return ut, false
// 	} else {
// 		return UserToken{}, true
// 	}

// }

// //検索メソッド
// //ユーザトークンの任意の条件に一致するユーザトークンを取得する
// //TODO: 検索に失敗するということの定義を考える
// //→指定条件で検索したところ、その条件にあうユーザトークンは
// //いなかった。これはエラーなのか？結果が０なだけで
// //検索には成功しているのではないか？
// //→この場合における「検索の失敗」とはSQLの構文エラーが起こることであり、
// //現状の実装だとそこのエラーハンドリングは呼び出し元が請け負っているので
// //Find()でエラーが発生することはありえないと思われる
// func (utm UserTokenModel) Find(ut UserToken) ([]UserToken, bool) {

// 	var r []UserToken
// 	utm.db.Where(&ut).Find(&r)

// 	//dbに問い合わせて存在していればユーザトークンを返す。なければエラーを返す ←？？
// 	if len(r) > 0 {
// 		return r, false
// 	} else {
// 		return []UserToken{}, true
// 	}
// }

// //更新メソッド
// //ユーザトークンの情報を更新する
// func (utm UserTokenModel) Update(id int) ([]string, bool) {
// 	var tu UserToken

// 	utm.db.AutoMigrate(&tu)
// 	utm.db.First(&tu, id)

// 	//ユーザトークンテーブルでは、基本的に情報の更新は行わない。

// 	//更新日を現在にする
// 	tu.Modified = time.Now()

// 	//バリデーションをかける
// 	msg, err := utm.Validate(tu)

// 	//バリデーションが成功していたら
// 	if !err {
// 		//セーブした結果がエラーであれば更新失敗
// 		if result := utm.db.Save(&tu); result.Error != nil {
// 			return []string{"データベースに保存することができませんでした。"}, true
// 		} else {
// 			return []string{}, false
// 		}
// 	} else {
// 		//バリデーションが失敗していたらそのエラーメッセージを返す
// 		return msg, true
// 	}

// }

// //削除メソッド
// //ユーザトークンを削除する
// func (utm UserTokenModel) Delete(id int) ([]string, bool) {

// 	//idで削除を実行する
// 	_, err := utm.GetById(id)
// 	if err { //削除するユーザトークンがいなかったらダメ
// 		return []string{"削除するユーザトークンが存在しません。"}, true
// 	}
// 	utm.db.Delete(&UserToken{}, id)
// 	_, err2 := utm.GetById(id)
// 	if err2 { //ユーザトークンが取得できなかったら成功
// 		return []string{"削除に成功しました。"}, false
// 	} else {
// 		return []string{"削除できませんでした。"}, true
// 	}

// }
