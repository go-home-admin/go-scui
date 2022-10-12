package gui

import "strings"

// 进度条
func (c *Column) progress() *Column {
	return c.Template(strings.ReplaceAll(`
<template #{k}="scope">
	<el-progress :percentage="scope.row.{k}" />
</template>
`, "{k}", c.Column.Prop))
}
