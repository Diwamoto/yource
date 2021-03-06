<template>
  <div class="menu d-flex justify-space-around">
      <div class="links" v-for="link in links" :key=link.name>
          <v-btn
            color="white"
            rounded
            large
            width="180"
            :href="link.link"
          >{{ link.name }}</v-btn>
      </div>
  </div>
</template>

<script>
export default {
    name: "Menu",
    data() {
        return{
            links: [
                {
                    name: "概要",
                    link: "/about",
                },
                {
                    name: "プライバシーポリシー",
                    link: "/privacy",
                },
                {
                    name: "サイト運営者について",
                    link: "/history",
                },

            ]
        }
    }
}
</script>

<style>
.menu{
    margin: 9% 30% 0 30%;
}
</style>