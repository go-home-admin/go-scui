package table

import (
	"strings"
)

// Image el-image
func (c *Column) Image() *Column {
	return c.Template(strings.ReplaceAll(`
<template #{k}="scope">
	<el-image
	  style="width: 100px; height: 100px;"
	  :src="scope.row.{k}"
	  :preview-src-list="[scope.row.{k}]"
	  preview-teleported="true"
	/>
</template>
`, "{k}", c.Column.Prop))
}
