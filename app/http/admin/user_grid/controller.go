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
	table.Column("姓名", "nickname").Width("150")
	table.Column("性别", "sex").Width("150").Filters([]*grid.Filter{{Text: "男", Value: "1"}, {Text: "女", Value: "0"}})
	table.Column("邮箱", "email").Width("250")
	table.Column("注册时间", "created_at").Width("150").Sortable(true)
	table.Column("进度", "progress").Avatar()

	action := table.NewAction()
	action.AddButton("删除").Confirm("/del?id={{ row.id }}")

	// 设置搜索栏
	filter := table.NewSearch()
	filter.LabelWidth = "120px"
	filter.Input("name", "名称").Placeholder("这里是提示语").Span(12)
	filter.Input("nick", "昵称").Span(12)
	filter.Input("nick2", "xvsasf").Span(6)

	return table
}

func (receiver Controller) Form() *grid_scui.Form {
	form := grid_scui.NewForm()

	return form
}
