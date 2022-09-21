<template>
	<div>
		<div class="sumber-data mt-14 ml-10 font-poppins font-semibold">
			<h1 class="text-left text-xl my-1 ml-20" style="color: #28a0f6">
				Sumber Data
			</h1>
			<h1 class="text-left text-4xl my-3 ml-20" style="color: #353d49">
				Dari mana dataset didapatkan?
			</h1>
			<div class="flex text-xl ml-16 mt-5 mb-10">
				<a
					href="https://bps.go.id/"
					target="_blank"
					class="bg-sumber-data1 sumber-data-size dark-blue rounded-md items-center justify-center text-center pt-5 mx-8"
					>BPS
				</a>
				<a
					href="https://bappenas.go.id/statistik"
					target="_blank"
					class="bg-sumber-data2 sumber-data-size dark-blue rounded-md items-center justify-center text-center pt-5 mx-8"
				>
					Bappenas
				</a>
				<a
					href="https://opendata.jabarprov.go.id/id/dataset"
					target="_blank"
					class="bg-sumber-data1 sumber-data-size dark-blue rounded-md items-center justify-center text-center pt-2 px-3 mx-8"
					>Open Data Jabar
				</a>
				<a
					href="https://data.jakarta.go.id/group"
					target="_blank"
					class="bg-sumber-data2 sumber-data-size dark-blue rounded-md items-center justify-center text-center pt-2 px-3 mx-8"
				>
					Open Data Jakarta
				</a>
				<a
					href="http://data.bandung.go.id/index.php/portal/dataset"
					target="_blank"
					class="bg-sumber-data1 sumber-data-size dark-blue rounded-md items-center justify-center text-center pt-2 px-3 mx-8"
					>Open Data Bandung
				</a>
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
	/* background: #d7e3fc; */
	background: #ffffff;
	color: #353d49;
	box-shadow: 0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1);
}
.bg-sumber-data1:hover {
	text-decoration-line: underline;
}
.bg-sumber-data2 {
	/* background: #abc3ff; */
	background: #ffffff;
	color: #353d49;
	box-shadow: 0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1);
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
	/* background: #53b2d0; */
	background: #ffffff;

	box-shadow: 0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1);
}
/* .bg-topik1:hover {
	text-decoration-line: underline;
} */
.bg-topik2 {
	/* background: #94b8d0; */
	background: #ffffff;

	box-shadow: 0 4px 6px -1px rgb(0 0 0 / 0.1), 0 2px 4px -2px rgb(0 0 0 / 0.1);
}
/* .bg-topik2:hover {
	text-decoration-line: underline;
} */
</style>

<script>
import Vue from "vue";

export default Vue.extend({
	name: "DataSourceCard",
	data() {
		return {
			scrollPosition: 0,
		};
	},
});
</script>
