<template>
  <div class="navbar">
    <v-navigation-drawer v-model="drawer" color="blue-grey darken-2" dark app>
      <v-sheet color="blue-grey darken-2" class="pa-4">
        <div>
          <h2>{{ space.Name }}</h2>
        </div>
      </v-sheet>
      <v-list class="channels">
        <v-list-item
          v-for="channel in channels"
          :key="channel.Name"
          @click="selectChannel(channel.Name)"
        >
          <v-list-item-content>
            <v-list-item-title
              ><h3>#{{ channel.Name }}</h3></v-list-item-title
            >
          </v-list-item-content>
        </v-list-item>
      </v-list>

      <!-- チャンネル追加アクション -->
      <v-dialog v-model="dialog" max-width="600px">
        <template v-slot:activator="{ on, attrs }">
          <v-list-item-content v-bind="attrs" v-on="on">
            <v-list-item-title
              ><h5 style="margin-left: 12px; color: white">
                <v-icon dense>mdi-plus</v-icon>チャンネルを追加する
              </h5></v-list-item-title
            >
          </v-list-item-content>
        </template>
        <v-card>
          <v-card-title>
            <span class="headline">チャンネルを追加する</span>
            <v-spacer></v-spacer>
            <v-btn text @click="dialog = false">
              <v-icon>mdi-close</v-icon>
            </v-btn>
          </v-card-title>
          <v-card-text>
            <v-container>
              <v-row>
                <v-col dense cols="12"> チャンネル名 </v-col>
                <v-col cols="12">
                  <v-text-field
                    v-model="channelName"
                    placeholder="zatsudan"
                    required
                    outlined
                  ></v-text-field>
                </v-col>
                <v-col dense cols="12"> 説明 </v-col>
                <v-col cols="12">
                  <v-text-field
                    v-model="channelDescription"
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
            <v-btn color="blue darken-1" text @click="addChannel">
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
    space: [Array, Object],
    channels: Array,
    userId: Number,
    switchChannel: {
      type: Function,
      required: true,
    },
  },
  data: function () {
    return {
      drawer: null,
      dialog: false,
      channelName: "",
      channelDescription: "",
    };
  },
  methods: {
    addChannel() {
      //チャンネル追加アクション
      const params = new URLSearchParams();
      params.append("Name", this.channelName);
      params.append("Description", this.channelDescription);
      this.$http
        .post(
          "https://" +
            this.$api +
            "/api/v1/spaces/" +
            this.space.Id +
            "/channels",
          params,
          {
            headers: {
              Authorization: "Bearer " + this.$cookies.get("token"),
            },
            withCredentials: true,
          }
        )
        .then((response) => {
          this.channels.push(response.data);
          (this.channelName = ""),
            (this.channelDescription = ""),
            (this.dialog = false);
          this.selectChannel(response.data.Name);
        })
        .catch(() => {});
    },
    selectChannel(name) {
      //指定された名前のチャンネルを選択する
      this.channels.forEach((ch) => {
        if (name == ch.Name) {
          this.switchChannel(ch);
        }
      });
    },
  },
};
</script>

<style>
.pa-4 {
  position: fixed;
  width: 256px;
  color: #455a64;
  border-width: 0 0 1px 0;
  border-style: solid;
  border-color: #000000;
  z-index: 100;
}
.channels {
  overflow: auto;
  margin-top: 68px;
  border-width: 1px 0 0 0;
  border-style: solid;
  border-color: #000000;
}
.v-navigation-drawer__content::-webkit-scrollbar {
  display: none;
}
</style>