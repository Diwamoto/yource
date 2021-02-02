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
              v-on:change="resetValidate"
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
              v-on:change="resetValidate"
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
              v-on:change="resetValidate"
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
              v-on:change="resetValidate"
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
        Conflict: false,
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
      login_submit (){
        this.$refs.form.validate()
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
            this.response = err.response.data
          }
        });
        
      },
      signup_submit (){
        this.$refs.form.validate()
        let params = new URLSearchParams()
        params.append('Email', this.Email)
        params.append('Password', this.Password)
        //ユーザを作成する。エラーが
        axios.post('http://yource.localhost/users',params
        ).then(() => {
            //登録成功
            window.location.href = '/mypage';
          }
        ).catch(err => {
          if(err.response) {
            switch (err.response.status){
              case 409:
                this.Conflict = true
                this.response = err.response.data[0]
            }
          }
        });
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
  margin: 6% 35% 35%;
  height: 45%;
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