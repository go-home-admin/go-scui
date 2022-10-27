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
	view.Column("城市", "city")
	view.Column("图片", "image")
	view.Column("内容", "content")

	action := view.NewAction()
	action.AddButton("删除").Delete().Options("type", "danger")
	action.AddButton("编辑").Edit().Options("type", "primary")

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
	f.Width("90%")
	f.Input("avatar", "头像")
	f.Input("nickname", "姓名")
	f.Select("sex", "性别").SelectOptions(gui.SelectOptions{
		{Label: "男", Value: 1},
		{Label: "女", Value: 0},
	})
	f.Input("phone", "手机号").SaveToInt()
	f.Input("email", "邮箱").Rules(true)
	f.DateTimePicker("created_at", "注册时间")
	f.CheckBox("city", "城市").Options(gui.SelectOptions{
		{Label: "广州", Value: "广州"},
		{Label: "成都", Value: "成都"},
	})
	f.File("image", "图片").Options(gui.FileLists{
		{Name: "ceshi", Url: "https://fuss10.elemecdn.com/3/63/4e7f3a15429bfda99bce42a18cdd1jpeg.jpeg?imageMogr2/thumbnail/360x360/format/webp/quality/100"},
	})
	f.Input("content", "内容").Options("type", "textarea")
}
