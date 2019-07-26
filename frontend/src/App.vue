<template>
  <div id="app">
    <div class="row m-3">
      <div class="mainMenu shadow-sm m-0 p-0 mt-3 mr-3 col-3 bg-white">
        <MainMenu/>
      </div>
      <div class="shadow-sm m-0 p-3 mt-3 col bg-white">
        <router-view/>
      </div>
    </div>
    <div class="mb-5">

    </div>
  </div>
</template>

<script>
import MainMenu from "./components/MainMenu.vue";
import "./assets/css/main.css";
import "./assets/css/bootstrap-4.3.1/bootstrap.min.css";

export default {
  name: "app",
  components: {
    MainMenu
  }
};
</script>
