import { createApp } from "vue";

import App from "./App.vue";
import router from "./router";
import  "../src/axios/axios";


import "./assets/css/bootstrap.css";
import "./assets/js/bootstrap.js";

const app = createApp(App);

app.use(router);

app.mount("#app");
