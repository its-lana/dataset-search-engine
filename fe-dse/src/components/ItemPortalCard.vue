<template>
	<div>
		<div
			v-for="(item, index) in data"
			:key="index"
			class="item-card w-full shadow-none mx-2 mb-7"
		>
			<div class="flex flex-row h-fit">
				<div class="ml3 mb1 w-full flex flex-col">
					<!-- Name -->
					<a
						:href="item.tautan_dataset_terkait"
						target="_blank"
						class="hover-underline font-lato w-full text-black font-bold text-xl mb1 mt2"
					>
						{{ item.judul }}
					</a>

					<div class="mt1 grid grid-cols-4 gap-2">
						<!-- Location -->
						<div>
							<!-- <ion-icon class="mr4" name="calendar-clear"></ion-icon> -->
							<ion-icon
								class="mr4 color-blue"
								name="calendar"
							></ion-icon>
							<span
								v-if="item.tahun_awal_data == item.tahun_akhir_data"
								class="font-normal font-lato text-sm text-l my-1"
							>
								{{ item.tahun_awal_data }}
							</span>
							<span
								v-else
								class="font-normal font-lato text-sm text-l my-1"
								>{{ item.tahun_awal_data }} -
								{{ item.tahun_akhir_data }}
							</span>
						</div>
						<div>
							<ion-icon
								class="mr4 color-blue"
								name="briefcase"
							></ion-icon>
							<a
								:href="item.tautan_sumber_data"
								style="color: #000000"
								class="hover-underline font-normal font-lato text-sm text-l"
								>{{ item.sumber_data }}</a
							>
						</div>
						<div>
							<ion-icon
								class="mr4 color-blue"
								name="business-sharp"
							></ion-icon>
							<span class="font-normal font-lato text-sm text-l my-1">
								{{ item.organisasi_terkait }}
							</span>
						</div>
						<div>
							<ion-icon
								class="mr4 color-blue"
								name="time-sharp"
							></ion-icon>
							<span class="font-normal font-lato text-sm text-l">
								{{ formatDate(item.last_updated) }}
							</span>
						</div>

						<div>
							<ion-icon
								class="color-green mr4"
								name="checkmark-circle"
							></ion-icon>
							<span class="font-normal font-lato text-sm text-l">
								{{ item.jenis_data }}
							</span>
						</div>
					</div>
					<!-- Facility Tags -->
				</div>
			</div>
		</div>
	</div>
</template>

<style scoped>
.color-green {
	color: #32cd32;
}
.color-blue {
	color: #013d6b;
}
.order-card {
	border: 1px solid #eeeeee;
	box-sizing: border-box;
}
.hover-underline:hover {
	text-decoration-line: underline;
}

.item-category {
	background: #ffffff;
	border: 1px solid #00b4a7;
	box-sizing: border-box;
	border-radius: 30px;
}

.item-card {
	border: 1px solid #eeeeee;
	box-sizing: border-box;
	border-radius: 16px;
}

.item-tag {
	margin: 5px 1%;
	background: rgba(0, 0, 0, 0.12);
	border-radius: 10px;
	font-size: 12px;
}
</style>

<script>
import Vue from "vue";

export default Vue.extend({
	// components: { StarCard },
	name: "ItemPortalCard",
	props: ["data"],

	data() {
		return {
			parent: "",
		};
	},

	created() {
		let pathArray = window.location.pathname.split("/");
		this.parent = pathArray[1];
	},

	methods: {
		formatDate(date) {
			let subDate = date.substring(0, 10).split("-"); //2021-11-18T00:00:00Z 12 Jan 2001
			let monthName = this.getMonthName(parseInt(subDate[1]));
			return subDate[2] + " " + monthName + " " + subDate[0];
		},
		getMonthName(monthNumber) {
			const date = new Date();
			date.setMonth(monthNumber - 1);

			return date.toLocaleString("en-US", { month: "short" });
		},
	},
});
</script>
