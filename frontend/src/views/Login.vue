<template>
  <div class="login">
    <div class="login-box rounded-lg">
      <h1>
        <a href="/"
          ><img
            class="login-logo-image"
            href="/"
            src="/img/logo_transparent.png"
        /></a>
      </h1>
      <div class="standard-login-box-wrapper">
        <h2 class="login-title">{{ model_jp }}</h2>
        <v-tabs fixed-tabs v-model="model">
          <v-tab href="#login">ログイン</v-tab>
          <v-tab href="#signup">新規登録</v-tab>
        </v-tabs>
        <v-tabs-items v-model="model">
          <v-tab-item value="login">
            <v-form
              ref="form"
              class="login-form"
              v-on:keyup.enter="login_submit"
              @submit.prevent
            >
              <v-text-field
                class="input_email"
                v-model="Email"
                label="メールアドレス"
                :rules="[email_required, email]"
                placeholder="user@example.com"
                outlined
                @keydown.enter="login_submit"
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
                @keydown.enter="login_submit"
              >
              </v-text-field>
              <v-alert v-if="Failed" type="error" dense>{{ response }}</v-alert>
              <v-btn
                color="primary"
                type="submit"
                ref="submit"
                :disabled="loading"
                large
                class="submit-button"
                @click="login_submit"
                >ログイン</v-btn
              >
            </v-form>
          </v-tab-item>
          <v-tab-item value="signup">
            <v-form
              ref="form"
              class="signup-form"
              v-on:keyup.enter="signup_submit"
              @submit.prevent
            >
              <v-text-field
                class="input_email"
                v-model="Email"
                label="メールアドレス"
                :rules="[email_required, email]"
                placeholder="user@example.com"
                outlined
                @keydown.enter="signup_submit"
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
                @keydown.enter="signup_submit"
              >
              </v-text-field>
              <v-alert v-if="Conflict" type="error" dense>{{
                response
              }}</v-alert>
              <v-btn
                color="primary"
                type="submit"
                ref="submit"
                :disabled="loading"
                large
                class="submit-button"
                @click="signup_submit"
                >新規登録</v-btn
              >
            </v-form>
            <v-progress-circular
                v-show="loading"
                indeterminate
                color="primary"
                class="loader mt-6"
              ></v-progress-circular>
          </v-tab-item>
        </v-tabs-items>
      </div>
      <!--
      <div class="sns-login-wrapper">
        <h2 class="login-title">SNSでログインする</h2>
        LINEログインのみアイコンを用意するため分ける
        <div class="sns-login">
          <v-btn
            depressed
            dark
            x-large
            :style="{ backgroundColor: '#00B900' }"
            class="sns-login-button"
          >
            <v-img src="/img/line_88.png" max-width="40" left></v-img>
            LINEでログイン
          </v-btn>
        </div>
         他はマテリアルアイコンを使用する 
        <div class="sns-login" v-for="item in sns" :key="item.name">
          <v-btn
            depressed
            dark
            x-large
            :style="{ backgroundColor: item.color }"
            class="sns-login-button"
          >
            <v-icon left>
              {{ item.icon }}
            </v-icon>
            {{ item.name }}でログイン
          </v-btn>
        </div>
      </div>
      -->
    </div>
  </div>
</template>

