package guid_scui

import "strings"

// Avatar el-avatar
func (c *Column) Avatar(opt ...string) *Column {
	opts := " :src=\"" + c.Column.Prop + "\""
	for _, s := range opt {
		opts = opts + " " + s
	}
	template := strings.ReplaceAll(`
<template #{k}="scope">
  <el-avatar {opts} />
</template>
`, "{k}", c.Column.Prop)
	template = strings.ReplaceAll(template, "{opts}", opts)
	return c.Template(template)
}
