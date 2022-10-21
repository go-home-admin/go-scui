package user_grid

import (
	"github.com/go-home-admin/go-admin/app/entity/mysql"
	"github.com/go-home-admin/go-admin/app/servers/gui/form"
	"github.com/go-home-admin/go-admin/app/servers/gui/table"
	"github.com/go-home-admin/go-admin/generate/proto/common/grid"
	"github.com/go-home-admin/home/app/http"
)

// Controller @Bean
type Controller struct{}

type GuiController struct {
	http.Context
	*table.View
}

func NewGuiContext(ctx http.Context) *GuiController {
	guid := &GuiController{Context: ctx}
	view := table.NewTable(guid)
	return &GuiController{Context: ctx, View: view}
}

func (g *GuiController) Grid(view *table.View) {
	view.SetDb(mysql.NewOrmUser().GetDB())
	view.Column("姓名", "nickname").Width("150")
	view.Column("性别", "sex").Width("150").Filters([]*grid.Filter{{Text: "男", Value: "1"}, {Text: "女", Value: "0"}})
	view.Column("邮箱", "email").Width("150")
	view.Column("注册时间", "created_at").Width("250").Sortable(true)
	view.Column("进度", "progress").Avatar()

	action := view.NewAction()

	action.AddButton("删除").Confirm("/del?id={{ row.id }}")
	action.AddButton("编辑").Edit()

	// 设置搜索栏
	filter := view.NewSearch()
	filter.Input("name", "名称").Placeholder("这里是提示语").Span(12)
	filter.Input("nick", "昵称").Span(12)
	filter.Input("sex", "性别").Span(6)

	header := view.NewHeader()
	header.Create()
}

func (g *GuiController) Form(f *form.DialogForm) {
	f.Input("nickname", "姓名")
	f.Switch("flag", "是否预售")
	f.DatePicker("start_date", "开始时间")
	f.DatePicker("end_date", "结束时间")
	f.DatePicker("start_month", "开始月份").Options("type", "month")
	f.DatePicker("end_month", "结束月份").Options("type", "month")
	f.ColorPicker("my_color", "衣服颜色")

	catList := make([]*form.SelectOptionItem, 0)
	catList = append(catList, &form.SelectOptionItem{
		Value: "1",
		Label: "A",
	})
	catList = append(catList, &form.SelectOptionItem{
		Value: "2",
		Label: "B",
	})
	f.Select("cat_id", "选择分类").AddSelectOptions(f.Form, catList)

	f.Radio("sex", "男").Options(":label", "0")
	f.Radio("sex", "女").Options(":label", "1")
}
