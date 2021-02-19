<template>
  <div class="navbar">
    <v-navigation-drawer
      v-model="drawer"
      color="indigo darken-4"
      dark
      app
    >
      <v-sheet
        color="indigo darken-4"
        class="pa-4"
      >

        <div><h2>{{ spaceName }}</h2></div>
      </v-sheet>

      <v-divider></v-divider>

      <v-list dense>
        <v-list-item
          link
        >
          <v-list-item-content>
            <v-list-item-title><h3>#main</h3></v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>
  </div>
</template>

<script>
export default {
  name: "Navbar",
  data: function () {
    return {
      drawer: null,
      userId: this.$cookies.get("id"),
      spaceName: "",
      channels: [],
    }
  },
  created() {
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
  }
}
</script>

<style>

</style>