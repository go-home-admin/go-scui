package menu

import (
	"encoding/json"
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
	guid.SetDb(mysql.NewOrmAdminMenu())
	return guid
}

func (g *GuiContext) Grid(view *table.View) {
	view.Column("ID", "id")
	view.Column("父级", "parent_id").Width("150")
	view.Column("排序", "order").Width("150")
	view.Column("组件名称", "name").Width("150")
	view.Column("组件", "component").Width("250")
	view.Column("地址", "path").Width("250")
	view.Column("菜单名称", "meta_title").Width("250").Display(func(val interface{}, row table.Row) interface{} {
		meta := row["meta"].(string)
		mm := make(map[string]interface{})
		json.Unmarshal([]byte(meta), &mm)
		return mm["title"]
	})
	view.Column("注册时间", "created_at").Width("250").Sortable(true)

	action := view.NewAction()
	action.AddButton("删除").Delete()
	action.AddButton("编辑").Edit()

	// 设置搜索栏
	filter := view.NewSearch()
	filter.Input("name", "名称").Span("8")

	header := view.NewHeader()
	header.Create()
}

func (g *GuiContext) Form(f *form.DialogForm) {
	f.LabelWidth("100px")
	f.Input("meta_title", "菜单名称")
	f.InputNumber("parent_id", "父级")
	f.Input("component", "组件").Placeholder("guid/index")
	f.Input("path", "地址")
}
