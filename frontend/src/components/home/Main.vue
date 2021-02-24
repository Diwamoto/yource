<template>
  <div class="main">
    <v-main app>
      <Menu :channel="channel"></Menu>

      <Posts :channel="channel" :posts="posts" ref="Posts"></Posts>

      <PostForm :channel="channel" :user-id="userId" :get-posts="getPostsForChild" ref="Form"></PostForm>
      
    </v-main>
  </div>
</template>

<script>
import Menu from '@/components/home/Menu.vue'
import Posts from '@/components/home/Post.vue'
import PostForm from '@/components/home/Form.vue'
export default {
  name: "Main",
  components: {
    Menu,
    Posts,
    PostForm
  },
  props: {
    channel: [Object, Array],
    posts: [Object, Array],
    userId: Number,
    getPosts: {
      type: Function,
      required: true,
    }
  },
  data() {
    return {
      dialog: false,
      postNotFound: false,
    }
  },
  updated() {
    this.$refs.Form.updateLabel()
  },
  methods: {
    scrollBottom(){
      this.$refs.Posts.scrollBottom()
    },
    getPostsForChild(){
      this.getPosts()
    }
  }
}
</script>

<style>
.posts-wrapper{
  position: fixed;
  height: 77.6%;
  width: 86.6%;
  margin-top: 68px;
  padding-top: 10px;
  overflow:auto;
}
.post{
  margin-left: 15px;
  height: 80px;
  width: 80%;
}
.welcome{
  margin-left: 30px;
  width: 86.6%;
}
.input-form{
  position: fixed;
  margin-top: 43%;
  width: 87%;
  padding-top: 10px;
  padding-left: 10px;
  border-width: 1px 0 0 0;
	border-style: solid;
	border-color: #909090;
}
.input-box{
  width: 98%;
}
.v-avatar{
  margin-top: 30px;
}
.content{
  margin-top: -3.7%;
  margin-left: 58px;
}
.post-user-name{
  margin-left: 9px;
  font-weight:bold;
}
.post-date{
  margin-left: 10px
}
</style>