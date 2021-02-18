<template>
  <div class="notfound">

    <Header></Header> 
    <v-img
        class="notfound-img"
        src="/img/notfound.svg"
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
  </div>
</template>

<script>
import Header from '@/components/Header.vue'
export default {
    name: "NotFoundComponent",
    components: { Header },

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