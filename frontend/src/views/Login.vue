<template>
  <div class="login">
    <div class="login-box rounded-lg">
      <img class="login-logo-image" src="/img/logo_transparent.png">
      <div class="divider"></div>
      <div class="login-title"><h3>ログイン</h3></div>
      <form class="login-form">
        <v-text-field
          v-model="Email"
          label="メールアドレス"
          :rules="[email_required, email]"
          placeholder="user@example.com"
          outlined
        >
        </v-text-field>
        <v-text-field
          class="password"
          v-model="Password"
          label="パスワード"
          :rules="[pass_required]"
          placeholder="Password"
          outlined
        >
        </v-text-field>
        <h4 v-if="Failed" class="login-failed-message">ユーザIDもしくはパスワードが正しくありません。</h4>
        <v-btn v-on:click="submit">ログイン</v-btn><br>
        <v-btn text href="/signup">新規登録はこちら</v-btn>
      </form>
      
    </div>
  </div>
</template>

<script>
import axios from 'axios'
export default {
    name:"Login",
    data: function () {
      return {
        Failed: false,
        error: false,
        message: '',
        Email: '',
        Password: '',
        email_required: Email => !!Email || "メールアドレスを入力してください",
        pass_required: Password => !!Password || "パスワードを入力してください",
        email: Email => (!!Email && /.+@.+\..+/.test(Email)) || "正しいメールアドレスを入力してください"
      }
    },
    methods: {
      submit (){
        let params = new URLSearchParams()
        params.append('Email', this.Email)
        params.append('Password', this.Password)
        //ログイン処理してユーザが存在すればログイン。いなければエラー
        axios.post('http://yource.localhost/users/login',params
        ).then(() => {
            //ログイン成功
            window.location.href = '/mypage';
          }
        ).catch(err => {
          if(err.response) {
            //ログイン失敗
            this.Failed = true
          }
        });
        
      }
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
  margin: 6% 35% 35%;
  height: 50%;
  background-color: white;
}
.login-logo-image{
  width: 150px;
  margin-top: 3%;
}
.login-title{
  margin-top: 50px;
  margin-bottom: 50px;
}
.login-failed-message{
  color: red;
}
.login-form{
  text-align: center;
  margin-left: 10%;
  margin-right: 10%;
}
.password{
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