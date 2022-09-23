package user_grid

import (
	"github.com/go-home-admin/go-admin/app/entity/mysql"
	"github.com/go-home-admin/go-admin/app/servers/grid_scui"
	"github.com/go-home-admin/go-admin/generate/proto/common/grid"
	"github.com/go-home-admin/home/app/http"
)

// Controller @Bean
type Controller struct {
}

func (receiver Controller) NewGridTable(ctx http.Context) *grid_scui.Table {
	table := grid_scui.NewTable(ctx, mysql.NewOrmUser().GetDB())
	table.Column("姓名", "name").Width("150")
	table.Column("性别", "sex").Width("150").Filters([]*grid.Filter{{Text: "男", Value: "男"}, {Text: "女", Value: "女"}})
	table.Column("邮箱", "email").Width("250")
	table.Column("注册时间", "datetime").Width("150").Sortable(true)
	table.Column("进度", "progress").Avatar()

	action := table.RowAction()
	action.Add()
	// 设置搜索栏
	filter := table.Search()
	filter.LabelWidth = "120px"
	filter.Input("name", "名称").Placeholder("这里是提示语").Span(12)
	filter.Input("nick", "昵称").Span(12)
	filter.Input("nick2", "xvsasf").Span(6)

	return table
}
