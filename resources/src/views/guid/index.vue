<template>
	<load-sync-views/>
</template>

<script>
import {loadTableView} from "@/api/model/load";
import {defineAsyncComponent} from 'vue'
import {useRouter} from 'vue-router'
import config from "@/config"
import http from "@/api/http"

export default {
	name: 'sync',
	components: {
		"loadSyncViews": defineAsyncComponent(
			() =>
				new Promise((resolve) => {
					loadTableView(useRouter().currentRoute._rawValue.path).then(res => {
						let com = res.data
						let view = {
							// props: com.props ? eval("(" + com.props + ")") : {},
							template: com.template,
							data() {
								return Object.assign(
									{
										config,
										http
									},
									com.data
								)
							},
							methods: com.methods ? eval("(" + com.methods + ")") : {},
							mounted() {
								com.mounted ? eval(com.mounted) : {}
							},
						}
						console.log(view, eval("(" + com.data + ")"))
						resolve(view);
					});
				})
		)
	}
}
</script>
