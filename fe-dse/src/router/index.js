import Vue from "vue";
import VueRouter from "vue-router";
import LandingPage from "../views/LandingPage.vue";
import DatasetPage from "../views/DatasetPage.vue";

Vue.use(VueRouter);

const routes = [
	{
		path: "/",
		name: "LandingPage",
		component: LandingPage,
	},
	{
		path: "/dataset",
		name: "DatasetPage",
		component: DatasetPage,
	},
];

const router = new VueRouter({
	mode: "history",
	base: process.env.BASE_URL,
	routes,
	props: true,
});

export default router;
