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

//controller.CreateUserActionのテスト
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
					"error": "入力されたメールアドレスは既に登録されています。",
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
			"/v1/users",
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

//controller.CreateUserActionのテスト
func TestGetUserByIdAction(t *testing.T) {

	tests := []struct {
		userId string
		want   preferResponse
	}{
		// { //TODO: テスト成功させる
		// 	//①: Id=1で検索
		// 	"1",
		// 	preferResponse{
		// 		code: http.StatusOK,
		// 		body: map[string]interface{}{
		// 			"Id": 1,
		// 		},
		// 	}, //Id=1のユーザが帰ってくるはず
		// },
		{
			//②: Idに適当な文字列を入れて検索
			"jkfdji",
			preferResponse{
				code: http.StatusBadRequest,
				body: map[string]interface{}{"error": "strconv.Atoi: parsing \"jkfdji\": invalid syntax"},
			}, //エラーが帰ってくるはず
		},
	}
	for i, tt := range tests {

		//テスト準備
		//レスポンス
		//ここに帰ってくる
		response := httptest.NewRecorder()
		//コンテキストを作成
		_, r := gin.CreateTestContext(response)
		//リクエストを格納
		request := httptest.NewRequest(
			http.MethodGet,
			"/"+tt.userId,
			nil,
		)
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("Accept", "application/json")

		r.GET("/:id", controller.GetUserByIdAction)

		// テストのコンテキストを持って実行
		r.ServeHTTP(response, request)

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
