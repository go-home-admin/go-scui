package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/app/entity/mysql"
	"github.com/go-home-admin/go-admin/app/servers/gui"
	"github.com/go-home-admin/go-admin/app/servers/gui/form"
	"github.com/go-home-admin/go-admin/app/servers/gui/table"
)

// GinHandleResource gin原始路由处理
func (receiver *Controller) GinHandleResource(ctx *gin.Context) {
	NewGuiContext(ctx).ActionHandle()
}

type GuiContext struct {
	*gui.GinHandle
	*table.View
}

func NewGuiContext(ctx *gin.Context) *GuiContext {
	guid := &GuiContext{GinHandle: gui.NewGui(ctx)}
	guid.View = table.NewTable(guid)
	guid.SetController(guid)
	guid.SetDb(mysql.NewOrmUser().GetDB())
	return guid
}

func (g *GuiContext) Grid(view *table.View) {
	view.Column("ID", "id").Width("80")
	view.Column("头像", "avatar").Avatar().Width("100")
	view.Column("姓名", "nickname").Width("150")
	view.Column("性别", "sex").Width("150").Filters([]gui.Filter{{Text: "男", Value: 1}, {Text: "女", Value: 0}})
	view.Column("邮箱", "email").Width("150")
	view.Column("注册时间", "created_at").Date().Width("250").Sortable(true)

	action := view.NewAction()
	action.AddButton("删除").Delete()
	action.AddButton("编辑").Edit()

	// 设置搜索栏
	filter := view.NewSearch()
	filter.Input("name", "名称")
	filter.Input("nick", "昵称")
	filter.Input("sex", "性别")

	header := view.NewHeader()
	header.Create()
}

func (g *GuiContext) Form(f *form.DialogForm) {
	f.LabelWidth("80px")
	f.Input("avatar", "头像")
	f.Input("nickname", "姓名")
	f.Select("sex", "性别").SelectOptions(gui.SelectOptions{
		{Label: "男", Value: 1},
		{Label: "女", Value: 0},
	})
	f.Input("phone", "手机号").SaveToInt()
	f.Input("email", "邮箱")
	f.DateTimePicker("created_at", "注册时间")
}
