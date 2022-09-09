<template>
	<div>
		<div class="absolute navbar-transparent z-10 w-full shadow-sm">
			<!-- Navbar -->
			<div class="flex items-center justify-between mr-14 px-4 py-1 sm:p-0">
				<a href="/">
					<img
						class="img-logo ml-5"
						style="max-width: 50%; height: auto"
						src="/img/logo-gits.png"
					/>
				</a>
				<div class="sm:hidden">
					<button
						@click="isOpen = !isOpen"
						type="button"
						class="block text-gray-500 hover:text-black focus:text-black focus:outline-none"
					>
						<svg class="h-6 w-6 fill-current" viewBox="0 0 24 24">
							<path
								v-if="isOpen"
								fill-rule="evenodd"
								d="M18.278 16.864a1 1 0 0 1-1.414 1.414l-4.829-4.828-4.828 4.828a1 1 0 0 1-1.414-1.414l4.828-4.829-4.828-4.828a1 1 0 0 1 1.414-1.414l4.829 4.828 4.828-4.828a1 1 0 1 1 1.414 1.414l-4.828 4.829 4.828 4.828z"
							/>
							<path
								v-if="!isOpen"
								fill-rule="evenodd"
								d="M4 5h16a1 1 0 0 1 0 2H4a1 1 0 1 1 0-2zm0 6h16a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2zm0 6h16a1 1 0 0 1 0 2H4a1 1 0 0 1 0-2z"
							/>
						</svg>
					</button>
				</div>

				<nav
					:class="isOpen ? 'block' : 'hidden'"
					class="px-2 mr-4 sm:flex sm:p-0"
				>
					<a
						href="/"
						class="dropdown navlist relative inline-block block flex items-center justify-center p-3 text-black font-lato font-semibold rounded mx-3"
					>
						<div class="text-black">Beranda</div>
					</a>
					<a
						href="/dataset"
						class="dropdown navlist relative inline-block block flex items-center justify-center p-3 text-black font-lato font-semibold rounded mx-3"
					>
						<div class="text-black">Dataset</div>
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
	background: #7ba0ff;
	border-bottom: 1px solid rgba(255, 255, 255, 1);
}

.navlist {
	transition: background 0.2s ease-in-out;
}

.bg-header {
	background: #c2d3ff;
	height: 57px;
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
		};
	},
});
</script>