<script>
export default {
  name: "Login",
  data: function () {
    return {
      model: "login",
      model_jp: "ログイン",
      Failed: false,
      Conflict: false,
      userId: null,
      show: false,
      error: false,
      response: "",
      message: "",
      Email: "",
      loading: false,
      Password: "",
      line_url:
        "https://access.line.me/oauth2/v2.1/authorize?response_type=code&client_id=1655708978&redirect_uri=https%3A%2F%2F0c583dca4f4a.ngrok.io%2Flogin&state=" +
        Math.random().toString(36).substring(2) +
        "&scope=profile%20openid%20email",
      email_required: (Email) => !!Email || "メールアドレスを入力してください",
      pass_required: (Password) => !!Password || "パスワードを入力してください",
      email: (Email) =>
        (!!Email && /.+@.+\..+/.test(Email)) ||
        "正しいメールアドレスを入力してください",
      sns: [
        {
          name: "GitHub",
          icon: "mdi-github",
          color: "#171515",
        },
        {
          name: "Twitter",
          icon: "mdi-twitter",
          color: "#00acee",
        },
        {
          name: "Facebook",
          icon: "mdi-facebook",
          color: "#3B5998",
        },
        {
          name: "Google+",
          icon: "mdi-google-plus",
          color: "#DB4437",
        },
      ],
    };
  },
  created() {
    //LINEログイン用
    //パラメータを持ってアクセスしてきた場合はlineAPIにトークンを貰いそれでログインする
    if (this.$route.query.code != null && this.$route.query.state != null) {
      const params = new URLSearchParams();
      params.append("grant_type", "authorization_code");
      params.append("code", this.$route.query.code);
      params.append("redirect_uri", "https://0c583dca4f4a.ngrok.io/login"); //TODO:プレースホルダーに直す
      params.append("client_id", process.env.VUE_APP_LINE_CLIENT_ID);
      params.append("client_secret", process.env.VUE_APP_LINE_CLIENT_SECRET);
      this.$http
        .post("https://api.line.me/oauth2/v2.1/token", params)
        .then((response) => {
          var lineJwt = response.data.id_token;
          var tokens = lineJwt.sprit(",");
          console.log(tokens);
        })
        .catch(() => {
          this.$cookies.set(
            "msg",
            "LINEでログインすることができませんでした。",
            { expires: "1H" }
          );
        });
    }

    //ログイン画面に遷移時にトークンが残っており、トークンがまだ有効であれば自動でログインさせる。
    //有効でなければ再ログインが必要な旨を表示する
    if (this.$cookies.get("token") != null) {
      this.$http
        .get(this.$api + "/v1/retrive", {
          headers: {
            Authorization: "Bearer " + this.$cookies.get("token"),
          },
          withCredentials: true,
        })
        .then((response) => {
          switch (response.status) {
            case 200: //ユーザidが取得できたらログインさせる。
              this.$router.push({ path: "app" }).catch(() => {});
          }
        })
        .catch(() => {
          this.$cookies.remove("token");
          this.$cookies.set("msg", "続けるにはログインが必要です。", {
            expires: "1H",
          });
        });
    }
  },
  mounted() {
    this.flash_msg()
    this.$setTitle("ログイン | yource", "yourceはポートフォリオを簡単に公開できるサービスです。")
  },
  updated() {
    this.flash_msg()
  },
  beforeUpdate() {
    switch (this.model) {
      case "login":
        this.model_jp = "ログイン";
        break;
      case "signup":
        this.model_jp = "新規登録";
        break;
    }
  },
  methods: {
    //ログインフォーム
    login_submit() {
      this.loading = true
      //バリデートに成功したらログイン処理を実行する
      if (this.$refs.form.validate()) {
        const params = new URLSearchParams();
        params.append("Email", this.Email);
        params.append("Password", this.Password);
        //クッキーを同時送信する
        //ログイン処理してユーザが存在すればログイン。いなければエラー
        this.$http
          .post(this.$api + "/v1/login", params, {
            withCredentials: true,
          })
          .then((response) => {
            //ログイン成功

            //jwtをサーバに保存
            this.$cookies.set("token", response.data.token, { expires: "1D" });

            //jwtを使ってuseridを取得する
            //ユーザIDを取得してくる
            this.$http
              .get(this.$api + "/v1/retrive", {
                headers: {
                  Authorization: "Bearer " + this.$cookies.get("token"),
                },
                withCredentials: true,
              })
              .then((response) => {
                switch (response.status) {
                  case 200: //ユーザidを保存する
                    this.userId = response.data.userId;

                    //ログインしたのち、トークンを使ってスペースを検索する。
                    //存在しなければ作成ページ、存在すればそのスペースに飛ばす
                    this.$http
                      .get(
                        this.$api + "/v1/users/" + this.userId + "/space",
                        {
                          headers: {
                            Authorization:
                              "Bearer " + this.$cookies.get("token"),
                          },
                          withCredentials: true,
                        }
                      )
                      .then((response) => {
                        //ステータスコードでスペースの有無を判定し存在すればそのページに、
                        //存在しなければ作成ページへ遷移
                        switch (response.status) {
                          case 200: //存在しているのでそのページに遷移
                            this.$router
                              .push({ path: "/app" })
                              .catch(() => {});
                        }
                      })
                      .catch((err) => {
                        switch (err.response.status) {
                          case 404: //未作成
                            this.$router.push({ path: "/new" }).catch(() => {});
                            break;
                        }
                      });
                }
              });
          })
          .catch((err) => {
            if (err.response) {
              //ログイン失敗

              this.Failed = true;
              //レスポンスコードで処理を分ける
              switch (err.response.status) {
                //認証失敗
                case 401:
                  this.response = err.response.data;
                  break;
                //サーバが応答してない
                case 404:
                  this.response =
                    "サーバーが一時的にダウンしています。管理者に連絡してください。";
              }
            }
          });
      }
      this.loading = false
    },

    //新規登録フォーム
    signup_submit() {
      this.loading = true
      //バリデートに成功したら
      if (this.$refs.form.validate()) {
        let params = new URLSearchParams();
        params.append("Email", this.Email);
        params.append("Password", this.Password);
        //ユーザを作成する。エラーがでたら失敗
        this.$http
          .post(this.$api + "/v1/signup", params, {
            withCredentials: true,
          })
          .then(() => {
            //登録成功
            //仮登録メールが送られている旨を通知する。
            this.$cookies.set(
              "ss",
              Math.random().toString(32).substring(2),
            );
            this.$router.push({ path: `/register` });
          })
          .catch((err) => {
            if (err.response) {
              switch (err.response.status) {
                case 404:
                  this.Conflict = true;
                  this.response =
                    "サーバーが一時的にダウンしています。管理者に連絡してください。";
                  break;
                case 409:
                  this.Conflict = true;
                  this.response = err.response.data.error
                    .replace("[", "")
                    .replace("]", "")
                    .replace('"', "")
                    .replace('"', "");
              }
            }
          });
      }
      this.loading = false
    },

    flash_msg() {
      //フラッシュメッセージ
      if (this.$cookies.get("msg") != null) {
        this.Failed = true;

        this.response = this.$cookies.get("msg");
        this.$cookies.remove("msg");
      }
    },
  },
};
</script>

