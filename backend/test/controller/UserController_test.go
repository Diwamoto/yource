package controller

import (
	"encoding/json"
	"main/controller"
	"main/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestCreateUserAction(t *testing.T) {

	tests := []struct {
		user model.User
		want preferResponse
	}{
		{
			//①: テーブルにレコードが何もない状態で作成するユーザ（コンフリクトは起きないはず）
			model.User{
				Email:    "test@example.com",
				Password: "test password",
				Name:     "test name",
			},
			preferResponse{
				code: http.StatusCreated,
				body: map[string]interface{}{
					"Email": "test@example.com",
					//"Password": "test password", //パスワードも調査した方が良いが、データベース内ではhash化されるため
					"Name": "test name",
				},
			}, //ユーザは作成できるはず
		},
		{
			//②: メールアドレスがかぶっているユーザ
			model.User{
				Email:    "test@example.com",
				Password: "test password",
				Name:     "test name",
			},
			preferResponse{
				code: http.StatusConflict,
				body: map[string]interface{}{
					"message": "[\"入力されたメールアドレスは既に登録されています。\"]",
				},
			}, //既に作成されているのでコンフリクトが起きるはず
		},
	}
	for i, tt := range tests {

		//テスト準備
		//リクエストを作成
		requestBody := strings.NewReader("Email=" + tt.user.Email + "&Name=" + tt.user.Name + "&Password=" + tt.user.Password)
		//レスポンス
		//ここに帰ってくる
		response := httptest.NewRecorder()
		//コンテキストを作成
		c, _ := gin.CreateTestContext(response)
		//リクエストを格納
		c.Request, _ = http.NewRequest(
			http.MethodPost,
			"/api/v1/users",
			requestBody,
		)
		//フォーム属性を付与
		c.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		// テストのコンテキストを持って実行
		controller.CreateUserAction(c)

		var responseBody map[string]interface{}
		_ = json.Unmarshal(response.Body.Bytes(), &responseBody)

		//ステータスコードがおかしいもしくは帰ってきたメッセージが想定と違えばダメ
		if response.Code != tt.want.code {
			t.Errorf("%d番目のテストが失敗しました。想定返却コード：%d, 実際の返却コード：%d", i+1, tt.want.code, response.Code)
		} else {
			//実際に帰ってきたレスポンスの中に想定された値が入っているかどうか
			for key := range tt.want.body {
				//値の存在チェック
				if _, exist := responseBody[key]; exist {

					//値の内容チェック
					if responseBody[key] != tt.want.body[key] {
						t.Errorf("%d番目のテストが失敗しました。想定されたキー「%s」の値:%s, 実際に返却された値:%s", i+1, key, tt.want.body[key], responseBody[key])
					} // else{
					//クリアはここだけ
					// }

				} else {
					t.Errorf("%d番目のテストが失敗しました。想定された「%s」がレスポンスボディに入っていません。", i+1, key)
				}
			}
		}
	}
}
