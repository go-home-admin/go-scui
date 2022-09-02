<template>
	<div id="guid">test-template</div>
</template>

<script>
	// import {GetAdminUser} from "@/api/model/admin_user";

	import { defineCustomElement } from 'vue'

	export default {
		name: 'tableCustomColumn',
		data() {
			return {
				guidInit: false,
			}
		},
		mounted() {
			const MyVueElement = defineCustomElement({
				// 这里是同平常一样的 Vue 组件选项
				props: {},
				emits: {},
				template: `<div>123</div>`,

				// defineCustomElement 特有的：注入进 shadow root 的 CSS
				styles: [`/* inlined css */`]
			})


			// 注册自定义元素
			// 注册之后，所有此页面中的 `<my-vue-element>` 标签
			// 都会被升级
			// customElements.define('my-vue-element', MyVueElement)

			// 你也可以编程式地实例化元素：
			// （必须在注册之后）
			document.body.appendChild(
				new MyVueElement({
					// 初始化 props（可选）
				})
			)
		}
	}
</script>

<style>
</style>
