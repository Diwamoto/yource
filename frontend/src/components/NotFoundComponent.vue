<template>
  <div class="notfound">

    <Header></Header> 
    <v-img
        class="notfound-img"
        src="/img/404.svg"
    ></v-img>
    <div class="message text-center">
        <h3>404 Notfound</h3>
        <p>ページが見つかりません。</p> 
        <p>お探しのページは存在しないか、削除された可能性があります。</p>
        <router-link to="/"><v-btn
  elevation="2"
  outlined
  raised
>トップページに戻る</v-btn></router-link>
    </div>
    <Footer></Footer>
  </div>
</template>

<script>
import Header from '@/components/top/Header.vue'
import Footer from '@/components/top/Footer.vue'
export default {
    name: "NotFoundComponent",
    components: { Header, Footer},

}
</script>

<style>
.notfound{
    height: 100vh;
}
.notfound-img{
    margin: 20vh auto 10vh auto;
    width: 30%;
}
.message{
    height: 23.25vh;
}
.message a{
    color: black;
    text-decoration: none;
}
</style>