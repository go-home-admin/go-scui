<template>
	<load-sync-views/>
</template>

<script>
import {loadTableView} from "@/api/model/load";
import {defineAsyncComponent } from 'vue'
import {useRouter} from 'vue-router'

export default {
	name: 'sync',
	components: {
		"loadSyncViews": defineAsyncComponent(
			() =>
				new Promise((resolve) => {
					loadTableView(useRouter().currentRoute._rawValue.path).then(res => {
						let com = res.data

						let view = {
							template: com.template,
							data() {
								return Object.assign({}, com.data)
							},
							methods: com.methods ? eval("(" + com.methods + ")") : {},
							mounted() {
								if (com.mounted){
									eval(com.mounted)
								}
							},
						}
						console.log(view)
						resolve(view);
					});
				})
		)
	}
}
</script>
