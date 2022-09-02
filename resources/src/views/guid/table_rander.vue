<template>
	<cmp :json="json"></cmp>
</template>

<script>
import { defineComponent, h } from 'vue';

const json = [
	{
		type: 'div',
		props: { id: 'd1', class: 'd1' },
		children: [
			{ type: 'a', props: { innerHTML: '链接1' } },
			{ type: 'p', props: { innerHTML: '文本1' } },
		],
	},
	{
		type: 'div',
		props: { id: 'd2', class: 'd2' },
		children: [
			{ type: 'a', props: { innerHTML: '链接2' } },
			{ type: 'p', props: { innerHTML: '文本2' } },
		],
	},
];

export default defineComponent({
	name: 'tableRender',
	components: {
		cmp: {
			props: {
				json: {
					type: Array,
					default() {
						return [];
					},
				},
			},
			render() {
				function r(arr) {
					return arr.map((item) => {
						return h(
							item.type,
							item.props,
							item.children ? r(item.children) : item.props
						);
					});
				}
				return r(this.json);
			},
		},
	},
	props: {
		json: {
			type: Array,
			default() {
				return json;
			},
		},
	},
});
</script>
