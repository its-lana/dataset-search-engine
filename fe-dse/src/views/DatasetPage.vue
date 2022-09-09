<template>
	<div>
		<div class="partnership font-sora">
			<Navbar :isTransparent="false" />
			<BannerDataset />
			<div class="flex ml3 mr3 mt2">
				<div id="filter-section" class="mr2">
					<SearchCard
						formTitle="Dataset Name"
						@passedSearchValue="getSearchValue"
					/>
					<!-- Filter Butoon -->
					<div
						style="width: 90vw; text-align: center"
						class="bg-navyblue rounded ml-2 py-3 my-6 hidden-sm-and-up"
						@click="isHidden = true"
					>
						<button class="font-normal text-white text-l">Filter</button>
					</div>
					<PortalFilterCard
						v-if="doSearch"
						:data="datasetFilter"
						portal_name="Dataset"
						@passedFilterValue="getFilterData"
						class="m-0"
					/>
				</div>
				<div class="w-full">
					<div class="flex justify-between">
						<p class="text-black font-medium text-2xl pb-0 my-auto">
							Data Found :
							<span class="font-bold">{{ totalDatasets }}</span> datasets
						</p>
						<div class="flex mr-10 mb-3">
							<p
								class="text-black font-medium text-xl mx-4 pb-0 my-auto"
							>
								Sort By :
							</p>
							<!-- <div
								class="border-2 bg-sorting my-auto rounded-md px-3 py-1"
							> -->
							<select
								v-model="sortBy"
								@change="onSortingChange()"
								class="font-lato font-medium text-xl border-2 bg-sorting my-auto rounded-md px-3 py-1"
							>
								<option value="judul.asc">Judul[A-Z]</option>
								<option value="judul.desc">Judul[Z-A]</option>
								<option value="last_updated.desc">Last Updated</option>
							</select>
							<!-- <ion-icon name="chevron-down"></ion-icon> -->
							<!-- </div> -->
						</div>
					</div>
					<div class="item-card w-auto shadow-none mx-2 mb-7"></div>
					<loading
						:active.sync="isLoading"
						:can-cancel="true"
						:on-cancel="onCancel"
						:is-full-page="fullPage"
					></loading>
					<ItemPortalCard v-if="pageData" :data="pageData" />
				</div>
			</div>
			<Pagination
				@passedPageNumber="getPageNumber"
				:totalPages="totalPages"
				pageRange="5"
			/>
			<Footer />
		</div>
	</div>
</template>

<style scoped>
.banner {
	background: #abc3ff;
	color: #013d6b;
}
.bg-sorting {
	/* background: #abc3ff; */
	border-style: solid;
	border-color: #0f0f0f;
}
.bg-sorting:active {
	/* background: #abc3ff; */
	/* border-style: solid; */
	/* border-color: #626161; */
	border-radius: 0.375rem;
}
#filter-section {
	max-width: 24.5%;
}
.item-card {
	border: 1px solid #eeeeee;
	box-sizing: border-box;
	border-radius: 16px;
}
</style>

<script>
import Vue from "vue";
import axios from "axios";
import Navbar from "../components/Navbar.vue";
import BannerDataset from "../components/BannerDataset.vue";
import SearchCard from "../components/SearchCard.vue";
import PortalFilterCard from "../components/PortalFilterCard.vue";
import ItemPortalCard from "../components/ItemPortalCard.vue";
import Pagination from "../components/Pagination.vue";
import filter from "../../public/jsonData/datasetFilter.json";
// import dummyDataset from "../../public/jsonData/dummyDataset1.json";
import Loading from "vue-loading-overlay";
import "vue-loading-overlay/dist/vue-loading.css";
import Footer from "../components/Footer.vue";

