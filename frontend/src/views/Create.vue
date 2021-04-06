<template>
  <div class="space">
    <v-main class="space-box rounded-lg">
      <h1 class="space-title">セットアップ</h1>
      <h5 class="space-subtitle ml-12 mt-12">
        あなたの情報を教えて下さい。
      </h5>
      <h5 class="space-subtitle ml-12">
        素敵な名前をつけましょう！
      </h5>
      <v-form>
        <v-text-field
          v-model="space"
          :rules="[space_rule]"
          class="ma-12 mt-12 text-center"
          append-outer-icon="mdi-send"
          outlined
          clear-icon="mdi-close"
          clearable
          counter="25"
          placeholder="myspace"
          label="スペースの名前"
          type="text"
          @click:append="toggleMarker"
          @click:append-outer="submit"
          @click:prepend="changeIcon"
          @click:clear="clearMessage"
        ></v-text-field>
      </v-form>
    </v-main>
  </div>
</template>

<script>
export default {
  name: "Create",
  data: function () {
    return {
      userId: "",
      space_rule: () =>
        (this.space != "ws" && this.space != "api" && !!this.space && /^[a-zA-Z]*$/.test(this.space)) ||
        "スペース名に使えるのは半角英字のみです。",
      show: false,
      space: "",
      marker: true,
      iconIndex: 0,
      icons: [
        "mdi-emoticon",
        "mdi-emoticon-cool",
        "mdi-emoticon-dead",
        "mdi-emoticon-excited",
        "mdi-emoticon-happy",
        "mdi-emoticon-neutral",
        "mdi-emoticon-sad",
        "mdi-emoticon-tongue",
      ],
    };
  },
  computed: {
    icon() {
      return this.icons[this.iconIndex];
    },
  },

  methods: {
    toggleMarker() {
      this.marker = !this.marker;
    },
    submit() {
      //スペースを作成
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
              var params = new URLSearchParams();
              params.append("Name", this.space);
              params.append("SubDomain", this.space);
              this.$http
                .post(
                  this.$api + "/v1/users/" + this.userId + "/space",
                  params,
                  {
                    headers: {
                      Authorization: "Bearer " + this.$cookies.get("token"),
                    },
                    withCredentials: true,
                  }
                )
                .then((response) => {
                  //スペースの作成終了
                  var space = response.data;

                  //デフォルトで#mainと#randomを追加する
                  params = new URLSearchParams();
                  params.append("Name", "main");
                  params.append("Description", "メインチャンネルです。");
                  this.$http
                    .post(
                      this.$api + "/v1/spaces/" + space.Id + "/channels",
                      params,
                      {
                        headers: {
                          Authorization: "Bearer " + this.$cookies.get("token"),
                        },
                        withCredentials: true,
                      }
                    )
                    .then(() => {})
                    .catch(() => {});
                  params = new URLSearchParams();
                  params.append("Name", "random");
                  params.append(
                    "Description",
                    "仕事以外の話はこちらでした方がいいでしょう。"
                  );
                  this.$http
                    .post(
                      this.$api + "/v1/spaces/" + space.Id + "/channels",
                      params,
                      {
                        headers: {
                          Authorization: "Bearer " + this.$cookies.get("token"),
                        },
                        withCredentials: true,
                      }
                    )
                    .then(() => {})
                    .catch(() => {});

                  //作成したスペースへ移動する
                  this.$router.push({ path: "home" }).catch(() => {});
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
        });
    },
    clearMessage() {
      this.space = "";
    },
    resetIcon() {
      this.iconIndex = 0;
    },
    changeIcon() {
      this.iconIndex === this.icons.length - 1
        ? (this.iconIndex = 0)
        : this.iconIndex++;
    },
  },
};
</script>

<style>
.space {
  height: 100%;
  background-color: #f3f3f3;
}
.space-box {
  background-color: white;
  margin: 15vh 30vw 10vh 30vw;
  height: 50%;
  width: 40vw;
  opacity: 0;
  box-shadow: 0 6px 6px -3px rgba(0, 0, 0, 0.2),
    0 10px 14px 1px rgba(0, 0, 0, 0.14), 0 4px 18px 3px rgba(0, 0, 0, 0.12) !important;
  animation-name: space;
  animation-duration: 0.5s;
  animation-timing-function: ease-out;
  animation-fill-mode: forwards;
}
.space-title {
  padding: 2% 0 2% 0;
  text-align: center;
  border-width: 0 0 2px 0;
  border-style: solid;
  border-color: #f0f0f0;
}
.privacy-subtitle {
  padding: 3% 10% 0 10%;
}
@keyframes space {
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