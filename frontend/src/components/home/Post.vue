<template>
    <div class="posts-wrapper" ref="posts">
        <div class="welcome">
            <h2>#{{ channel.Name }} へようこそ！</h2>
            <h5>説明: {{ channel.Description }}</h5>
        </div>
        <div v-if="posts.length > 0" class="posts">
            <div class="post" v-for="post in posts" :key="post.id">
                <v-row>
                    <v-col cols="12">
                        <v-avatar
                        color="primary"
                        size="50"
                        tile
                        dark
                        class="rounded"
                        ><span style="color:white">{{ getAvatarPrefix(post.User.Name) }}</span></v-avatar>
                    <span class="post-user-name">{{ post.User.Name }}</span><small class="post-date">{{ formatPostDate(post.Date) }}</small>
                    </v-col>
                    <v-col cols="12" class="content">
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
        }
    }

}
</script>

<style>

</style>