package controller

import (
	"errors"
	"main/config/database"
	"main/controller"
	"main/model"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

//テストで使用する帰り値用のオブジェクト
//codeとレスポンスのボディを見て判断する
type preferResponse struct {
	code int                    //httpステータスコード
	body map[string]interface{} //帰ってくる文字列
}

//テストメイン関数
//コントローラのテストを管理する
func TestMain(m *testing.M) {

	//ginをテストモードに設定
	gin.SetMode(gin.TestMode)

	//まずデータベース接続オブジェクトを作成する
	db := database.GetInstance("test")
	//テストに必要なテーブルを全て作成する
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.UserProfile{})
	db.AutoMigrate(&model.Space{})
	db.AutoMigrate(&model.Channel{})
	db.AutoMigrate(&model.Post{})

	//テスト用のユーザを作成

	// //ユーザ以外のテストに使用するテストユーザを作成
	// mtu := model.User{
	// 	Email:    "master@example.com",
	// 	Password: "4AeNkWVisJ",
	// 	Name:     "master name",
	// 	Phone:    "028-0728-9727",
	// 	Status:   1,
	// 	Profile:  model.UserProfile{},
	// }
	// mtu.Created = time.Now()
	// mtu.Modified = time.Now()
	// db.Create(&mtu)

	// //ユーザ以外のテストに使用するテストユーザを作成その２
	// mtu2 := model.User{
	// 	Email:    "master2@example.com",
	// 	Password: "4AeNkWVisJ",
	// 	Name:     "master name",
	// 	Phone:    "028-0728-9727",
	// 	Status:   1,
	// 	Profile:  model.UserProfile{},
	// }
	// mtu2.Created = time.Now()
	// mtu2.Modified = time.Now()
	// db.Create(&mtu2)

	// //スペース以外のテストに使用するテストスペースを作成
	// mts := model.Space{
	// 	UserId:      1,
	// 	Name:        "master name",
	// 	Description: "master disc",
	// 	SubDomain:   "master",
	// 	Status:      true,
	// 	Publish:     true,
	// }
	// mts.Created = time.Now()
	// mts.Modified = time.Now()
	// db.Create(&mts)

	// //チャンネル以外のテストに使用するテストチャンネルを作成
	// mtc := model.Channel{
	// 	SpaceId:     1,
	// 	Name:        "master name",
	// 	Description: "master disc",
	// }
	// mtc.Created = time.Now()
	// mtc.Modified = time.Now()
	// db.Create(&mtc)

	//テストを実行
	code := m.Run()

	//テスト用のデータベースの全てのテーブルを破棄
	db.DropTable(&model.User{})
	db.DropTable(&model.UserProfile{})
	db.DropTable(&model.Space{})
	db.DropTable(&model.Channel{})
	db.DropTable(&model.Post{})
	db.Close()
	os.Exit(code)
}

//sendMailのテスト
func TestSendMail(t *testing.T) {
	type args struct {
		to       string
		receiver string
		subject  string
		templete string
		data     map[string]string
		option   map[string]string
	}
	tests := []struct {
		args    args
		wantErr error
	}{
		{
			//①: 正しい入力値
			args{
				to:       "test@example.com",
				receiver: "test",
				subject:  "test",
				templete: "verify_email",
			},
			nil, //メールは飛ぶはず
		},
		{
			//②: 存在しないテンプレート
			args{
				to:       "test@example.com",
				receiver: "test",
				subject:  "test",
				templete: "NotFound",
			},
			errors.New("open /go/src/github.com/Diwamoto/yource/view/mail/NotFound.html: no such file or directory"), //エラーになるはず
		},
	}
	for i, tt := range tests {
		_, err := controller.SendMail(tt.args.to, tt.args.subject, tt.args.templete, tt.args.data, tt.args.option)
		if err != tt.wantErr {
			if err.Error() != tt.wantErr.Error() {
				t.Errorf("%d番目のテストが失敗しました。想定結果：%s、実際の結果：%#v", i+1, tt.wantErr.Error(), err.Error())
			}
		}
	}
}
