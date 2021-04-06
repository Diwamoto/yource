<template>
  <div class="input-form">
    <v-form @submit.prevent>
      <v-text-field
        ref="newPostTextField"
        :label="label"
        v-model="newPost"
        outlined
        class="input-box"
        append-outer-icon="mdi-send"
        clear-icon="mdi-close"
        clearable
        type="text"
        :hint="hint"
        v-on:keydown.meta.enter="submit"
        @click:append-outer="submit"
        @click:clear="clearMessage"
      ></v-text-field>
    </v-form>
  </div>
</template>

<script>
export default {
  Name: "Form",
  props: {
    channel: [Object, Array],
    userId: Number,
    getPosts: {
      type: Function,
      required: true,
    },
  },
  data() {
    return {
      socket: new WebSocket("wss://" + process.env.VUE_APP_WEBSOCKET_URL),
      label: "",
      newPost: "",
      hint: "",
    };
  },
  created() {
    //ラベルを読み込む
    this.updateLabel();

    //websocket経由でメッセージの追加の通知をもらった際に投稿を更新する
    this.socket.onmessage = () => {
        this.getPosts();
    };
  },
  mounted() {
    //システムバーにwifiアイコンを表示する
    this.socket.onopen = () =>{
      this.$parent.$parent.$parent.$refs.Systembar.loadComplete()
    }
  },
  updated() {
    //フォームに文字が入力されていれば送信ヒントメッセージを追加
    if (this.newPost != "") {
      this.hint = "cmd + enterでも送信することができます。";
    } else {
      this.hint = "";
    }

    //ラベルを読み込む
    this.updateLabel();
  },
  methods: {
    //投稿を送信する
    submit() {
      if (this.newPost.length > 0) {
        const params = new URLSearchParams();
        params.append("Content", this.newPost);
        params.append("UserId", this.userId);
        this.$http
          .post(
              this.$api +
              "v1/channels/" +
              this.channel.Id +
              "/posts",
            params,
            {
              headers: {
                Authorization: "Bearer " + this.$cookies.get("token"),
              },
              withCredentials: true,
            }
          )
          .then(() => {
            //投稿が作成できたら投稿一覧をリセットし、入力欄を削除する。
            this.getPosts();
            
            //同時にws経由でメッセージの追加を送信
            this.socket.send(JSON.stringify(
                {
                    message: this.newPost
                }
            ));
            this.newPost = "";
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
                  this.response = err.response.data[0];
              }
            }
          });
      }
    },
    clearMessage() {
      this.newPost = "";
      this.hint = "";
    },
    //入力フォームのラベルを更新する
    updateLabel() {
      this.label = "#" + this.channel.Name + "へメッセージを投稿する";
    },
  },
};
</script>

<style>
.input-form {
  position: fixed;
  margin-top: 85vh;
  width: calc(100vw - 256px);
  padding-top: 1vh;
  padding-left: 1vw;
  border-width: 1px 0 0 0;
  border-style: solid;
  border-color: #909090;
}
.input-box {
  width: 98%;
}
</style>