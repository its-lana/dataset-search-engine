<template>
	<div>
		<div class="partnership font-sora">
			<Navbar :isTransparent="false" />
			<BannerLandingPage />
			<DataSourceCard />
			<div class="item-card w-auto shadow-none mx-32 mt-24"></div>
			<TopicGroupCard />
			<Footer />
			<!-- <footer
				class="bottom-0 bg-navyblue mt-20"
				style="height: 300px; background: #013d6b"
			></footer> -->
		</div>
	</div>
</template>

<style scoped>
.item-card {
	border: 1px solid #eeeeee;
	box-sizing: border-box;
	border-radius: 16px;
}
.banner {
	background: #abc3ff;
}
.dark-blue {
	color: #013d6b;
}
.color-text-mulai {
	color: rgb(0, 0, 0);
}
.sumber-data-size {
	width: 150px;
	height: 70px;
}
.topik-size {
	height: 50px;
}
.bg-sumber-data1 {
	background: #d7e3fc;
}
.bg-sumber-data1:hover {
	text-decoration-line: underline;
}
.bg-sumber-data2 {
	background: #abc3ff;
}
.bg-sumber-data2:hover {
	text-decoration-line: underline;
}
.bg-mulai {
	background: #ffee93;
	border-style: solid;
	border-color: #abc3ff;
	/* color: #013d6b; */
	/* background: #5376ce; */
}
.bg-mulai:hover {
	background: #ffda07;
	/* color: white; */
	/* background: #5376ce; */
}
.bg-topik1 {
	background: #53b2d0;
}
/* .bg-topik1:hover {
	text-decoration-line: underline;
} */
.bg-topik2 {
	background: #94b8d0;
}
/* .bg-topik2:hover {
	text-decoration-line: underline;
} */
</style>

<script>
import Vue from "vue";
// import axios from "axios";
import Navbar from "../components/Navbar.vue";
import BannerLandingPage from "../components/BannerLandingPage.vue";
import DataSourceCard from "../components/DataSourceCard.vue";
import TopicGroupCard from "../components/TopicGroupCard.vue";
import Footer from "../components/Footer.vue";

export default Vue.extend({
	components: {
		Navbar,
		BannerLandingPage,
		DataSourceCard,
		TopicGroupCard,
		Footer,
	},
	name: "LandingPage",
	data() {
		return {
			isHidden: false,
		};
	},
});
</script>