export default Vue.extend({
	components: {
		Navbar,
		BannerDataset,
		SearchCard,
		PortalFilterCard,
		ItemPortalCard,
		Pagination,
		Loading,
		Footer,
	},
	name: "DatasetPage",
	data() {
		return {
			isHidden: false,
			isLoading: false,
			fullPage: true,
			datasetFilter: filter,
			BE_API: "http://localhost:3000/api",
			pageData: [{}],
			totalDatasets: null,
			totalPages: 100,
			searchValue: {},
			filterValue: {},
			sortBy: "judul.asc",
			pageNumber: 1,
			doSearch: false,
			prevQuery: null,
			currentQuery: null,
		};
	},
	async created() {
		try {
			this.currentQuery = this.getQuery();
			this.isLoading = true;
			// simulate AJAX
			setTimeout(() => {
				this.isLoading = false;
			}, 3000);
			await this.getData(this.currentQuery);
		} catch (error) {
			console.log(error);
		}
	},
	// async updated() {
	// 	try {
	// 		console.log("ada updaet url dari : " + this.prevRoute);
	// 		if (this.currentQuery !== this.prevQuery) {
	// 			console.log("masuk if apdet");
	// 			console.log(this.currentQuery + "===" + this.prevQuery);
	// 			this.prevQuery = this.currentQuery;
	// 			this.isLoading = true;
	// 			// simulate AJAX
	// 			setTimeout(() => {
	// 				this.isLoading = false;
	// 			}, 4000);
	// 			await this.getData(this.getQuery());
	// 		}
	// 	} catch (error) {
	// 		console.log(error);
	// 	}
	// },
	methods: {
		getQuery() {
			let uri = window.location.href.split("?");
			if (uri.length == 2) {
				return uri[1];
			}
			return "page=1&sortBy=judul.asc";
		},
		onCancel() {
			console.log("User cancelled the loader.");
		},
		async getData(query) {
			try {
				const getAllDatasetURI = this.BE_API + `/get_all_datasets?` + query;
				const getInfoDataURI = this.BE_API + `/get_total_data?` + query;
				const responseDataset = await axios.get(getAllDatasetURI, {
					headers: {
						"Content-Type": "application/json",
						"X-API-Key": "RAHASIA",
					},
				});
				const responseInfoData = await axios.get(getInfoDataURI, {
					headers: {
						"Content-Type": "application/json",
						"X-API-Key": "RAHASIA",
					},
				});

				this.pageData = responseDataset.data.data;
				this.totalDatasets = await responseInfoData.data.total_data;
				this.totalPages = await responseInfoData.data.total_page;
				console.log("ini total page : " + this.totalPages);
				console.log(responseDataset.data.data);
			} catch (error) {
				console.log(error);
			}
		},
		getPageNumber(value) {
			this.pageNumber = value;
			console.log("ini di getpagenumber");
			this.updateRoute();
		},
		getSearchValue(value) {
			this.searchValue.value = value;
			this.doSearch = true;
			console.log("ini di getsearch");
			this.updateRoute();
		},
		onSortingChange() {
			console.log("ini di getsorting");
			this.updateRoute();
		},
		getFilterData(value) {
			this.filterValue.tahunAwal = value.tahunAwal;
			this.filterValue.tahunAkhir = value.tahunAkhir;
			this.filterValue.topikPrefer = value.topikPrefer;
			this.filterValue.sumberDataPrefer = value.sumberDataPrefer;
			console.log("ini di getfilter");

			this.updateRoute();
		},
		async updatePage() {
			try {
				console.log("ada updaet page dari : " + this.prevQuery);
				if (this.currentQuery !== this.prevQuery) {
					console.log("masuk if apdet");
					console.log(this.currentQuery + "===" + this.prevQuery);
					this.prevQuery = this.currentQuery;
					this.isLoading = true;
					// simulate AJAX
					setTimeout(() => {
						this.isLoading = false;
					}, 4000);
					await this.getData(this.getQuery());
				}
			} catch (error) {
				console.log(error);
			}
		},
		updateRoute() {
			console.log("ini di update route: ");
			let filter = {};
			let path = "";
			// console.log(this.filterValue.topikPrefer);

			if (typeof this.filterValue.topikPrefer !== "undefined") {
				if (this.filterValue.topikPrefer !== "All") {
					filter.filter = "topik." + this.filterValue.topikPrefer;
				}
			}

			if (typeof this.filterValue.sumberDataPrefer !== "undefined") {
				if (this.filterValue.sumberDataPrefer !== "All") {
					if (typeof filter.filter === "undefined") {
						filter.filter =
							"sumber_data." + this.filterValue.sumberDataPrefer;
					} else {
						filter.filter +=
							"&sumber_data." + this.filterValue.sumberDataPrefer;
					}
				}
			}

			if (
				typeof this.filterValue.tahunAwal !== "undefined" &&
				typeof this.filterValue.tahunAkhir !== "undefined"
			) {
				filter.tahun =
					this.filterValue.tahunAwal + "." + this.filterValue.tahunAkhir;
			}

			if (typeof this.searchValue.value !== "undefined") {
				filter.search = this.searchValue.value;
			}
			filter.page = this.pageNumber;

			for (let prop in filter) {
				path += `${prop}=${filter[prop]}&`;
			}
			path += "sortBy=" + this.sortBy;

			this.prevQuery = this.currentQuery;
			this.currentQuery = path;
			path = "/dataset?" + path;

			console.log(this.currentQuery + "==?" + this.prevQuery);

			this.$router.push({
				path: path,
			});
			this.updatePage();
		},
	},
});
</script>
