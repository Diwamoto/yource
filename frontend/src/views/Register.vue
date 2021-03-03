<template>
  <div class="register">
    <Header></Header>
    <div class="register-box rounded-lg">
      <h1 class="register-title">仮登録完了</h1>
      <div class="register-body">
        入力いただいたメールアドレス宛てに、本登録を行うためのURLをメール送付いたしました。
        メールのURLより本登録画面に進んでいただき、アカウントを有効化してください。
        <div class="mt-4 px-5">
          <small
            >※仮登録完了メールが届かない、入力いただいたメールアドレスが間違っている可能性がございます。</small
          >
        </div>
        <div class="px-5">
          <small
            >※仮登録の有効期限は24時間となっておりますので、お早めに本登録をよろしくお願いいたします。</small
          >
        </div>
      </div>
      <v-btn color="primary" elevation="2" class="btn-top" large link href="/"
        >Topページへ戻る</v-btn
      >
    </div>
    <Footer></Footer>
  </div>
</template>

<script>
import Header from "@/components/top/Header.vue";
import Footer from "@/components/top/Footer.vue";
export default {
  name: "Register",
  components: {
    Header,
    Footer,
  },
  created(){
    //直アクセスを弾く
    if (this.$cookies.get("ss") == null){
      this.$router.push({ name: '404' ,params: { url: "register" } })
    } else {
      this.$cookies.remove("ss")
    }
  }, 
  methods: {
  }
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
  margin: 7% 37%;
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