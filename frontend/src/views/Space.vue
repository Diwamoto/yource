<template>
  <v-app id="inspire">
    
    <Systembar></Systembar>

    <Navbar></Navbar>

    <router-view></router-view>
  </v-app>
</template>

<script>
import Systembar from '@/components/mypage/Systembar.vue'
import Navbar from '@/components/mypage/Navbar.vue'
export default {
  components: {
    Systembar,
    Navbar
  },
  name : "Space",
  data: () => ({
  }),
}
</script>