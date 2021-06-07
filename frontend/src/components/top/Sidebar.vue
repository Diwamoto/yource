<template>
  <v-navigation-drawer
    v-model="hamburger"
    absolute
    temporary
    right
  >
    <v-list-item>

      <v-list-item-content>
        <v-list-item-title class="text-h6">
            Yource
          </v-list-item-title>
      </v-list-item-content>
    </v-list-item>

    <v-divider></v-divider>

    <v-list dense>
      <v-list-item
        v-for="item in items"
        :key="item.name"
        :href="item.link"
        link
      >
        <v-list-item-content>
          <v-list-item-title>{{ item.name }}</v-list-item-title>
        </v-list-item-content>
      </v-list-item>
    </v-list>
  </v-navigation-drawer>
</template>

<script>
export default {
    name:'Sidebar',
    data: function(){
      return{
        hamburger: null,
        items: [{'name': 'ログイン', 'link': 'login'}, {'name': '概要', 'link': 'about'}, {'name': 'プライバシーポリシー', 'link':'privacy'}, {'name': '制作者について', 'link':'https://diwamoto.dev'}],
      }
      
    },
    methods:{
      toggle(){
        this.hamburger = !this.hamburger;
      }
    }
}
</script>

<style>

</style>