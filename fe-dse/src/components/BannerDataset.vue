<template>
	<div>
		<div
			class="banner mb-5 font-lato font-semibold text-5xl flex items-center justify-between"
		>
			<div class="ml-16 mr-4">
				<h1 class="my-4">Temukan</h1>
				<h1 class="my-4">Dataset yang kamu inginkan</h1>
				<h1 class="my-4">dengan mudah dan cepat!</h1>
				<!-- <h1 class="my-4">Topik : <span class="font-semibold">All</span></h1>
				<h1 class="my-4">
					Tahun : <span class="font-semibold">2015 - 2022</span>
				</h1> -->
			</div>

			<div class="ml-10">
				<img
					class="ml-5"
					style="max-width: 80%; height: auto"
					src="/img/filter-data.png"
				/>
			</div>
		</div>
	</div>
</template>

<style scoped>
.banner {
	/* background: #abc3ff; */
	background: #28a0f6;
	color: #ffffff;
	height: 300px;
	/* color: #013d6b; */
}
</style>

<script>
import Vue from "vue";
// import axios from "axios";

export default Vue.extend({
	name: "BannerDataset",
	data() {
		return {
			scrollPosition: 0,
		};
	},
});
</script>
