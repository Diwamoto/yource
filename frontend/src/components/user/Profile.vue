<template>
  <div class="profile">
    <v-form class="profile-form rounded-lg" elevation="12" @submit.prevent>
      <h1 class="profile-title">プロフィール</h1>
      <div class="spacer"></div>
      <v-row>
        <v-col cols=5>icon</v-col>
        <v-col cols=6>名前：{{user.Name}}<router-link :to="user"><small>変更する</small></router-link></v-col>
      </v-row>
      <v-row>
        <v-col offset=5>ニックネーム：{{user.Nickname}}<router-link :to="user"><small>変更する</small></router-link></v-col>
      </v-row>
      <div class="profile-form-field">
        <v-text-field
          v-model="profile.Profile"
          label="自己紹介"
          required
          outlined
          prepend-icon="mdi-account-box"
        ></v-text-field>
      </div>
      <div class="profile-form-field">
        <v-menu
          ref="menu"
          v-model="menu"
          :close-on-content-click="false"
          :return-value.sync="profile.Birthday"
          transition="scale-transition"
          offset-y
          min-width="auto"
        >
          <template v-slot:activator="{ on, attrs }">
            <v-text-field
              v-model="profile.Birthday"
              label="誕生日"
              required
              outlined
              prepend-icon="mdi-cake-variant"
              class="profile-form-field"
              v-bind="attrs"
              v-on="on"
            ></v-text-field>
          </template>
          <v-date-picker v-model="profile.Birthday" no-title scrollable>
            <v-spacer></v-spacer>
            <v-btn text color="primary" @click="menu = false"> Cancel </v-btn>
            <v-btn text color="primary" @click="$refs.menu.save(profile.Birthday)">
              OK
            </v-btn>
          </v-date-picker>
        </v-menu>
      </div>
      <div class="profile-form-field">
        <v-text-field
          v-model="profile.Hometown"
          label="出身地"
          required
          outlined
          prepend-icon="mdi-home-city"
          class="profile-form-field"
        ></v-text-field>
      </div>
      <div class="profile-form-field">
        <v-text-field
          v-model="profile.Job"
          label="職業"
          required
          outlined
          prepend-icon="mdi-briefcase"
          class="profile-form-field"
        ></v-text-field>
      </div>
      <div class="profile-form-field">
        <v-text-field
          v-model="profile.Twitter"
          label="ツイッターURL"
          required
          outlined
          prepend-icon="mdi-twitter"
          class="profile-form-field"
        ></v-text-field>
      </div>
      <div class="profile-form-field">
        <v-text-field
          v-model="profile.Facebook"
          label="フェイスブックURL"
          required
          outlined
          prepend-icon="mdi-facebook"
          class="profile-form-field"
        ></v-text-field>
      </div>
      <div class="profile-form-field">
        <v-text-field
          v-model="profile.Instagram"
          label="インスタグラムURL"
          required
          outlined
          prepend-icon="mdi-instagram"
          class="profile-form-field"
        ></v-text-field>
      </div>
      <div class="profile-form-field">
        <v-text-field
          v-model="profile.Other"
          label="その他webサイトURL"
          required
          outlined
          prepend-icon="mdi-web"
          class="profile-form-field"
        ></v-text-field>
      </div>
      <v-dialog
        v-model="dialog"
        width="600px"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-btn class="profile-form-submit" color="primary" v-bind="attrs" v-on="on"> 変更する </v-btn>
        </template>
        <v-card>
          <v-card-title>
            <span class="headline">以下の内容でよろしいですか？</span>
          </v-card-title>
          <v-card-text>
            <p>自己紹介：{{omit(profile.Profile)}}</p>
            <p>誕生日：{{profile.Birthday}}</p>
            <p>出身地：{{profile.Hometown}}</p>
            <p>職業：{{profile.Job}}</p>
            <p>ツイッターURL：{{profile.Twitter}}</p>
            <p>フェイスブックURL：{{profile.Facebook}}</p>
            <p>インスタグラムURL：{{profile.Instagram}}</p>
            <p>その他webサイトURL：{{profile.Other}}</p>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
              color="red darken-1"
              text
              @click="dialog = false"
            >
              キャンセル
            </v-btn>
            <v-btn
              color="green darken-1"
              text
              @click="dialog = false;submit();"
              type="submit"
            >
              変更する
            </v-btn>
          </v-card-actions>
        </v-card>
     </v-dialog>
    </v-form>
    <Footer></Footer>
  </div>
</template>

<script>
import Footer from "@/components/top/Footer.vue";
export default {
  components: { Footer },
  props: {
    flashMsg: {
      type: Function,
      required: true,
    },
  },
  data() {
    return {
      user: null,
      profile: null,
      menu: false,
      dialog: false,
    };
  },
  created() {
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
            this.$http
              .get(this.$api + "/v1/users/" + this.userId, {
                headers: {
                  Authorization: "Bearer " + this.$cookies.get("token"),
                },
                withCredentials: true,
              })
              .then((response) => {
                this.user = response.data[0];
                this.profile = this.user.Profile;
                this.profile.Birthday = this.profile.Birthday.split("T")[0]
              })
              .catch(() => {
                this.toLoginPage();
              });
        }
      })
      .catch(() => {
        this.toLoginPage();
      });
  },
  methods: {
    //自己紹介を省略する
    omit(s){
      return s.length > 100 ? s.slice(0, 100) + "…" : s;
    },
    //送信する
    submit() {
      //ポスト投げる
      const params = new URLSearchParams();
        params.append("Profile", this.profile.Profile);
        params.append("Birthday", this.profile.Birthday);
        params.append("Hometown", this.profile.Hometown);
        params.append("Job", this.profile.Job);
        params.append("Twitter", this.profile.Twitter);
        params.append("Facebook", this.profile.Facebook);
        params.append("Instagram", this.profile.Instagram);
        params.append("Other", this.profile.Other);
    this.$http
      .put(this.$api + "/v1/users/" + this.userId + "/profile", params, {
        headers: {
          Authorization: "Bearer " + this.$cookies.get("token"),
        },
        withCredentials: true,
      })
      .then((response) => {
        switch (response.status) {
          case 200: //変更成功
           this.flashMsg("success","ユーザの情報を変更しました。")
        }
      })
      .catch((err) => {
        this.flashMsg("error",err.response.data.error)
      });
    },
  },
};
</script>

<style>
.profile {
  background-color: #e2e2e2;
}
.profile-form {
  margin: 10vh 25vw 10vh 25vw;
  width: 50vw;
  height: 100%;
  background-color: white;
  box-shadow: 0 6px 6px -3px rgba(0, 0, 0, 0.2),
    0 10px 14px 1px rgba(0, 0, 0, 0.14), 0 4px 18px 3px rgba(0, 0, 0, 0.12) !important;
  opacity: 0;
  animation-name: profile;
  animation-duration: 0.5s;
  animation-timing-function: ease-out;
  animation-fill-mode: forwards;
}
.profile-title {
  padding: 2% 0 2% 0;
  text-align: center;
  border-width: 0 0 2px 0;
  border-style: solid;
  border-color: #f0f0f0;
}
.profile-form-field {
  margin: 1vh 7vw 1vh 7vw;
}
.spacer {
  height: 30px;
}
.profile-form-submit{
  margin: 1vh 0 3vh 45.5%;
}
@keyframes profile {
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