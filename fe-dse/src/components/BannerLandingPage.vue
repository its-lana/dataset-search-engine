<template>
	<div>
		<div
			class="grid grid-cols-3 gap-0 mb-5 h-56 font-lato font-semibold text-3xl"
		>
			<div class="pl-10 banner pt-6">
				<img
					class="pl-5"
					style="max-width: 60%; height: auto"
					src="/img/statistik.png"
				/>
				<!-- <h1 class="my-4">Tahun : 2021 - 2022</h1> -->
			</div>
			<div
				class="col-span-2 pl-8 pr-4 pt-10 text-center items-center justify-center"
				style="background: #d7e3fc"
			>
				<h1 class="mb-4 dark-blue">
					Selamat Datang di Website Pencarian Dataset!
				</h1>
				<h1 class="my-4 dark-blue">
					Sekarang, pencarian data menjadi lebih mudah dan cepat
				</h1>
				<div
					class="mx-auto bg-mulai border-4 pt-1 text-center items-center justify-center"
					style="width: 400px; height: 50px"
				>
					<a href="/dataset" class="color-text-mulai"
						>Mulai Pencarian Dataset >></a
					>
				</div>
			</div>
		</div>
	</div>
</template>

<style scoped>
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

export default Vue.extend({
	name: "BannerLandingPage",
	data() {
		return {
			scrollPosition: 0,
		};
	},
});
</script>
