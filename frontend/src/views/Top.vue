<template>
  <div class="home" v-scroll="handleScroll">
    <Header></Header>
    <v-main id="main">
      <v-container>
        <div :class="{title:true, fadein_top:true}"> 
          <v-row>
            <v-col xs5 class="message">
              <h1>Your Space is here.</h1>
              <h2>自分を表現しよう。</h2>
            </v-col>
            <v-col xs6>
              <v-img class="team" src="/img/team.svg" width="500"></v-img>
            </v-col>
          </v-row>
        </div>
        <div :class="{fadein_left:fadein1, hide:!fadein1}">
          <v-row>
            <v-col md6 offset-md="3" class="message">
              <h1>Yourceについて</h1>
              <h4>
                Yourceは自分のポートフォリオを公開し、自分のコミュニティを作ることのできる場所です。
              </h4>
              <h4>
                もちろん、ポートフォリオを公開せずに、自分の活動を表現する場所としても使えます。
              </h4>
            </v-col>
          </v-row>
        </div>
        <div :class="{card:true, fadein_right:fadein2, hide:!fadein2}">
          <v-row>
            <v-col xs5>
              <v-img class="hero" src="/img/hero.svg" width="500"></v-img>
            </v-col>
            <v-col xs6 class="message">
              <h1>Bloadcast your mind.</h1>
              <h3>気持ちを発信しよう。</h3>
            </v-col>
          </v-row>
        </div>
        <div :class="{fadein_left:fadein3, hide:!fadein3}">
          <v-row>
            <v-col md6 offset-md="3" class="message">
              <h1>活動を発信する</h1>
              <h4>自分個人の価値を他の誰かに発信することができます。</h4>
              <h4>活動を発信する場として活用しましょう。</h4>
            </v-col>
          </v-row>
        </div>
        <div :class="{card:true, fadein:fadein4, hide:!fadein4}">
          <v-row>
            <v-col xs6 class="message">
              <h1>Discuss awesome agenda.</h1>
              <h3>未来を話し合おう。</h3>
            </v-col>
            <v-col xs5>
              <v-img class="chat" src="/img/chat.svg" width="500"></v-img>
            </v-col>
          </v-row>
        </div>
      </v-container>
    </v-main>
    <Footer></Footer>
  </div>
</template>

<script>
import Header from "@/components/top/Header.vue";
import Footer from "@/components/top/Footer.vue";

export default {
  name: "Home",
  data: function () {
    return {
      fadein1: false,
      fadein2: false,
      fadein3: false,
      fadein4: false,
    };
  },
  components: {
    Header,
    Footer,
  },
  mounted(){
    this.$setTitle("yource | yourceはポートフォリオを簡単に公開できるサービスです。", "yourceはポートフォリオを簡単に公開できるサービスです。")
  },
  methods: {
    //スクロールに応じて表示する内容を決める
    handleScroll: function() {
      if (window.scrollY > 85){
        this.fadein1 = true
      }
      if (window.scrollY > 615){
        this.fadein2 = true
      }
      if (window.scrollY > 1136){
        this.fadein2 = true
      }
      if (window.scrollY > 1350){
        this.fadein3 = true
      }
      if (window.scrollY > 1802){
        this.fadein4 = true
      }
    }
  }
};
</script>

<style>
.home{
  height: 100vh;
}
.title {
  height: 900px;
  padding-top: 16%;
  padding-left: 10%;
  margin: 0 auto;
  text-align: left;
}
.title h2 {
  margin-top: 1%;
}
.message {
  margin-top: 5%;
  margin-left: 5%;
}
.card {
  margin-top: 15%;
  margin-left: 5%;
  height: 600px;
}
.card .row {
  margin-top: 10%;
}
.card .team {
  margin-left: 10%;
}
.card .message {
  margin-left: 10%;
}
.fadein{
  animation: fadein 1s;
}
.hide{
  opacity: 0;
}
.fadein_top{
  animation: fadein_top 1.5s;
}
.fadein_left{
  animation: fadein_left 1.5s;
}
.fadein_right{
  animation: fadein_right 1.5s;
}


@keyframes fadein {
  0% {
    opacity: 0;
    transform: translateY(20px);
  }
  100% {
    opacity: 1;
    transform: translateY(0px);
  }
}
@keyframes fadein_top {
  0% {
    opacity: 0;
    transform: translateY(-20px);
  }
  100% {
    opacity: 1;
    transform: translateY(0px);
  }
}
@keyframes fadein_left {
  0% {
    opacity: 0;
    transform: translateX(-20px);
  }
  100% {
    opacity: 1;
    transform: translateY(0px);
  }
}
@keyframes fadein_right {
  0% {
    opacity: 0;
    transform: translateX(20px);
  }
  100% {
    opacity: 1;
    transform: translateY(0px);
  }
}

</style>