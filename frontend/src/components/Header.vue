<template>
  <div class="header">
    <v-app-bar flat class="app_bar" color="white" app elevate-on-scroll scroll-target="card">
      <v-row>
        <!-- logo -->
        <v-col md="1" lg="2" offset-md="1">
          <router-link to="/">
            <v-img
              class="main-logo_image"
              src="/img/logo_transparent.png"
              width="150"
            ></v-img>
          </router-link>
        </v-col>
      
        <!-- /logo -->

        <!-- menu -->
        <v-col md="1" offset-md="6" class="menu-element">
          <router-link to="/about">概要</router-link>
        </v-col>

        <v-col md="1" class="menu-element btn-join">
          <router-link to="/join">
            <v-btn
              color="red"
              elevation="5"
              large
              outlined
              rounded
            >新規登録</v-btn>
          </router-link>
        </v-col>
        <v-col md="1" class="hamburger">
          <v-app-bar-nav-icon></v-app-bar-nav-icon>
        </v-col>
        <!-- /menu -->
      </v-row>  
    </v-app-bar> 
  </div>
</template>

<script>
export default {
  name: 'Header',
  
}
</script>

<style>
.header{
  height: 5%;
  margin-top: 30px;
  margin-left: 70px; 
  margin-bottom: 30px;
}
.main-logo_image{
  margin-top: 16px;
  margin-left: 2%;
}
.menu-element{
  margin-left: -1%;
  margin-top: 40px;
}
.menu-element a{
  color: black;
  text-decoration: none;
}
.btn-join{
  margin-top: 30px;
  color: white;
}
.hamburger{
  margin-top: 25px;
}
</style>