<template>
  <div class="main">
    <v-main app>
      <Menu :channel="channel"></Menu>

      <Posts :channel="channel" :posts="posts" ref="Posts"></Posts>

      <PostForm
        :channel="channel"
        :user-id="userId"
        :get-posts="getPostsForChild"
        ref="Form"
      ></PostForm>
    </v-main>
  </div>
</template>

<script>
import Menu from "@/components/home/Menu.vue";
import Posts from "@/components/home/Post.vue";
import PostForm from "@/components/home/Form.vue";
export default {
  name: "Main",
  components: {
    Menu,
    Posts,
    PostForm,
  },
  props: {
    channel: [Object, Array],
    posts: [Object, Array],
    userId: Number,
    getPosts: {
      type: Function,
      required: true,
    },
  },
  data() {
    return {
      dialog: false,
      postNotFound: false,
    };
  },
  updated() {
    this.$refs.Form.updateLabel();
  },
  methods: {
    scrollBottom() {
      this.$refs.Posts.scrollBottom();
    },
    getPostsForChild() {
      this.getPosts();
    },
  },
};
</script>

<style>
</style>