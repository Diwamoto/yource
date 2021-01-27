<template>
  <v-app>
  <v-app-bar
        color="white"
        dense
      >
        <v-app-bar-nav-icon>
          <v-img
            alt="Vuetify Logo"
            class=""
            contain
            src="/img/logo.png"
            transition="scale-transiton"
            width="60"
            height="40"
          />
        </v-app-bar-nav-icon>
  
        <v-toolbar-title>Yource</v-toolbar-title>
  
        <v-spacer></v-spacer>
  
        <v-menu
          right
          bottom
        >
          <template v-slot:activator="{ on, attrs }">
            <v-btn
              icon
              v-bind="attrs"
              v-on="on"
            >
              <v-icon>mdi-dots-vertical</v-icon>
            </v-btn>
          </template>
  
          <v-list>
            <v-list-item
              v-for="n in 5"
              :key="n"
              @click="() => {}"
            >
              <v-list-item-title>Option {{ n }}</v-list-item-title>
            </v-list-item>
          </v-list>
        </v-menu>
      </v-app-bar>
  </v-app>
</template>

<script></script>

<style>
.v-sheet .theme--light .v-toolbar .v-toolbar--dense .v-app-bar .v-app-bar--is-scrolled .white{
  height:30px
}
</style>