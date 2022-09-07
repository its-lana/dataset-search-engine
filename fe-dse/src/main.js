import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import vuetify from "./plugins/vuetify";
import "@/index.css";
// import axios from 'axios'
import "tw-elements";

// axios.defaults.baseURL = 'https://web-jtk-2022.herokuapp.com/';

Vue.config.productionTip = false;

new Vue({
	router,
	vuetify,
	render: (h) => h(App),
}).$mount("#app");
