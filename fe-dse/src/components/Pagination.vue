<template>
	<div class="flex justify-center font-lato mb2 mt1">
		<nav class="flex items-center">
			<button
				@click="backPage"
				:class="{
					'pagination-item-active': currentPage === 1,
					'pagination-item': currentPage !== 1,
				}"
			>
				<ion-icon class="pt-2 mr-2" name="chevron-back-outline"></ion-icon>
			</button>
			<button
				v-if="currentPage > 3"
				@click="firstPage"
				:class="{
					'pagination-item-active': currentPage === 1,
					'pagination-item': currentPage !== 1,
				}"
			>
				1
			</button>
			<span v-if="itemStart > 1" class="mr-1 text-xl mx-3">...</span>
			<!-- <span v-if="itemStart > 1" class="text-xl">/</span> -->
			<div
				v-for="pageNumber in pageNumbers.slice(itemStart - 1, itemStop)"
				:key="pageNumber"
			>
				<button
					@click="goToPage(pageNumber)"
					:class="{
						'pagination-item-active': pageNumber === currentPage,
						'pagination-item': pageNumber !== currentPage,
					}"
					class="mx-3"
				>
					{{ pageNumber }}
				</button>
				<!-- <span :class="{ hidden: pageNumber === itemStop }" class="text-xl"
					>/</span
				> -->
			</div>
			<div v-if="itemStop < totalPages - 1">
				<span class="mr-1 mx-3">...</span>
				<button @click="goToPage(totalPages)" class="pagination-item mx-3">
					{{ totalPages }}
				</button>
			</div>
			<button
				@click="nextPage"
				:class="{
					'pagination-item-active': currentPage === totalPages,
					'pagination-item': currentPage !== totalPages,
				}"
			>
				<ion-icon
					class="pt-2 ml-2"
					name="chevron-forward-outline"
				></ion-icon>
			</button>
			<!-- <a
				href="#"
				@click="goToPage(totalPages)"
				:class="{
					'pagination-item-active': currentPage === totalPages,
					'pagination-item': currentPage !== totalPages,
				}"
				>Last</a
			> -->
		</nav>
	</div>
</template>

<style scoped>
.pagination-item-active {
	color: #00b4a7;
	pointer-events: none;
	font-size: 20px;
	text-decoration: none;
	margin: 0 4px;
}

.pagination-item:hover {
	color: #00b4a7;
}

.pagination-item {
	color: #142e46;
	font-size: 20px;
	/* text-decoration: underline; */
	margin: 0 4px;
}
</style>

<script>
import Vue from "vue";

export default Vue.extend({
	name: "Pagination",
	props: ["totalPages", "pageRange"],

	data() {
		return {
			// totalPages: 1,
			itemStart: 0,
			itemStop: 0,
			currentPage: 1,
			pageNumbers: [0],
			perPage: 10,
			// pageData: [{}],
		};
	},

	methods: {
		passPageNumber() {
			this.$emit("passedPageNumber", this.currentPage);
		},

		// updatePageData() {
		// 	this.pageData = this.data.slice(
		// 		(this.currentPage - 1) * this.perPage,
		// 		this.currentPage * this.perPage
		// 	);
		// },

		updateStartStop() {
			if (this.currentPage <= Math.floor((this.pageRange - 1) / 2)) {
				this.itemStart = 1;
			} else {
				this.itemStart =
					this.currentPage - Math.floor((this.pageRange - 1) / 2);
			}
			if (
				this.currentPage >=
				this.totalPages - Math.ceil((this.pageRange - 1) / 2)
			) {
				this.itemStop = this.totalPages;
			} else {
				this.itemStop =
					this.currentPage + Math.ceil((this.pageRange - 1) / 2);
			}
			// this.updatePageData();
			// this.passPageData();
		},

		goToPage(page) {
			this.currentPage = page;
			this.updateStartStop();
			this.passPageNumber();
		},
		backPage() {
			this.currentPage -= 1;
			this.updateStartStop();
			this.passPageNumber();
		},
		nextPage() {
			this.currentPage += 1;
			this.updateStartStop();
			this.passPageNumber();
		},
		firstPage() {
			this.currentPage = 1;
			this.updateStartStop();
			this.passPageNumber();
		},
	},

	created() {
		// this.totalPages = Math.ceil(this.data.length / this.perPage);
		this.updateStartStop();
		for (let i = 1; i <= this.totalPages; i++) {
			this.pageNumbers[i - 1] = i;
		}
		// this.pageData = this.data.slice(
		// 	this.currentPage * 0 * this.perPage,
		// 	this.currentPage * this.perPage
		// );
		this.passPageNumber();
	},
});
</script>
