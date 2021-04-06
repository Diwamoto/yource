<template>
  <div class="register">
    <Header></Header>
    <div class="register-box rounded-lg">
      <h1 class="register-title">{{ title }}</h1>
      <div class="register-body">
        {{ body }}
      </div>
      <div class="text-center">
        <v-btn
          color="primary"
          elevation="2"
          class="btn-top text-center mt-10"
          large
          link
          href="/login"
          >ログイン画面に遷移する</v-btn
        >
      </div>
    </div>
    <Footer></Footer>
  </div>
</template>

<script>
import Header from "@/components/top/Header.vue";
import Footer from "@/components/top/Footer.vue";
export default {
  name: "Verify",
  components: {
    Header,
    Footer,
  },
  data() {
    return {
      title: "",
      body: "",
    };
  },
  created() {
    //直アクセスを弾く
    if (this.$route.query.token == null || this.$route.query.token == "") {
      this.$router.push({ name: "404", params: { url: "register" } });
    }

    //認証APIを叩いて確認する
    let params = new URLSearchParams();
    params.append("Token", this.$route.query.token);
    this.$http
      .post(this.$api + "/v1/verify", params, {
        withCredentials: true,
      })
      .then(() => {
        this.title = "本登録完了";
        this.body =
          "メールアドレスが認証できました。ログイン画面よりログインしてください。引き続きyourceをよろしくお願いいたします。";
      })
      .catch((err) => {
        switch (err.response.status) {
          case 400:
            this.$router.push({ name: "404" });
            break;
          case 409:
            this.title = "トークンが不正です。";
            this.body =
              "入力されたトークンの整合性が取れません。該当メールアドレスのユーザとトークンを削除いたしましたので、大変申し訳ありませんが、再度登録をよろしくお願いいたします。";
        }
      });
  },
  methods: {},
};
</script>

<style>
.register {
  height: 100%;
  background-color: #f3f3f3;
}
.register-box {
  background-color: white;
  margin: 15vh 32.5vw 10vh 32.5vw;
  height: 61vh;
  width: 35vw;
  opacity: 0;
  box-shadow: 0 6px 6px -3px rgba(0, 0, 0, 0.2),
    0 10px 14px 1px rgba(0, 0, 0, 0.14), 0 4px 18px 3px rgba(0, 0, 0, 0.12) !important;
  animation-name: register;
  animation-duration: 0.5s;
  animation-timing-function: ease-out;
  animation-fill-mode: forwards;
}
.register-title {
  padding: 2% 0 2% 0;
  text-align: center;
  border-width: 0 0 2px 0;
  border-style: solid;
  border-color: #f0f0f0;
}
.register-body {
  padding: 8vh 10% 0 10%;
  line-height: 170%;
}
.btn-top {
  text-transform: none !important;
}
@keyframes register {
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