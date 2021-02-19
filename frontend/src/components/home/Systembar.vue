<template>
  <div class="system">
    <v-system-bar color="primary" dark app>
      <v-spacer></v-spacer>
      <span>yource - {{ spaceName }}</span>
      <v-spacer></v-spacer>
      <span>{{ date }}</span>
      <span>{{ time }}</span>
    </v-system-bar>
  </div>
</template>

<script>
export default {
  name: "Systembar",
  data: () => {
    return {
      date: "",
      time: "",
      month: 0,
      day: 0,
      hour: 0,
      minute: 0,
      prefix: "",//午前 or 午後
      week: ['(日)', '(月)', '(火)', '(水)', '(木)',  '(金)', '(土)'] ,
      spaceName:  "",
      userId: ""
    }
  },
  created() {

      //現在日時を取得
      //先にコンマを消す処理が走るため
      const JST = new Date().toLocaleString({ timeZone: 'Asia/Tokyo' });
      let now = new Date(JST)
      this.month = now.getMonth() + 1
      this.day = now.getDate()
      this.hour = now.getHours()
      //午前か午後の判定
      if (this.hour > 12){
        this.prefix = "午後"
        this.hour -= 12
      } else {
        this.prefix = "午前"
      }
      this.minute = now.getMinutes()

    //スペースを取得してくる
    this.userId = this.$cookies.get("id")
    this.$http.get('https://' + this.$api + '/api/v1/users/' + this.userId + '/space',{
      headers: {
        "Authorization" : "Bearer " + this.$cookies.get("token")
      },
      withCredentials: true
    })
    .then(response => {

      switch (response.status){
        case 200: //名前を挿入
          this.spaceName = response.data.Name
        }
      })
    .catch(()=> {
    })
  },
  mounted: function()  { 
    setInterval(this.renderComma, 500); 
    setInterval(this.updateTime, 1000); 
  },
  methods: {
    renderComma: function(){
      //間のコンマを描画しない
      this.time = '　' + this.prefix + this.zeroPadding(this.hour, 2) + ' ' + this.zeroPadding(this.minute, 2)
      this.date = this.month + "月" + this.zeroPadding(this.day, 2) + "日" + this.week[this.day % 7]
    },
    updateTime: function() { 
      //現在日時を取得
      const JST = new Date().toLocaleString({ timeZone: 'Asia/Tokyo' });
      let now = new Date(JST)
      this.month = now.getMonth() + 1
      this.day = now.getDate()
      this.hour = now.getHours()
      //午前か午後の判定
      if (this.hour > 12){
        this.prefix = "午後"
        this.hour -= 12
      } else {
        this.prefix = "午前"
      }
      this.minute = now.getMinutes()

      //表示用変数に代入
      this.time = '　' + this.prefix + this.zeroPadding(this.hour, 2) + ':' + this.zeroPadding(this.minute, 2)
      this.date = this.month + "月" + this.zeroPadding(this.day, 2) + "日" + this.week[this.day % 7]
    },
    zeroPadding: function(num, len) {
      let zero = '';

      // 0の文字列を作成
      for(var i = 0; i < len; i++) {
        zero += '0';
      }

      // zeroの文字列と、数字を結合し、後ろ２文字を返す
      return (zero + num).slice(-len);
    }
  }
}
</script>

<style>

</style>