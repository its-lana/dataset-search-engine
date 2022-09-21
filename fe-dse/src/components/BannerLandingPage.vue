<template>
	<div>
		<div
			class="banner grid grid-cols-3 gap-0 mb-5 h-80 font-lato font-semibold text-4xl"
		>
			<div
				class="col-span-2 pl-8 pr-4 text-center my-auto items-center justify-center"
			>
				<h1 class="mb-6">Selamat Datang di Website Pencarian Dataset!</h1>
				<h1 class="my-6">Sekarang, pencarian data menjadi lebih cepat</h1>
				<div
					class="mx-auto bg-mulai border-4 pt-1 text-center rounded-lg items-center justify-center"
					style="width: 500px; height: 60px"
				>
					<a href="/dataset" class="color-text-mulai"
						>Mulai Pencarian Dataset >></a
					>
				</div>
			</div>
			<div class="pr-10 pt-6">
				<img class="pr-5" src="/img/statistik.jpg" />
				<!-- <h1 class="my-4">Tahun : 2021 - 2022</h1> -->
			</div>
		</div>
	</div>
</template>

<style scoped>
.banner {
	/* background: #abc3ff; */
	background: #28a0f6;
	/* background: #ffffff; */
	/* color: rgb(0, 0, 0); */
	color: #ffffff;
}
.dark-blue {
	color: #013d6b;
}
.color-text-mulai {
	/* color: rgb(0, 0, 0); */
	color: #ffffff;
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
	/* background: #ffee93; */
	border-style: solid;
	/* border-color: #abc3ff; */
	/* color: #013d6b; */
	background: #ff8b49;
	/* border-color: #ff8b49; */
}
.bg-mulai:hover {
	background: #f37c38;
	/* border-color: #f37c38; */
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

export default Vue.extend({
	name: "BannerLandingPage",
	data() {
		return {
			scrollPosition: 0,
		};
	},
});
</script>
