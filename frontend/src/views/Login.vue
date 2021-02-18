<template>
  <div class="login">
    <div class="login-box rounded-lg">
      <a href="/"><img class="login-logo-image" href="/" src="/img/logo_transparent.png"></a>
      <v-tabs fixed-tabs v-model="model">
        <v-tab href="#login">ログイン</v-tab>
        <v-tab href="#signup">新規登録</v-tab>
      </v-tabs>
      <v-tabs-items v-model="model">
        <v-tab-item value="login">
          <h2 class="login-title">ログイン</h2>
          <v-form ref="form" class="login-form" v-on:keyup.enter="login_submit" @submit.prevent>
            <v-text-field
              class="input_email"
              v-model="Email"
              label="メールアドレス"
              :rules="[email_required, email]"
              placeholder="user@example.com"
              outlined
            >
            </v-text-field>
            <v-text-field
              class="input_password"
              v-model="Password"
              label="パスワード"
              :type="show ? 'text' : 'password'"
              :append-icon="show ? 'mdi-eye' : 'mdi-eye-off'"
              :rules="[pass_required]"
              placeholder="Password"
              outlined
              @click:append="show = !show"
            >
            </v-text-field>
            <h4 v-if="Failed" class="login-failed-message">{{ response }}</h4>
            <v-btn color="primary" type="submit" @click="login_submit">ログイン</v-btn>
          </v-form>  

        </v-tab-item>
        <v-tab-item value="signup">
          <h2 class="login-title">新規登録</h2>
          <v-form ref="form" class="signup-form" v-on:keyup.enter="signup_submit" @submit.prevent>
            <v-text-field
              class="input_email"
              v-model="Email"
              label="メールアドレス"
              :rules="[email_required, email]"
              placeholder="user@example.com"
              outlined
            >
            </v-text-field>
            <v-text-field
              class="input_password"
              v-model="Password"
              label="パスワード"
              :type="show ? 'text' : 'password'"
              :append-icon="show ? 'mdi-eye' : 'mdi-eye-off'"
              :rules="[pass_required]"
              placeholder="Password"
              outlined
              @click:append="show = !show"
            >
            </v-text-field>
            <h4 v-if="Conflict" class="login-failed-message">{{ response }}</h4>
            <v-btn color="primary" type="submit" @click="signup_submit">新規登録</v-btn>
          </v-form>  
        </v-tab-item>
      </v-tabs-items>
      
    </div>
  </div>
</template>

<script>
import axios from 'axios'
export default {
    name:"Login",
    data: function () {
      return {
        model: 'login',
        Failed: false,
        url: process.env.VUE_APP_API_URL,
        Conflict: false,
        userId: null,
        show: false,
        error: false,
        response: '',
        message: '',
        Email: '',
        Password: '',
        email_required: Email => !!Email || "メールアドレスを入力してください",
        pass_required: Password => !!Password || "パスワードを入力してください",
        email: Email => (!!Email && /.+@.+\..+/.test(Email)) || "正しいメールアドレスを入力してください"
      }
    },
    methods: {

      //ログインフォーム
      login_submit (){
        //バリデートに成功したらログイン処理を実行する
        if(this.$refs.form.validate()){
          const params = new URLSearchParams();
          params.append('Email', this.Email);
          params.append('Password', this.Password);
          //クッキーを同時送信する
          //ログイン処理してユーザが存在すればログイン。いなければエラー
          axios.post('https://' + this.url  + '/api/v1/login',params,{
            withCredentials: true
          }
          ).then(response => {
              //ログイン成功

              this.userId = response.data.id

              //ログインしたのち、トークンを使ってスペースを検索する。
              //存在しなければ作成ページ、存在すればそのスペースに飛ばす
              axios.get('https://' + this.url + '/api/v1/users/' + this.userId + '/space',{
                headers: {
                  "Authorization" : "Bearer " + response.data.token
                },
                withCredentials: true
              })
              .then(response => {
                //ステータスコードでスペースの有無を判定し存在すればそのページに、
                //存在しなければ作成ページへ遷移
                switch (response.status){
                  case 200: //存在しているのでそのページに遷移
                    this.$router.push( { path: `/space/${this.userId}` }).catch(()=>{});
                }
              })
              .catch(err => {
                //40x系はcatch()に入るのでこちらで処理
                console.log("aaa")
                switch (err.response.status){
                  case 404: //未作成
                    this.$router.push( { path: "new" }).catch((err)=>{ console.log(err)});
                    break;
                }
              })
            }
          ).catch(err => {
            if(err.response) {
              //ログイン失敗

              this.Failed = true
              //レスポンスコードで処理を分ける
              switch (err.response.status){
                //認証失敗
                case 401:
                  this.response = err.response.data
                  break
                //サーバが応答してない
                case 404:
                  this.response = "サーバーが一時的にダウンしています。管理者に連絡してください。"
              }
            }
          });
        }
      },

      //新規登録フォーム
      signup_submit (){
        //バリデートに成功したら
        if(this.$refs.form.validate()){
          let params = new URLSearchParams()
          params.append('Email', this.Email)
          params.append('Password', this.Password)
          //ユーザを作成する。エラーがでたら失敗
          axios.post('https://' + this.url + '/api/v1/signup',params,{
            headers: {
              "Authorization" : "Bearer " + this.$cookies.get("token")
            },
            withCredentials: true
          }).then(() => {
              //登録成功
              //新規登録したユーザはまずスペースを作成する
              window.location.href = '/space/new';
            }
          ).catch(err => {
            if(err.response) {
              switch (err.response.status){
                case 404:
                  this.Conflict = true
                  this.response = "サーバーが一時的にダウンしています。管理者に連絡してください。"
                  break
                case 409:
                  this.Conflict = true
                  this.response = err.response.data[0]
              }
            }
          });
        }
      },
    }
}
</script>

<style>
.login{
  height: 100%;
  background-color: #f3f3f3;
}
.login-box{
  text-align: center;
  margin: 6% 35% 0 35%;
  height: 60%;
  background-color: white;
}
.login-logo-image{
  width: 150px;
  margin-top: 3%;
}
.login-title{
  margin-top: 50px;
}
.login-failed-message{
  color: red;
}
.login-form{
  text-align: center;
  margin-top: 10%;
  margin-left: 10%;
  margin-right: 10%;
}
.signup-form{
  text-align: center;
  margin-top: 10%;
  margin-left: 10%;
  margin-right: 10%;
}
.input_email input{
  ime-mode:disabled;
}
.input_password{
  margin-top: 5%;
}
.input_password input{
  ime-mode:disabled;
  margin-top: 5%;
}
.divider{
  width: 90%;
  height: 0px;
  margin-left: 5%;
  border-width: 1px 0 0 0;
	border-style: solid;
	border-color: #f0f0f0;
}
</style>