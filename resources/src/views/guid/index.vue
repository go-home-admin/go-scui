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

						resolve({
							// props: com.props ? eval("(" + com.props + ")") : {},
							template: com.template,
							data() {
								return eval("(" + com.data + ")")
							},
							methods: com.methods ? eval("(" + com.methods + ")") : {},
							mounted() {
								this.config = config
								this.http = http

								com.mounted ? eval(com.mounted) : {}
							},
						});
					});
				})
		)
	}
}
</script>
