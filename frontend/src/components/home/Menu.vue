<template>
  <div class="menu">
    <div class="menu-items-wrap">
      <div class="menu-items-channel">#{{ channel.Name }}</div>

      <div class="menu-items-icons">
        <v-dialog v-model="dialog" width="500" height="400">
          <template v-slot:activator="{ on, attrs }">
            <v-btn text icon v-bind="attrs" v-on="on" style="margin-right: 10%">
              <v-icon large>mdi-information-outline</v-icon>
            </v-btn>
          </template>

          <v-card height="300">
            <v-card-title class="headline">
              <span>#{{ channel.Name }}</span>
              <v-spacer></v-spacer>
              <v-btn text icon v-on:click="dialog = false"
                ><v-icon>mdi-close</v-icon></v-btn
              >
            </v-card-title>
            <v-divider></v-divider>
            <v-card-text>
              <span>{{ channel.Description }}</span>
            </v-card-text>
          </v-card>
        </v-dialog>
        <router-link to="/user/profile"
          ><v-icon large>mdi-cog</v-icon></router-link
        >
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "Menu",
  props: {
    channel: [Array, Object],
  },
  data: function () {
    return {
      dialog: false,
    };
  },
};
</script>

<style>
.menu {
  position: fixed;
  height: 7vh;
  width: 100vw;
  border-width: 0 0 1px 0;
  border-style: solid;
  border-color: #d0d0d0;
}
.menu-items-wrap {
  margin-left: 1.5vw;
  padding-top: 2vh;
}
.menu-items-channel {
  float: left;
  font-weight: 900;
}
.menu-items-icons {
  margin-top: -5px;
  margin-left: 78vw;
}
.menu-items-icons a {
  text-decoration: none;
}
.v-card__text {
  margin-top: 10px;
}
</style>