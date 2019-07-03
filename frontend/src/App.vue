<template>
  <div id="app">
    <div class="row">
      <div class="col-3">
        <MainMenu/>
      </div>
      <div class="col">
        <router-view/>
      </div>
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
