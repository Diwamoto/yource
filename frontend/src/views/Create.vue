<template>
  <div class="space">

    <v-main class="space-title">

        <h1>あなたのコミュニティもしくはチームの名前を教えてください</h1>
        <p>素敵な名前をつけましょう！</p>
        
        <v-form>
              <v-text-field
                v-model="message"
                append-outer-icon="mdi-send"
                outlined
                clear-icon="mdi-close"
                clearable
                counter="25"
                placeholder="myspace"
                label="コミュニティーの名前"
                type="text" 
                @click:append="toggleMarker"
                @click:append-outer="submit"
                @click:prepend="changeIcon"
                @click:clear="clearMessage"
              ></v-text-field>
          
        </v-form>
    </v-main>

  </div>
</template>

<script>
export default {
  name:"Create",
  data: function(){
    return {
      userId: "",
      show: false,
      message: '',
      marker: true,
      iconIndex: 0,
      icons: [
        'mdi-emoticon',
        'mdi-emoticon-cool',
        'mdi-emoticon-dead',
        'mdi-emoticon-excited',
        'mdi-emoticon-happy',
        'mdi-emoticon-neutral',
        'mdi-emoticon-sad',
        'mdi-emoticon-tongue',
      ],
    }
  },
  computed: {
    icon () {
      return this.icons[this.iconIndex]
    },
  },

  methods: {
    toggleMarker () {
      this.marker = !this.marker
    },
    submit () {
      //スペースを作成
      this.userId = this.$cookies.get("id")
      const params = new URLSearchParams();
      params.append("Name", this.message)
      params.append("SubDomain", this.message)
      this.$http.post('https://' + this.$api + '/api/v1/users/' + this.userId + '/space',params,{
        headers: {
          "Authorization" : "Bearer " + this.$cookies.get("token")
        },
        withCredentials: true
      }).then(() => {
        //スペースの作成終了
        //作成したスペースへ移動する
        this.$router.push( { path: "home" }).catch(()=>{});
      }).catch(err => {
        if(err.response) {
          switch (err.response.status){
            case 404:
              this.Conflict = true
              this.response = "サーバーが一時的にダウンしています。管理者に連絡してください。"
              break
            case 409:
              this.Conflict = true
              this.response = err.response.data[0]
          }
        }
      });
    },
    clearMessage () {
      this.message = ''
    },
    resetIcon () {
      this.iconIndex = 0
    },
    changeIcon () {
      this.iconIndex === this.icons.length - 1
        ? this.iconIndex = 0
        : this.iconIndex++
    },
  },
}
</script>

<style>
.space-title{
  width: 900px;
  margin: 15% auto;
}
</style>