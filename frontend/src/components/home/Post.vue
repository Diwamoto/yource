<template>
    <div class="posts-wrapper" ref="posts">
        <div class="welcome" v-bind:style="marginTop">
            <h2>#{{ channel.Name }} へようこそ！</h2>
            <h5>説明: {{ channel.Description }}</h5>
        </div>
        <div v-if="posts.length > 0" class="posts">
            <div class="posts-item" v-for="post in posts" :key="post.id">
                <v-row>
                    <v-col cols="12">
                        <v-avatar
                        color="primary"
                        size="50"
                        tile
                        dark
                        class="rounded"
                        ><span style="color:white">{{ getAvatarPrefix(post.User.Name) }}</span></v-avatar>
                    <span class="posts-item-user-name">{{ post.User.Name }}</span><small class="posts-item-date">{{ formatPostDate(post.Date) }}</small>
                    </v-col>
                    <v-col cols="12" class="posts-item-content">
                        {{ post.Content }}
                    </v-col>
                </v-row>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    Name: "Post",
    props: {
        channel: [Object, Array],
        posts: [Object, Array],
    },
    data() {
        return {
            marginTop: {
                "margin-top" : "650px"
            }
        }
    },
    mounted(){

        //一番下までスクロールする
        //DOM要素にはmounted以降でしか動かせないためmountedに
        this.scrollBottom()
    },
    updated() {
    
        //最下部までスクロールする
        this.scrollBottom()

        //マージンを動的に設定する。
        this.dynamicMargin()
    },
    methods: {
        
        //一番下までスクロールする
        scrollBottom(){
            this.$refs.posts.scrollTop = this.$refs.posts.scrollHeight
        },
        //日付をフォーマットする。今日の投稿であれば先頭に「今日」と言う文字と時刻を付け足す。
        formatPostDate(rawPostDate){
            //文字列からmoment形式に変換
            var postDate = this.$moment(new Date(rawPostDate))
            var nowDate = this.$moment()
            //日付を今日と比較する。
            if (nowDate.diff(postDate,"days") == 0){
                return "今日" + postDate.format("HH:mm")
            } else {
                return postDate.format("YYYY/MM/DD")
            }
        },
        //アイコンがないユーザに対して、そのユーザの名前の先頭二文字をとってきて使う。
        //もし英語であれば大文字にする。
        getAvatarPrefix(Name){
            return Name.slice( 0, 2 ).toUpperCase()
        },

        //マージンを動的に設定する。
        //投稿一つのサイズが50px(アバターに依存する)なので、一つ投稿があるごとにmarginを50px減らす。
        //投稿を表示する欄のサイズが11個表示するのがちょうどいいので11個以上あれば0にして、それ以外は個数*50px減らす。
        //投稿がなければ元の数値に戻す。
        dynamicMargin(){
            if (this.posts.length > 11) {
                this.marginTop["margin-top"] = "0px"
            }else if (this.posts.length > 0) {
                this.marginTop["margin-top"] = (600 - 50 * this.posts.length) + "px"
            }else if (this.posts.length == 0) {
                this.marginTop["margin-top"] = "650px"
            }
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
.posts-item{
  margin-left: 15px;
  height: 80px;
  width: 80%;
}
.welcome{
  margin-left: 30px;
  width: 86.6%;
}
.posts-item-user-name{
  margin-left: 9px;
  font-weight:bold;
}
.posts-item-date{
  margin-left: 10px
}
.v-avatar{
  margin-top: 30px;
}
.posts-item-content{
  margin-top: -3.7%;
  margin-left: 58px;
}
</style>