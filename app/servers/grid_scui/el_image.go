package grid_scui

import "strings"

// Image el-image
func (c *Column) Image() *Column {
	return c.Template(strings.ReplaceAll(`
<template #{k}="scope">
  <div class="demo-image__preview">
    <el-image
      style="width: 100px; height: 100px"
      :src="scope.row.{k}"
      :preview-src-list="['scope.row.{k}']"
      :initial-index="4"
      fit="cover"
    />
  </div>
</template>
`, "{k}", c.Column.Prop))
}
