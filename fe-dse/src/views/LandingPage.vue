<template>
	<div>
		<div class="partnership font-sora">
			<Navbar :isTransparent="false" />
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
			<div class="sumber-data my-5 font-lato font-semibold">
				<h1 class="text-center text-3xl my-10">Sumber Data</h1>
				<div class="flex mx-30 text-xl items-center justify-center mt-5">
					<a
						href="https://bps.go.id/"
						target="_blank"
						class="bg-sumber-data1 sumber-data-size dark-blue rounded-md items-center justify-center text-center pt-5 mx-4"
						>BPS
					</a>
					<a
						href="bappenas.go.id/statistik"
						target="_blank"
						class="bg-sumber-data2 sumber-data-size dark-blue rounded-md items-center justify-center text-center pt-5 mx-4"
					>
						Bappenas
					</a>
					<a
						href="https://opendata.jabarprov.go.id/id/dataset"
						target="_blank"
						class="bg-sumber-data1 sumber-data-size dark-blue rounded-md items-center justify-center text-center pt-2 px-3 mx-4"
						>Open Data Jabar
					</a>
					<a
						href="https://data.jakarta.go.id/group"
						target="_blank"
						class="bg-sumber-data2 sumber-data-size dark-blue rounded-md items-center justify-center text-center pt-2 px-3 mx-4"
					>
						Open Data Jakarta
					</a>
					<a
						href="http://data.bandung.go.id/index.php/portal/dataset"
						target="_blank"
						class="bg-sumber-data1 sumber-data-size dark-blue rounded-md items-center justify-center text-center pt-2 px-3 mx-4"
						>Open Data Bandung
					</a>
				</div>
			</div>
			<div class="sumber-data my-5 font-lato font-semibold">
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
			<footer
				class="bottom-0 bg-navyblue mt-20"
				style="height: 300px; background: #013d6b"
			></footer>
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
// import axios from "axios";
import Navbar from "../components/Navbar.vue";

export default Vue.extend({
	components: {
		Navbar,
	},
	name: "Tes",
	data() {
		return {
			isHidden: false,
		};
	},
});
</script>
