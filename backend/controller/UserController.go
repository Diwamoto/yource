/*
Package controller is handle some requests.
*/
package controller

import (

	//標準ライブラリ
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	//自作ライブラリ
	"main/model"

	//githubライブラリ
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

//TODO: 変数の具体名化

//ユーザー作成アクション
//POSTされた要素でデータを作成する
func CreateUserAction(c *gin.Context) {

	um := model.NewUserModel("default")

	password, _ := bcrypt.GenerateFromPassword([]byte(c.PostForm("Password")), bcrypt.DefaultCost)

	u := model.User{
		Email:    c.PostForm("Email"),
		Password: string(password),
		Name:     c.PostForm("Name"),
		Phone:    c.PostForm("Phone"),
		Status:   2, //メールアドレス認証ができるまでステータスは有効にならない
		Profile:  model.UserProfile{},
	}
	u.Created = time.Now()
	u.Modified = time.Now()

	user, err := um.Create(u)
	//エラーじゃなければuserの情報を返す
	if err == nil {

		//ユーザのメールアドレスからurlセーフな死活監視トークンを生成する。
		encode_str := base64.StdEncoding.EncodeToString([]byte(user.Email))
		r := strings.NewReplacer("=", "-", "/", "_", "+", ".")
		token := r.Replace(encode_str)
		//トークンモデルに一時保存
		utm := model.NewUserTokenModel("default")
		ut := model.UserToken{
			UserId: user.Id,
			Token:  token,
			Expire: time.Now().Add(24 * time.Hour), //有効期限は一日
		}
		utm.Create(ut)

		mailData := map[string]string{
			"Token": token,
		}

		//メールを送信する
		_, err := SendMail(user.Email, "[yource]仮登録完了のお知らせ", "verify_email", mailData, nil)
		if err != nil {
			//作成できなければエラーメッセージを返す。
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			//送信できれば作成できた旨を返却する。
			c.JSON(http.StatusCreated, user)
		}
	} else {
		//作成できなければエラーメッセージを返す。
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})

	}
}

//検索用アクション
//パラメータを解析して、検索用のオブジェクトに挿入してモデルにて検索する
func SearchUserAction(c *gin.Context) {

	um := model.NewUserModel("default")

	//クエリより検索文字列を取得して、構造体に入れる
	status, _ := strconv.Atoi(c.Query("status"))
	u := model.User{
		Email:    c.Query("email"),
		Name:     c.Query("name"),
		Nickname: c.Query("nickname"),
		Phone:    c.Query("phone"),
		Status:   status,
		//とりあえずプロフィールからユーザを検索するのは非対応
	}
	users, err := um.Find(u)
	//検索した結果が0件でもエラーにはならない。
	//検索した条件が間違えていればエラーに入る
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if len(users) > 0 {
			c.JSON(http.StatusOK, users)
		} else {
			c.JSON(http.StatusNotFound, gin.H{})
		}

	}

}

//ユーザの情報を返すアクション
//GETでパラメータのユーザの情報を取得する
func GetUserByIdAction(c *gin.Context) {

	um := model.NewUserModel("default")

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := um.GetById(id)
	if err == nil {
		if len(user) == 0 {
			c.JSON(http.StatusNotFound, []model.User{})
		} else {
			c.JSON(http.StatusOK, user)
		}
	} else { //エラーが発生した場合はそのエラーを返す
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

//全てのユーザの情報を返すアクション
//GETでパラメータのユーザの情報を取得する
func GetAllUserAction(c *gin.Context) {

	um := model.NewUserModel("default")

	users, err := um.Find(model.User{})
	if err != nil {
		c.JSON(http.StatusOK, users)
	} else {
		c.JSON(http.StatusNotFound, gin.H{})
	}
}

//ユーザの情報を更新するアクション
//PUTでフォームの情報からユーザの情報を更新する
func UpdateUserAction(c *gin.Context) {

	um := model.NewUserModel("default")

	userId, _ := strconv.Atoi(c.Param("id"))
	//ユーザを取得し、取得できたら更新をかける
	_, err := um.GetById(userId)
	if err == nil {
		//フォームから更新内容を取得したユーザ構造体を作成
		var u model.User
		u.Email = c.PostForm("Email")
		u.Password = c.PostForm("Password")
		u.Name = c.PostForm("Name")
		u.Phone = c.PostForm("Phone")
		Status, _ := strconv.Atoi(c.PostForm("Status"))
		u.Status = Status
		user, err2 := um.Update(userId, u)
		if err2 == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusConflict, err2.Error())
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{})
	}
}

