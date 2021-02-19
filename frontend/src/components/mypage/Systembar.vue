<template>
  <div class="system">
    <v-system-bar color="primary" dark app>
      <v-spacer></v-spacer>
      <span>yource - {{ space }}</span>
      <v-spacer></v-spacer>
      <v-icon>mdi-wifi-strength-4</v-icon>
      <v-icon>mdi-signal-cellular-outline</v-icon>
      <v-icon>mdi-battery</v-icon>
      <span>12:30</span>
    </v-system-bar>
  </div>
</template>

<script>
//import axios from 'axios'
export default {
  name: "Systembar",
  data: () => {
    return {
      space:  "",
      userId: ""
    }
  },
  created() {
    console.log(this.$cookies.keys())
    //スペースを取得してくる
    // this.userId = this.$cookies.get("userId")
    // axios.get('https://' + this.url + '/api/v1/users/' + this.userId + '/space',{
    //   headers: {
    //     "Authorization" : "Bearer " + this.$cookies.keys()
    //   },
    //   withCredentials: true
    // })
    // .then(response => {

    //   switch (response.status){
    //     case 200: //名前を挿入
    //       this.space = response.data
    //     }
    //   })
    // .catch(()=> {
    // })
  }
}
</script>

<style>

</style>