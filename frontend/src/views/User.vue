<template>
  <div class="user">
      <transition><div class="flash-wrapper" v-if="flashOn"><v-alert :type="flashType">{{ flashMsg }}</v-alert></div></transition>
      <router-view :user="user" :flash-msg="flash"></router-view>
  </div>
</template>

<script>
export default {
  name:"User",
  data() {
    return {
      userId : 0, 
      user : [],
      flashMsg: "",
      flashType: "",
      flashOn: false,
    }
  },
  mounted(){
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
          case 200: //ユーザ情報を取得
            this.userId = response.data["userId"];
            //ユーザIDからユーザの情報を取得
            this.$http.get(this.$api + "/v1/users/" + this.userId, {
              headers: {
                Authorization: "Bearer " + this.$cookies.get("token"),
              },
              withCredentials: true,
            }).then((response) => {
              this.user = response.data[0]
            }).catch(() => {
              this.toLoginPage()
            })
        }
      })
      .catch(() => {
        this.toLoginPage()
      });
    this.$setTitle("ユーザ情報変更 | yource", "yourceはポートフォリオを簡単に公開できるサービスです。")
  },
  methods: {
    toLoginPage(){
      //ユーザ情報が取得できなかったら再ログインさせる
      this.$cookies.remove("token");
      this.$cookies.set(
        "msg",
        "続けるにはログインが必要です。",
        {expires: "1H",}
      );
      this.$router.push({ path: "/login" });
    },
    flash(type, text){
      this.flashMsg = text
      this.flashType = type
      this.flashOn = true
      setTimeout(() => {
        this.flashOn = false}
        ,2000
      )
    },
  }
}
</script>

<style>
.user{
  width: 100vw;
  background-color: #e2e2e2;
}
.flash-wrapper{
  top: 0;
  left: 0;
  width: 50vw;
  margin: 0 25vw 0 25vw;
  z-index: 100;
  position: fixed;
}
.v-enter,
.v-leave-to {
  opacity: 0;
}

.v-enter-active,
.v-leave-active {
  transition: 1s;
}
</style>