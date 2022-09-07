<template>
	<div class="font-lato search-container">
		<div class="rounded py-5" style="background: #d7e3fc">
			<div class="text-blue font-semibold text-xl ml5">Search</div>
			<div class="text-blue text-lg mt-5 ml-5">{{ formTitle }} :</div>
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
						class="email-subscribe-btn w-11/12 text-blue font-normal text-base"
					>
						Search..
					</button>
				</div>
			</form>
		</div>
		<!-- <div style="width: 90vw" class="font-lato ml-2 hidden-sm-and-up py-3">
			<form
				class="w-full email-input-box-long flex justify-center items-center px-3"
				onsubmit="return false"
			>
				<input
					id="search-card"
					class="w-full font-normal text-base ml-5 mr-5 outline-none color-navyblue text-truncate"
					type="text"
					placeholder="Ciletuh Palabuhanratu Geopark"
				/>
				<button
					type="submit"
					@click="getSearchValue"
					class="email-subscribe-btn-mobile text-white font-normal flex justify-center items-center text-xl"
				>
					<ion-icon name="search-sharp" class="text-white"></ion-icon>
				</button>
			</form>
		</div> -->
	</div>
</template>

<style scoped>
@media screen and (width) {
}

.email-subscribe-btn-long {
	width: 165px;
	height: 50px;
	background: #00b4a7;
	box-sizing: border-box;
	border-radius: 10px;
	transition: background 0.1s ease-in-out;
}
.email-subscribe-btn-mobile {
	width: 80px;
	height: 50px;
	background: #00b4a7;
	box-sizing: border-box;
	border-radius: 10px;
	transition: background 0.1s ease-in-out;
}
.email-subscribe-btn {
	height: 40px;
	background: #abc3ff;
	box-sizing: border-box;
	border: 0.3px solid #142e46;
	border-radius: 5px;
	transition: background 0.1s ease-in-out;
}

.email-subscribe-btn:hover {
	background: #abc3ff;
}

.email-input-box-long {
	height: 64px;
	box-sizing: border-box;
	background: #ffffff;
	box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.15);
	border-radius: 10px;
}
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
