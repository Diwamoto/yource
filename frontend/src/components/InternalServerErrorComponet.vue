<template>
  <div class="notfound">

    <Header></Header> 
    <v-img
        class="notfound-img"
        src="/img/500.svg"
    ></v-img>
    <div class="message text-center">
        <h3>500 InternalServerError</h3>
        <p>サーバ内部のエラーにより表示できません。</p> 
        <p>しばらく時間をおいてからやり直してください。</p>
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
    name: "InternalServerErrorComponent",
    components: { Header, Footer },

}
</script>

<style>
.notfound{
    height: 100%;
}
.notfound-img{
    margin: 200px auto 100px auto;
    width: 500px;
}
.message a{
    color: black;
    text-decoration: none;
}
</style>