<template>
  <div class="home">
    <Systembar :space-name="space.Name" ref="Systembar"></Systembar>

    <Navbar
      :space="space"
      :channels="channels"
      :user-id="userId"
      :switch-channel="switchChannel"
    ></Navbar>

    <Main
      :channel="channel"
      :posts="posts"
      :user-id="userId"
      ref="Main"
      :get-posts="getPosts"
    ></Main>
  </div>
</template>

<script>
import Systembar from "@/components/home/Systembar.vue";
import Navbar from "@/components/home/Navbar.vue";
import Main from "@/components/home/Main.vue";
export default {
  components: {
    Systembar,
    Navbar,
    Main,
  },
  name: "Home",
  data() {
    return {
      userId: null,
      space: [],
      channels: [],
      channel: [],
      posts: [],
    };
  },
  mounted() {
    //ユーザIDを取得してくる
    this.$http
      .get(this.$api + "/api/v1/retrive", {
        headers: {
          Authorization: "Bearer " + this.$cookies.get("token"),
        },
        withCredentials: true,
      })
      .then((response) => {
        switch (response.status) {
          case 200: //ユーザ情報を取得
            this.userId = response.data["userId"];

            //スペースの名前を取得してくる
            this.$http
              .get(
                  this.$api +
                  "/api/v1/users/" +
                  this.userId +
                  "/space",
                {
                  headers: {
                    Authorization: "Bearer " + this.$cookies.get("token"),
                  },
                  withCredentials: true,
                }
              )
              .then((response) => {
                switch (response.status) {
                  case 200: //名前を挿入
                    this.space = response.data;
                    this.channels = this.space.Channels;
                    this.switchChannel(this.channels[0]);
                }
              })
              .catch(() => {
                //ここに来るのは本来発生し得ないが、何らかのタイミングでスペースが消えてしまった場合は
                //↑再ログイン時にhomeに強制的に飛ばすため発生する
                //作成ページに飛ばす
                //this.$router.push( { path: "new" }).catch((err)=>{ console.log(err)});
              });
        }
      })
      .catch(() => {
        // //ユーザ情報が取得できなかったら再ログインさせる
        // this.$cookies.remove("token");
        // this.$cookies.set(
        //   "msg",
        //   "続けるにはログインが必要です。",
        //   3600,
        //   "/",
        //   "localhost",
        //   true,
        //   "None"
        // );
        // this.$router.push({ path: "/login" }).catch((err) => {
        //   console.log(err);
        // });
      });
  },
  methods: {
    //投稿を取得する
    getPosts() {
      //一度投稿を初期化
      this.posts = [];
      this.$http
        .get(
            this.$api +
            "/api/v1/channels/" +
            this.channel.Id +
            "/posts",
          {
            headers: {
              Authorization: "Bearer " + this.$cookies.get("token"),
            },
            withCredentials: true,
          }
        )
        .then((response) => {
          switch (response.status) {
            case 200:
              this.posts = response.data;
          }
        })
        .catch(() => {});
    },
    //チャンネルを変更する
    switchChannel(c) {
      this.channel = c;
      this.getPosts();
      this.$refs.Main.scrollBottom();
    },
  },
};
</script>

<style>
body::-webkit-scrollbar {
  display: none;
}
</style>