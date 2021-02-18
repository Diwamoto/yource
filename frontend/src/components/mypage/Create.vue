<template>
  <div class="space">

    <v-main>
      <v-content class="space-title">

        <h1>あなたのコミュニティもしくはチームの名前を教えてください</h1>
        <p>この名前がサブドメインになります。</p>
        <v-form>
          <v-text-field
              v-model="message"
              append-outer-icon="mdi-send"
              outlined
              clear-icon="mdi-close"
              clearable
              placeholder="myspace"
              label="コミュニティーの名前"
              type="text" 
              @click:append="toggleMarker"
              @click:append-outer="sendMessage"
              @click:prepend="changeIcon"
              @click:clear="clearMessage"
            ></v-text-field>
        </v-form>
      </v-content>
    </v-main>

  </div>
</template>

<script>
export default {
  name:"Create",
  data: function(){
    return {
      password: 'Password',
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
    sendMessage () {
      this.resetIcon()
      this.clearMessage()
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
  width: 700px;
  margin: 15% auto;
}
</style>