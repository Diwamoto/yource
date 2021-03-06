<template>
  <div class="about">
    <Header></Header>
    <Menu></Menu>
    <div class="about-box rounded-lg">
      <h1 class="about-title">yourceについて</h1>
      <div class="about-body-wrapper">
        <h1 class="about-main-title">Let's talk about the future.</h1>
        <div class="about-body about-img-wrapper">
          <v-img class="about-img" src="/img/about.svg"></v-img>
        </div>
      </div>
      <div class="about-body-wrapper">
        <div class="about-body">
          yourceは自分のポートフォリオを公開し、自分のコミュニティを作ることのできる場所です。
        </div>
        <div class="about-body">
          slackのような見た目で自分をフォローしてくれる人たちとチャットすることができます。
        </div>
      </div>
    </div>
    <Footer></Footer>
  </div>
</template>

<script>
import Header from "@/components/top/Header.vue";
import Footer from "@/components/top/Footer.vue";
import Menu from "@/components/top/Menu.vue";
export default {
  name: "About",
  components: {
    Header,
    Footer,
    Menu,
  }
};
</script>

<style>
.about {
  height: 100%;
  background-color: #f3f3f3;
}
.about-box {
  background-color: white;
  margin: 8vh 25vw 10vh 25vw;
  height: 100vh;
  width: 50vw;
  opacity: 0;
  box-shadow: 0 6px 6px -3px rgba(0, 0, 0, 0.2),
    0 10px 14px 1px rgba(0, 0, 0, 0.14), 0 4px 18px 3px rgba(0, 0, 0, 0.12) !important;
  animation-name: about;
  animation-duration: 0.5s;
  animation-timing-function: ease-out;
  animation-fill-mode: forwards;
}
.about-title {
  padding: 2% 0 2% 0;
  text-align: center;
  border-width: 0 0 2px 0;
  border-style: solid;
  border-color: #f0f0f0;
}
.about-main-title{
  font-size: 4.5em;
  padding: 2% 0 2% 0;
  text-align: center;
}
.about-img-wrapper{
  margin: 0 2vw 0 2vw;
}
.about-img{
  margin: 0 5vw 0 5vw;
  width: 70%;
}
.about-subtitle {
  padding: 3% 10% 0 10%;
}
.about-body-wrapper {
  padding: 6% 10% 0 10%;
}
.about-body{
  padding: 0 5% 0 5%;
}
@keyframes about {
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