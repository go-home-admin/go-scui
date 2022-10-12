package table

import (
	"strings"
)

// 进度条
func (c *Column) progress() *Column {
	return c.Template(strings.ReplaceAll(`
<Template #{k}="scope">
	<el-progress :percentage="scope.row.{k}" />
</Template>
`, "{k}", c.Column.Prop))
}
