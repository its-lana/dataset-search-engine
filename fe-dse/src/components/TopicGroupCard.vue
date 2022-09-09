<template>
	<div>
		<div class="sumber-data mt-5 mb-10 font-lato font-semibold">
			<h1 class="text-center text-3xl mt-16 mb-10">Topik</h1>
			<div class="grid grid-cols-2 gap-4 mx-32 text-xl">
				<h1
					class="bg-topik1 topik-size dark-blue rounded-md items-center justify-center text-center pt-2 px-3 mx-1"
				>
					Ekonomi dan Industri
				</h1>
				<h1
					class="bg-topik2 topik-size dark-blue rounded-md items-center justify-center text-center pt-2 px-3 mx-1"
				>
					Infrastruktur
				</h1>
				<h1
					class="bg-topik2 topik-size dark-blue rounded-md items-center justify-center text-center pt-2 px-3 mx-1"
				>
					Sosial dan Kependudukan
				</h1>
				<h1
					class="bg-topik1 topik-size dark-blue rounded-md items-center justify-center text-center pt-2 px-3 mx-1"
				>
					Kesehatan
				</h1>
				<h1
					class="bg-topik1 topik-size dark-blue rounded-md items-center justify-center text-center pt-2 px-3 mx-1"
				>
					Pemuda dan Olahraga
				</h1>
				<h1
					class="bg-topik2 topik-size dark-blue rounded-md items-center justify-center text-center pt-2 px-3 mx-1"
				>
					Wisata dan Budaya
				</h1>
				<h1
					class="bg-topik2 topik-size dark-blue rounded-md items-center justify-center text-center pt-2 px-3 mx-1"
				>
					Pendidikan
				</h1>
				<h1
					class="bg-topik1 topik-size dark-blue rounded-md items-center justify-center text-center pt-2 px-3 mx-1"
				>
					Politik dan Pemerintahan
				</h1>
				<h1
					class="bg-topik1 topik-size dark-blue rounded-md items-center justify-center text-center pt-2 px-3 mx-1"
				>
					Komunikasi dan Informasi
				</h1>
				<h1
					class="bg-topik2 topik-size dark-blue rounded-md items-center justify-center text-center pt-2 px-3 mx-1"
				>
					Lingkungan Hidup dan SDA
				</h1>
				<h1
					class="bg-topik2 topik-size dark-blue rounded-md items-center justify-center text-center pt-2 px-3 mx-1"
				>
					Keagamaan
				</h1>
				<h1
					class="bg-topik1 topik-size dark-blue rounded-md items-center justify-center text-center pt-2 px-3 mx-1"
				>
					Pertahanan dan Keamanan
				</h1>
				<h1
					class="bg-sumber-data2 col-span-2 topik-size dark-blue rounded-md items-center justify-center text-center pt-2 px-3 mx-auto"
					style="width: 600px"
				>
					Perhubungan
				</h1>
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
/* .bg-sumber-data1 {
	background: #d7e3fc;
}
.bg-sumber-data1:hover {
	text-decoration-line: underline;
}

.bg-sumber-data2:hover {
	text-decoration-line: underline;
} */

.bg-sumber-data2 {
	background: #abc3ff;
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
	name: "TopicGroupCard",
	data() {
		return {
			scrollPosition: 0,
		};
	},
});
</script>