//ユーザの削除アクション
func DeleteUserAction(c *gin.Context) {

	um := model.NewUserModel("default")
	userId, _ := strconv.Atoi(c.Param("id"))
	err := um.Delete(userId)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{})
	} else {
		c.JSON(http.StatusConflict, err.Error())
	}
}

//ログインアクション
func LoginAction(c *gin.Context) {

	godotenv.Load(os.Getenv("ENV_PATH"))

	um := model.NewUserModel("default")
	var user model.User
	user.Email = c.PostForm("Email")
	user.Status = 1 //有効なユーザ
	users, err := um.Find(user)

	//ユーザを取得でき、且ハッシュ化されたパスワードが等しければログイン成功
	if err == nil && len(users) == 1 {
		user := users[0]
		err1 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.PostForm("Password")))
		if err1 == nil {

			//ログインに成功したらjwtを発行しクッキーとredisに保存

			//jwtの作成
			//headerのセット
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"token": hash(user.Name),
				"iat":   time.Now(),
				"exp":   time.Now().Add(time.Hour * 24).Unix(),
			})

			//電子署名
			tokenString, _ := token.SignedString([]byte(os.Getenv("SIGN_KEY")))

			//セッションに保存
			session := sessions.Default(c)
			session.Set(tokenString, user.Id)
			session.Save()

			//クッキーに保存する処理はレスポンスではなくvue-cookieで明示的に実行する
			// c.SetCookie("token", tokenString, 3600, "/", "localhost", true, true)
			// c.SetCookie("userId", fmt.Sprint(user.Id), 3600, "/", "localhost", true, true)

			c.JSON(http.StatusOK, gin.H{"token": tokenString})
		} else {
			c.JSON(http.StatusUnauthorized, "メールアドレスもしくはパスワードが間違っています。")
		}
	} else {
		c.JSON(http.StatusUnauthorized, "メールアドレスもしくはパスワードが間違っています。")
	}

}

//ヘッダのjwtからユーザ情報を取得するアクション
func RetriveUserByJWTAction(c *gin.Context) {

	tokenString := c.Request.Header.Get("Authorization")
	//"Bearer"が含まれていたら削除する。仮に含まれていなかったら通さない
	if strings.Contains(tokenString, "Bearer") {
		tokenString = strings.TrimLeft(strings.Replace(tokenString, "Bearer", "", -1), " ")
	} else {
		c.JSON(http.StatusUnauthorized, "不正なログインを検知しました。")
	}

	session := sessions.Default(c)
	userId := session.Get(tokenString)

	//MEMO: jwtからparseする処理がうまく動かないので、いったんセッションはそのままトークンから呼び出す。
	// claims := jwt.MapClaims{}
	// jt, err := jwt.ParseWithClaims(tokenstring, claims, func(token *jwt.Token) (interface{}, error) {
	// 	return []byte(os.Getenv("SIGNKEY")), nil
	// })
	if userId != nil {

		c.JSON(http.StatusOK, gin.H{"userId": userId})

	} else {
		c.JSON(http.StatusUnauthorized, "ログインしていません")
	}

}

//ユーザ有効化アクション
//新規登録を行ったのち、メールアドレスの死活監視を行ったユーザを有効化しログインできるようにする
func VerifyUserAction(c *gin.Context) {

	//フォームのトークンを読む
	token := c.PostForm("Token")

	utm := model.NewUserTokenModel("default")

	user, err := utm.IsValid(token)

	//問題なければユーザを有効化する。
	//トークンも合わせて削除しStatusOKを返す
	if err == nil {
		um := model.NewUserModel("default")

		user.Status = 1 //有効(ログインできる)
		_, err := um.Update(user.Id, user)
		if err == nil {
			err := utm.Delete(token)
			if err == nil {
				c.JSON(http.StatusOK, "")
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func hash(s string) string {
	h := sha256.Sum256([]byte(s))
	hash := hex.EncodeToString(h[:])

	return hash
}
