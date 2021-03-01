<template>
  <div class="profile rounded-lg" elevation="12">
    <v-form @submit.prevent>
      <v-text-field
        v-model="profile_field"
        label="自己紹介"
        required
        outlined
        dense
      ></v-text-field>
      <v-text-field
        v-model="Birthday"
        label="誕生日"
        required
        outlined
        dense
      ></v-text-field>
      <v-text-field
        v-model="Hometown"
        label="出身地"
        required
        outlined
        dense
      ></v-text-field>
      <v-text-field
        v-model="Job"
        label="職業"
        required
        outlined
        dense
      ></v-text-field>
      <v-text-field
        v-model="Twitter"
        label="ツイッターURL"
        required
        outlined
        dense
      ></v-text-field>
      <v-text-field
        v-model="Facebook"
        label="フェイスブックURL"
        required
        outlined
        dense
      ></v-text-field>
      <v-text-field
        v-model="Instagram"
        label="インスタグラムURL"
        required
        outlined
        dense
      ></v-text-field>
      <v-text-field
        v-model="Other"
        label="その他webサイトURL"
        required
        outlined
        dense
      ></v-text-field>
      <v-btn class="mr-4" type="submit" > 変更する </v-btn>
    </v-form>
  </div>
</template>

<script>
export default {
  name: "Profile",
  props: {
    user: [Object, Array],
  },
  data() {
    return {
      profile: null,
      profile_field: "",
      Birthday: null,
      Hometown: "",
      Job: "",
      Twitter: "",
      Facebook: "",
      Instagram: "",
      Other: "",
    };
  },
  mounted() {
    this.profile = this.user.Profile;
  },
  methods: {
      submit () {
        this.$http
      },
  }
};
</script>

<style>
.profile {
  margin: 10vh 25vw 0 25vw;
  width: 50vw;
  background-color: white;
  box-shadow: 0 6px 6px -3px rgba(0, 0, 0, 0.2),
    0 10px 14px 1px rgba(0, 0, 0, 0.14), 0 4px 18px 3px rgba(0, 0, 0, 0.12) !important;
  opacity: 0;
  animation-name: profile;
  animation-duration: 0.5s;
  animation-timing-function: ease-out;
  animation-fill-mode: forwards;
}
@keyframes profile {
  0% {
    opacity: 0;
    transform: translateY(20px);
  }
  100% {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>