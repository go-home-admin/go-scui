package user_grid

import (
	"github.com/go-home-admin/go-admin/app/entity/mysql"
	"github.com/go-home-admin/go-admin/app/servers/gui/form"
	t "github.com/go-home-admin/go-admin/app/servers/gui/table"
	"github.com/go-home-admin/go-admin/generate/proto/common/grid"
	"github.com/go-home-admin/home/app/http"
)

// Controller @Bean
type Controller struct{}

type GuiContext struct {
	http.Context
}

func NewGuiContext(ctx http.Context) *GuiContext {
	return &GuiContext{Context: ctx}
}

func (g *GuiContext) NewGridTable() *t.Table {
	table := t.NewTable(g, mysql.NewOrmUser().GetDB())
	table.Column("姓名", "nickname").Width("150")
	table.Column("性别", "sex").Width("150").Filters([]*grid.Filter{{Text: "男", Value: "1"}, {Text: "女", Value: "0"}})
	table.Column("邮箱", "email").Width("150")
	table.Column("注册时间", "created_at").Width("250").Sortable(true)
	table.Column("进度", "progress").Avatar()

	action := table.NewAction()
	action.AddButton("删除").Confirm("/del?id={{ row.id }}")
	action.AddButton("编辑").Edit()

	// 设置搜索栏
	filter := table.NewSearch()
	filter.Input("name", "名称").Placeholder("这里是提示语").Span(12)
	filter.Input("nick", "昵称").Span(12)
	filter.Input("sex", "性别").Span(6)

	header := table.NewHeader()
	header.Create()

	return table
}

func (g *GuiContext) Form(f *form.DialogForm) {
	f.Input("nickname", "名称")

}
