<template>
  <v-app id="inspire">
    
    <Systembar :space-name="space.Name"></Systembar>

    <Navbar :space-name="space.Name" :channels="channels"></Navbar>

    <Main :channel="channel"></Main>

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
  data () {
    return {
      userId: null,
      space: [],
      channels: [],
      channel: [] 
    }
  },
  created() {
    //ユーザIDを取得してくる
    this.$http.get('https://' + this.$api + '/api/v1/retrive',{
      headers: {
        "Authorization" : "Bearer " + this.$cookies.get("token")
      },
      withCredentials: true
    })
    .then(response => {

      switch (response.status){
        case 200: //ユーザ情報を取得
          this.userId = response.data["userId"]

          //スペースの名前を取得してくる
          this.$http.get('https://' + this.$api + '/api/v1/users/' + this.userId + '/space',{
            headers: {
              "Authorization" : "Bearer " + this.$cookies.get("token")
            },
            withCredentials: true
          })
          .then(response => {

            switch (response.status){
              case 200: //名前を挿入
                this.space = response.data
                this.channels = this.space.Channels
                this.channel = this.channels[0]
              }
            })
          .catch(()=> {
            //ここに来るのは本来発生し得ないが、何らかのタイミングでスペースが消えてしまった場合は
            //↑再ログイン時にhomeに強制的に飛ばすため発生する
            //作成ページに飛ばす
            this.$router.push( { path: "create" }).catch((err)=>{ console.log(err)});
          })
        
        }
      })
    .catch(()=> {
      //ユーザ情報が取得できなかったら再ログインさせる
      this.$cookies.remove("token")
      this.$cookies.set("msg", "続けるにはログインが必要です。", 3600, "/", "localhost", true, "None")
      this.$router.push( { path: "login" }).catch((err)=>{ console.log(err)});
    })
  },
}
</script>