<template>
  <div class="navbar">
    <v-navigation-drawer
      app
      color="#808080"
    >
      <v-list>
        <v-list-item v-for="ch in channels" :key="ch.name" :to="ch.link">
          <v-list-item-icon>
            <v-icon>mdi-table</v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title>{{ ch.name }}</v-list-item-title>
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
      channels: [
        { name: 'main', link: "/channels/main"},
        { name: 'zatsudan', link: "/channels/zatsudan"},
      ],
    }
  },
}
</script>

<style>

</style>