<template>
	<load-views />
</template>

<script>
import {GetAdminUser} from "@/api/model/admin_user";
import {defineAsyncComponent} from 'vue/dist/vue.esm-bundler.js'

export default {
	name: 'sync',
	components: {
		"loadViews": defineAsyncComponent(
			() =>
				new Promise((resolve) => {
					GetAdminUser().then(res => {
						console.log(res)
						resolve({
							template: " <el-progress :percentage=\"100\" status=\"success\" />",
						});
					});
				})
		)
	},
}
</script>
