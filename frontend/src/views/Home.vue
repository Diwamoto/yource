<template>
  <div class="home">
    <Header></Header>
    <div class="title">
      <v-row>
        <v-col xs5 class="message">
          <h1>Your Space is here.</h1>
          <h2>自分を表現しよう。</h2>
        </v-col>
        <v-col xs6>
          <v-img
            class="team"
            src="/img/team.svg"
            width="500"
          ></v-img>
        </v-col>
      </v-row>
    </div>

    <div class="card">
      <v-row>
        <v-col xs5>
          <v-img
            class="hero"
            src="/img/hero.svg"
            width="500"
          ></v-img>
        </v-col>
        <v-col xs6 class="message">
          <h1>Bloadcast your mind.</h1>
          <h3>気持ちを発信しよう。</h3>
        </v-col>  
      </v-row>
    </div>

    <div class="card">
      <v-row>
        <v-col xs6 class="message">
          <h1>Discuss awesome agenda.</h1>
          <h3>未来を話し合おう。</h3>
        </v-col>
        <v-col xs5>
          <v-img
            class="chat"
            src="/img/chat.svg"
            width="500"
          ></v-img>
        </v-col>
      </v-row>
    </div>
    <Footer></Footer>
  </div>
</template>

<script>
import Header from '@/components/Header.vue'
import Footer from '@/components/Footer.vue'
export default {
  name: 'Home',
  components: {
    Header,
    Footer,
  }
}
</script>

<style>
.title {
  height: 1000px;
  padding-top: 10%;
  padding-left: 10%;
  margin: 0 auto;
  text-align: left;
}
.title h2{
  margin-top: 1%;
}
.message{
  margin-top: 5%;
  margin-left: 5%;
}

.card{
  margin-left: 5%;
  height: 700px;
}
.card .row{
  margin-top: 10%;
}
.card .team{
  margin-left: 10%;
}
.card .message{
  margin-left: 10%;
}
</style>