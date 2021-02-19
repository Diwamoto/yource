<template>
  <v-app id="inspire">
    
    <Systembar></Systembar>

    <Navbar></Navbar>

    <Main></Main>
  </v-app>
</template>

<script>
import Systembar from '@/components/home/Systembar.vue'
import Navbar from '@/components/home/Navbar.vue'
import Main from '@/components/home/Main.vue'
export default {
  components: {
    Systembar,
    Navbar,
    Main
  },
  name : "Home",
  data: () => ({
  }),
}
</script>