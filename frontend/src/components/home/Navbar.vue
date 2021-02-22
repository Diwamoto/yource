<template>
  <div class="navbar">
    <v-navigation-drawer
      v-model="drawer"
      color="blue darken-3"
      dark
      app
    >
      <v-sheet
        color="blue darken-3"
        class="pa-4"
      >

        <div><h2>{{ spaceName }}</h2></div>
      </v-sheet>

      <v-divider></v-divider>

      <v-list dense v-for="channel in channels" :key="channel.Name">
        <v-list-item
          link
        >
          <v-list-item-content>
            <v-list-item-title><h3>#{{ channel.Name }}</h3></v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list>

      <!-- チャンネル追加アクション -->
      <v-dialog
        v-model="dialog"
        max-width="600px"
      >
        <template v-slot:activator="{ on, attrs }">
          <v-list-item-content v-bind="attrs" v-on="on">
            <v-list-item-title><h5 style="margin-left: 12px; color:white"><v-icon dense>mdi-plus</v-icon>チャンネルを追加する</h5></v-list-item-title>
          </v-list-item-content>
        </template>
        <v-card>
          <v-card-title>
            <span class="headline">チャンネルを追加する</span>
            <v-spacer></v-spacer>
            <v-btn
              text
              @click="dialog = false"
            >
              <v-icon>mdi-close</v-icon>
            </v-btn>
          </v-card-title>
          <v-divider></v-divider>
          <v-card-text>
            <v-container>
              <v-row>
                <v-col dense cols="12">
                  チャンネル名
                </v-col>
                <v-col cols="12">
                  <v-text-field
                    placeholder="zatsudan"
                    required
                    outlined
                  ></v-text-field>
                </v-col>
                <v-col dense cols="12">
                  説明
                </v-col>
                <v-col cols="12">
                  <v-text-field
                    label="説明"
                    placeholder="このチャンネルは〇〇の為の物です。"
                    outlined
                  ></v-text-field>
                </v-col>
              </v-row>
            </v-container>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn
              color="blue darken-1"
              text
              @click="addChannel"
            >
              <v-icon>mdi-plus</v-icon>追加する
            </v-btn>
          </v-card-actions>
        </v-card>
      </v-dialog>
      <!-- /チャンネル追加アクション -->
    </v-navigation-drawer>
  </div>
</template>

<script>
export default {
  name: "Navbar",
  props: {
    spaceName: String,
    channels: Array
  },
  data: function () {
    return {
      drawer: null,
      dialog: false,
    }
  },
  methods: {
    addChannel(){
    }
  }
}
</script>

<style>

</style>