<template>
  <div class="home">
    <Header></Header>
    <v-main id="main">
      <v-container>
        <div class="title">
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
        <v-scroll-x-transition>
          <div class="about">
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
        </v-scroll-x-transition>
        <v-scroll-x-transition>
          <div class="card">
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
        </v-scroll-x-transition>
        <v-scroll-x-transition>
          <div class="about">
            <v-row>
              <v-col md6 offset-md="3" class="message">
                <h1>活動を発信する</h1>
                <h4>自分個人の価値を他の誰かに発信することができます。</h4>
                <h4>活動を発信する場として活用しましょう。</h4>
              </v-col>
            </v-row>
          </div>
        </v-scroll-x-transition>
        <v-scroll-x-transition>
          <div class="card">
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
        </v-scroll-x-transition>
      </v-container>
    </v-main>
    <Footer></Footer>
  </div>
</template>

<script>
import Header from "@/components/Header.vue";
import Footer from "@/components/Footer.vue";
export default {
  name: "Home",
  data: function () {
    return {};
  },
  components: {
    Header,
    Footer,
  },
};
</script>

<style>
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
</style>