<template>
	<login-popup />
</template>

<script>
	import {GetAdminUser} from "@/api/model/admin_user";
	import { defineAsyncComponent } from 'vue/dist/vue.esm-bundler.js'


	export default {
		name: 'tableCustomColumn',
		components: {
			"LoginPopup" :  defineAsyncComponent(
				() =>
					new Promise((resolve) => {
						GetAdminUser().then(res=>{
							console.log(res)
							resolve({
								template: " <el-progress :percentage=\"100\" status=\"success\" />",
							});
						});
					})
			)
		},
		data() {
			return {
				show: false
			}
		}
	}
</script>

<style>
</style>
