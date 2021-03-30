<template>
  <div class="maintenance">
    <div class="maintenance-box rounded-lg">
      <h1 class="maintenance-title">メンテナンス中</h1>
      <div class="maintenance-body-wrapper">
        <h5 class="maintenance-main-title">Yourceは現在メンテナンス中です。ご迷惑をおかけして申し訳ありません。</h5>
        <div class="maintenance-body maintenance-img-wrapper">
          <v-img class="maintenance-img" src="/img/maintenance.svg"></v-img>
        </div>
        <div style="margin-top:20px;text-align:center;">お問い合わせ先: <a href="https://diwamoto.dev#contact" target="_brank">diwamoto.dev</a></div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "maintenance",
};
</script>

<style>
.maintenance {
  height: 100%;
  background-color: #f3f3f3;
}
.maintenance-box {
  background-color: white;
  margin: 15vh auto 0 auto;
  padding-bottom: 10vh;
  width: 50vw;
  opacity: 0;
  box-shadow: 0 6px 6px -3px rgba(0, 0, 0, 0.2),
    0 10px 14px 1px rgba(0, 0, 0, 0.14), 0 4px 18px 3px rgba(0, 0, 0, 0.12) !important;
  animation-name: maintenance;
  animation-duration: 0.5s;
  animation-timing-function: ease-out;
  animation-fill-mode: forwards;
}
.maintenance-title {
  padding: 2% 0 2% 0;
  text-align: center;
  border-width: 0 0 2px 0;
  border-style: solid;
  border-color: #f0f0f0;
}
.maintenance-main-title {
  font-size: 1.2em;
  padding: 2% 0 2% 0;
  text-align: center;
}
.maintenance-img-wrapper {
  margin: 0 2vw 0 2vw;
}
.maintenance-img {
  margin: 0 5vw 0 5vw;
  width: 70%;
}
.maintenance-subtitle {
  padding: 3% 10% 0 10%;
}
.maintenance-body-wrapper {
  padding: 6% 10% 0 10%;
}
.maintenance-body {
  padding: 0 5% 0 5%;
}
@keyframes maintenance {
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