<style>
.login {
  height: 100%;
  background-color: #f3f3f3;
}
.login-box {
  text-align: center;
  margin: 10vh auto;
  height: 70vh;
  width: 350px;
  background-color: white;
  box-shadow: 0 6px 6px -3px rgba(0, 0, 0, 0.2),
    0 10px 14px 1px rgba(0, 0, 0, 0.14), 0 4px 18px 3px rgba(0, 0, 0, 0.12) !important;
  animation-name: login;
  animation-duration: 0.5s;
  animation-timing-function: ease-out;
  animation-fill-mode: forwards;
}
.standard-login-box-wrapper {
  padding: 0 10px 0 10px;
  height: 80%;
  width: 100%;
  /* border-width: 0 1px 0 0;
  border-style: solid;
  border-color: #d0d0d0; */
}
.login-logo-image {
  width: 30%;
  padding: 2vh 0 2vh 0;
}
.login-title {
  font-weight: 600;
}
.login-failed-message {
  color: red;
}
.login-form {
  text-align: center;
  margin-top: 4vh;
  margin-left: 5%;
  margin-right: 5%;
}
.signup-form {
  text-align: center;
  margin-top: 4vh;
  margin-left: 5%;
  margin-right: 5%;
}
.input_email input {
  ime-mode: disabled;
}
.input_password {
  margin-top: 10vh;
}
.input_password input {
  ime-mode: disabled;
}
.submit-button {
  margin-top: 2vh;
  width: 300px;
}
.sns-login-wrapper {
  height: 100%;
}
.sns-login-button {
  width: 300px;
  text-transform: none !important;
  margin: 10px 40px 10px 20px;
}
.v-btn__content {
  text-decoration: none;
}
.v-icon {
  width: 40px;
}
@keyframes login {
  0% {
    opacity: 0;
    transform: translateY(20px);
  }
  100% {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>