package controller

import (
	"bytes"
	"main/config"
	"os"
	"text/template"

	"github.com/joho/godotenv"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

//メールを送る sendGridを使う
func SendMail(to string, subject string, templete string, data map[string]string, option map[string]string) (*rest.Response, error) {

	godotenv.Load(os.Getenv("ENV_PATH"))
	from := mail.NewEmail("info", config.ToString("sendGridFrom"))

	//送信者名の指定があればつける
	var mailTo *mail.Email
	if receiver, ok := option["receiver"]; ok {
		mailTo = mail.NewEmail(receiver, to)
	} else {
		mailTo = mail.NewEmail("", to)
	}

	//テンプレートを呼び出す
	body, err := parseTemplete(templete, data)
	//もし読み出せなかったらエラーを返す
	if err != nil {
		return nil, err
	}
	message := mail.NewSingleEmail(from, subject, mailTo, "", body)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	return client.Send(message)

}

//テンプレートをファイルから読み込み、変数を埋め込む。
func parseTemplete(fileName string, data map[string]string) (string, error) {

	//ファイルを開く
	rootPath := config.ToString("rootPath")
	t, err := template.New(fileName + ".html").ParseFiles(rootPath + "/view/mail/" + fileName + ".html")
	if err != nil {
		return "", err
	}
	//バッファに一度書き出して文字列にして出力する。
	var result bytes.Buffer
	if err = t.Execute(&result, data); err != nil {
		return "", err
	} else {
		return result.String(), nil
	}

}
