<template>
	<div class="font-lato mt5 mb-5 w-full">
		<!-- Filter Result -->
		<div class="filter-border px-3 py-2">
			<div class="flex justify-between">
				<span class="color-navyblue text-xl font-semibold"
					>Filter By :</span
				>
				<button
					@click="resetFilter"
					style="color: #1976d2"
					class="font-bold hover:underline text-sm"
				>
					Reset Filter
				</button>
			</div>
			<p class="mt2 text-sm mb-0">Showing results based on categories</p>
		</div>
		<!-- Tahun -->
		<div class="filter-border px-3 py-2">
			<p class="mb3 color-navyblue font-semibold text-xl">Tahun Data</p>
			<div class="flex">
				<input
					v-model="tahunAwal"
					@change="onChange()"
					class="input-tahun w-16 px-2 rounded font-normal text-l ml-5 mr-5 border-4"
					type="text"
					placeholder="2015"
				/>
				<h1 class="text-xl font-bold pb-0">-</h1>
				<input
					v-model="tahunAkhir"
					@change="onChange"
					class="input-tahun w-16 px-2 rounded font-normal text-l ml-5 mr-5 border-4"
					type="text"
					placeholder="2022"
				/>
			</div>
		</div>

		<!-- Sumber Data -->
		<div class="filter-border px-3 py-2">
			<p class="mb2 color-navyblue font-semibold text-xl">
				Preferensi Sumber Data
			</p>
			<!-- With Radio Button Filter (Accomodation) -->
			<div v-for="(type, index) in data.sumberDataPreference" :key="index">
				<input
					type="radio"
					:value="data.sumberDataPreference[index]"
					v-model="selectedSumberData"
					@change="updateFilterData()"
				/>
				<label class="text-base color-darkgray">
					{{ data.sumberDataPreference[index] }}
				</label>
			</div>
		</div>
		<!-- Preferensi Topik -->
		<div class="filter-border px-3 py-2">
			<p class="mb2 color-navyblue font-semibold text-xl">
				Preferensi Topik
			</p>
			<!-- With Radio Button Filter (Accomodation) -->
			<div v-for="(type, index) in data.topikPreference" :key="index">
				<input
					type="radio"
					:value="data.topikPreference[index]"
					v-model="selectedTopik"
					@change="updateFilterData()"
				/>
				<label class="text-base color-darkgray">
					{{ data.topikPreference[index] }}
				</label>
			</div>
		</div>
	</div>
</template>

<style scoped>
.search-card {
	background: #00b4a7;
	border: 3px solid #f94c4c;
}

.input-tahun {
	/* background: #00b4a7; */
	border: 3px solid #080606;
}

input[type="range"] {
	width: 50%;
}

.price-range-input {
	border: 1px solid #eeeeee;
	box-sizing: border-box;
	border-radius: 25px;
}

.filter-border {
	border: 1px solid #eeeeee;
	box-sizing: border-box;
}
</style>

<script>
import Vue from "vue";

export default Vue.extend({
	name: "PortalFilterCard",
	props: ["data", "portal_name"],

	data() {
		return {
			selectedTopik: "All",
			selectedSumberData: "All",
			tahunAwal: null,
			tahunAkhir: null,
			isReset: false,
		};
	},

	async created() {
		await this.data;
		this.passFilterValue();
	},

	methods: {
		resetFilter() {
			// this.$router.push({
			// 	path: "/dataset",
			// });
			this.tahunAwal = null;
			this.tahunAkhir = null;
			this.passFilterValue();
		},
		updateFilterData() {
			this.isReset = false;
			this.passFilterValue();
			// console.log("abis ganti filter :");
			// console.log(this.selectedTopik);
			// console.log(this.selectedSumberData);
		},
		onChange() {
			this.updateFilterData();
			// console.log("abis di enter : ");
			// console.log(this.tahunAwal);
			// console.log(this.tahunAkhir);
		},
		// getSliderValue() {
		// 	if (document.getElementById("slider1")) {
		// 		let val1 = document.getElementById("slider1") as HTMLInputElement;
		// 		this.sliderValTemp = Number(val1.value);
		// 		this.sliderVal1 = 1000000 - Number(val1.value);
		// 	}
		// 	if (document.getElementById("slider2")) {
		// 		let val2 = document.getElementById("slider2") as HTMLInputElement;
		// 		this.sliderVal2 = Number(val2.value);
		// 	}
		// 	this.updateFilterData();
		// },
		passFilterValue() {
			let dataSelected = {};
			dataSelected.topikPrefer = this.selectedTopik;
			dataSelected.sumberDataPrefer = this.selectedSumberData;
			dataSelected.isReset = this.isReset;
			if (this.tahunAwal !== null) {
				dataSelected.tahunAwal = this.tahunAwal;
			}
			if (this.tahunAkhir !== null) {
				dataSelected.tahunAkhir = this.tahunAkhir;
			}

			this.$emit("passedFilterValue", dataSelected);
		},
	},
});
</script>
