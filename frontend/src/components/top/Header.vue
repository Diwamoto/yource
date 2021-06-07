<template>
  <div v-bind:class="[header, header_scroll]">
    <v-row>
      <!-- logo -->
      <v-col cols="6" md="2">
        <router-link to="/">
          <v-img
            v-bind:class="[logo, logo_scroll]"
            class="main-logo_image"
            src="/img/logo_transparent.png"
          ></v-img>
        </router-link>
      </v-col>

      <!-- logo -->

      <!-- menu -->
      <v-col md="1" offset="7" v-bind:class="[me, me_scroll]">
        <router-link to="/about" style="color: black"> 概要 </router-link>
      </v-col>

      <v-col cols="1" v-bind:class="[mej, mej_scroll]">
        <router-link to="/login">
          <v-btn color="red" elevation="5" large outlined rounded
            >ログイン</v-btn
          >
        </router-link>
      </v-col>

      <v-col
        cols="1"
        md="1"
        offset="4"
        offset-sm="4"
        offset-md="0"
        v-bind:class="[meh, meh_scroll]"
      >
        <v-app-bar-nav-icon class="hamburger" @click="toggleSidebar"></v-app-bar-nav-icon>
      </v-col>
    </v-row>
  </div>
</template>

<script>
export default {
  name: "Header",
  data: function () {
    return {
      header: "header",
      header_scroll: "",
      logo: "main-logo-image",
      logo_scroll: "",
      logo_width: "150",
      me: "menu-element d-none d-md-flex d-xl-flex",
      mej: "menu-element-join d-none d-md-flex d-xl-flex",
      meh: "menu-element-hamburger",
      me_scroll: "",
      mej_scroll: "",
      meh_scroll: "",
    };
  },
  mounted() {
    window.addEventListener("scroll", this.handleScroll);
  },
  methods: {
    handleScroll() {
      if (window.scrollY > 15) {
        this.header_scroll = "header_scrolling";
        this.logo_scroll = "main-logo-image_scrolling";
        this.me_scroll = "menu-element_scrolling";
        this.mej_scroll = "menu-element-join_scrolling";
        this.meh_scroll = "menu-element-hamburger_scrolling";
        this.logo_width = "110";
      } else {
        this.header_scroll = "";
        this.logo_scroll = "";
        this.me_scroll = "";
        this.mej_scroll = "";
        this.meh_scroll = "";
        this.logo_width = "150";
      }
    },
    toggleSidebar(){
      this.$parent.$refs.Sidebar.toggle();
    }
  },
};
</script>

<style>
.header {
  position: fixed;
  top: 0;
  width: 100%;
  z-index: 5;
  background-color: white;
  height: 80px;
  transition: all 0.2s ease-in-out;
}
.header_scrolling {
  height: 50px;
  box-shadow: 0px 8px 22px -5px rgba(0, 0, 0, 0.2);
  transition: all 0.2s ease-in-out;
}
.main-logo-image {
  width: 130px;
  margin-top: 3%;
  margin-left: 45%;
  transition: all 0.2s ease-in-out;
}
.main-logo-image_scrolling {
  width: 80px;
  margin-top: 8px;
  margin-left: 45%;
  transition: all 0.2s ease-in-out;
}
.menu-element {
  margin-left: -1%;
  margin-top: 30px;
  transition: all 0.2s ease-in-out;
}
.menu-element-join {
  margin-left: -1%;
  margin-top: 20px;
  transition: all 0.2s ease-in-out;
}
.menu-element-hamburger {
  margin-left: 19px;
  margin-top: 23px;
  transition: all 0.2s ease-in-out;
}
.menu-element a {
  color: black;
  text-decoration: none;
}
.menu-element-join a {
  color: black;
  text-decoration: none;
}
.menu-element_scrolling {
  margin-top: 15px;
  transition: all 0.2s ease-in-out;
}
.menu-element-join_scrolling {
  margin-top: 2px;
  transition: all 0.2s ease-in-out;
}
.menu-element-hamburger_scrolling {
  margin-top: 7px;
  transition: all 0.2s ease-in-out;
}
.btn-join {
  margin-top: 30px;
  color: white;
}
/* スマホ用css */
@media screen and (max-width: 480px) {
  .header {
    height: 80px;
  }
  .header_scrolling {
    height: 60px;
  }
  .main-logo-image {
    width: 100px;
    margin-top: 7%;
    margin-left: 20%;
    transition: all 0.2s ease-in-out;
  }
  .main-logo-image_scrolling {
    width: 100px;
    margin-top: 8px;
    margin-left: 20%;
    transition: all 0.2s ease-in-out;
  }
  .menu-element-hamburger {
    margin-left: 19px;
    margin-top: 20px;
    transition: all 0.2s ease-in-out;
  }
  .menu-element-hamburger_scrolling {
    margin-top: 16px;
    transition: all 0.2s ease-in-out;
  }
}
</style>