<template>
	<div>
		<div class="absolute navbar-transparent z-10 w-full shadow">
			<!-- Navbar -->
			<div class="flex items-center justify-between mr-14 px-4 py-4 sm:p-0">
				<a href="/">
					<img
						class="img-logo ml-16"
						style="max-width: 50%; height: auto"
						src="/img/logo-gits.png"
					/>
				</a>

				<nav class="px-2 mr-10 flex p-0">
					<a
						href="/"
						class="dropdown navlist relative inline-block block flex items-center justify-center p-3 font-lato font-semibold rounded text-lg mx-4"
						:class="{ underline: !isDatasetPage }"
					>
						<div class="">Beranda</div>
					</a>
					<a
						href="/dataset"
						class="dropdown navlist relative inline-block block flex items-center justify-center p-3 font-lato font-semibold text-lg rounded mx-4"
						:class="{ underline: isDatasetPage }"
					>
						<div class="">Dataset</div>
					</a>
				</nav>
			</div>
		</div>
		<div class="bg-header"></div>
	</div>
</template>

<style scoped>
.navlist:hover {
	/* background: rgba(2, 75, 117, 0.4); */
	/* background: #7ba0ff; */
	background: #28a0f6;
	color: #ffffff;
	border-bottom: 1px solid rgba(255, 255, 255, 1);
}

.underline {
	/* text-decoration-color: ; */
	text-decoration-line: none;
	border-bottom: 2px solid #28a0f6;
}

.navlist {
	transition: background 0.2s ease-in-out;
	color: #353d49;
}

.bg-header {
	background: #ffffff;
	height: 90px;
	color: black;
}

.dropdown-content a:hover {
	background: rgba(0, 0, 0, 0.1);
	border-bottom: 1px solid rgba(0, 0, 0, 0.5);
}

.dropdown-content a {
	color: black;
	padding: 12px 16px;
	text-decoration: none;
	display: block;
	transition: background 0.2s ease-in-out;
}

.dropdown-content {
	background: #f1f1f1;
	top: 100%;
	left: 0;
	min-width: 160px;
	box-shadow: 0px 8px 16px 0px rgba(0, 0, 0, 0.2);
	border-radius: 4px;
	z-index: 1;
}

.dropdown:hover .dropdown-content {
	display: block !important;
}

.navbar-transparent {
	background: rgba(255, 255, 255, 0);
}

.img-logo {
	height: auto;
	width: auto;
	max-height: 72px;
	max-width: 250px;
}

.search-container {
	height: 260px;
	background: #142e46;
}
.email-subscribe-btn {
	width: 80px;
	height: 38px;
	background: #8d5cf6;
	border: 0.3px solid #142e46;
	box-sizing: border-box;
	border-radius: 5px;
	transition: background 0.1s ease-in-out;
}

.email-subscribe-btn:hover {
	background: #8d5cf6;
}

.email-input-box {
	width: 250px;
	height: 38px;
	background: #f7f7f7;
	border: 1px solid #142e46;
	box-sizing: border-box;
	border-radius: 5px;
}
</style>

<script>
import Vue from "vue";
// import axios from "axios";

export default Vue.extend({
	name: "Navbar",
	props: ["isTransparent"],

	data() {
		return {
			isOpen: false,
			state: "close",
			scrollPosition: 0,
			article: [],
			isDatasetPage: false,
		};
	},
	created() {
		if (this.$route.name == "DatasetPage") {
			this.isDatasetPage = true;
		} else {
			this.isDatasetPage = false;
		}
	},
});
</script>
