<template>
	<div class="font-poppins search-container">
		<div class="rounded py-5 shadow" style="background: #013d6b">
			<div class="text-white font-semibold text-2xl ml5">Search</div>
			<div class="text-white text-lg mt-5 ml-5">{{ formTitle }} :</div>
			<form class="flex flex-wrap mt3 w-full" onsubmit="return false">
				<div class="email-input-box flex w-full items-center ml5 mr5">
					<input
						id="search-card"
						class="w-full font-normal text-l ml-5 mr-5 outline-none color-navyblue"
						style="background: #f7f7f7"
						type="text"
						placeholder="Perpustakaan"
					/>
				</div>
				<div class="flex justify-center items-center mt5 w-full">
					<button
						type="submit"
						@click="getSearchValue"
						class="email-subscribe-btn w-11/12 text-blue font-semibold text-base"
					>
						Search..
					</button>
				</div>
			</form>
		</div>
	</div>
</template>

<style scoped>
@media screen and (width) {
}

.email-subscribe-btn {
	height: 40px;
	background: #28a0f6;
	box-sizing: border-box;
	border: 0.3px solid #28a0f6;
	border-radius: 5px;
	transition: background 0.1s ease-in-out;
}
/* .email-subscribe-btn:hover {
	background: #f37c38;
	color: #ffffff;
} */

.email-input-box {
	width: 274px;
	height: 48px;
	box-sizing: border-box;
	background: #f7f7f7;
	box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.15);
	border: 1px solid #142e46;
	border-radius: 5px;
}
</style>

<script>
import Vue from "vue";
export default Vue.extend({
	name: "SearchCard",
	props: ["formTitle", "isLong"],

	data() {
		return {
			searchValue: "",
		};
	},

	methods: {
		getSearchValue() {
			this.searchValue = document.getElementById("search-card").value;
			console.log(this.searchValue);
			this.passSearchValue();
		},

		passSearchValue() {
			this.$emit("passedSearchValue", this.searchValue);
		},
	},
});
</script>
