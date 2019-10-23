import Vue from "vue";
import App from "./App.vue";
import router from "./router";

Vue.config.productionTip = false;
Vue.config.devtools = true;

import * as Wails from "@wailsapp/runtime";

Wails.Init(() => {
  new Vue({
    router,
    render: h => h(App)
  }).$mount("#app");
});
