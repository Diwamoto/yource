<template>
  <div id="app">
    <div id="header">
      <!--リンクタグを生成します。-->
      <router-link to="/">top</router-link>
      <router-link to="/test2">about</router-link>
      <router-link to="/test3">contact</router-link>
    </div>
    <p>test1</p>
    <!--上記のリンクタグで設定したコンポーネントが表示される場所です。-->
    <router-view>
    </router-view>
  </div>
</template>

<!--コンポーネントの名前を定義します。-->
<script>
export default {
  name: 'App'
}
</script>

<!--スタイルの指定をします-->
<style>
body {
  margin: 0;
}

#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#header {
  height: 40px;
  background: white;
  box-shadow: 0px 3px 3px rgba(0,0,0,0.1);
  display: flex;
  justify-content: center;
  align-items: center;
}

#header a {
  text-decoration: none;
  color: #2c3e50;
  margin: 0 10px;
  padding: 3px 10px;
  background: #5ccebf;
}
</style>
