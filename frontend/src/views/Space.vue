<template>
  <div class="mypage">
    <Systembar></Systembar>
    
    <MypageNavbar></MypageNavbar>

    <MypageHeader></MypageHeader>

    <router-view></router-view>

  </div>
</template>

<script>
import MypageNavbar from '@/components/mypage/Navbar.vue'
import MypageHeader from '@/components/mypage/Header.vue'
import Systembar from '@/components/mypage/Systembar.vue'
export default {
  name:"Space",
  components: {
    MypageNavbar,
    MypageHeader,
    Systembar,
  },
}
</script>

<style>

</style>