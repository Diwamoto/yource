<template>
  <div class="header">
    <v-row>
      <!-- logo -->
      <!-- TODO:タイトルとロゴを統一した画像にする -->
      
      <v-col sm="1">
        <router-link to="/summary">
          <v-img
            width="75"
            class="main-logo_image"
            src="/img/logo.png"
          ></v-img>
        </router-link>
      </v-col>
      <v-col sm="3">
        <router-link to="/summary">
        <v-img
          width="130"
          class="main-logo_name"
          src="/img/logo_name.png"
        ></v-img>
         </router-link>
      </v-col>
     
      <!-- /logo -->

      <!-- menu -->
      <v-col sm="1" offset-sm="3" class="menu-element">
        <router-link to="/summary">概要</router-link>
      </v-col>

      <v-col sm="1" class="menu-element">
        <router-link to="/howtojoin">参加方法</router-link>
      </v-col>

      <v-col sm="1" class="menu-element">
        <router-link to="/join">会員登録</router-link>
      </v-col>

      <!-- /menu -->
     
    </v-row>   
  </div>
</template>

<script>
// @ is an alias to /src

export default {
  name: 'Home',
  
}
</script>

<style>
.header{
  margin-top: 30px;
  margin-left: 70px; 
}
.main-logo_name{
  margin-top: 16px;
  margin-left: -20px;
}
.menu-element{
  margin-top: 40px;
}
.menu-element a{
  color: black;
  text-decoration: none;
}
</style>