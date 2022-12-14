package table

import (
	_ "embed"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/app/servers/gui"
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
	"github.com/go-home-admin/go-admin/app/servers/gui/form"
	"github.com/go-home-admin/go-admin/app/servers/gui/html"
	"github.com/go-home-admin/home/app"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strconv"
)

type View struct {
	*base.View
	*gui.ViewDB
	*gui.ViewPrimary
	gin        *gin.Context
	Controller GuiController
	// 存储列表信息
	columns []*Column
	uri     string

	action *RowAction
}

type GuiController interface {
	Gin() *gin.Context
	Grid(f *View)
	Form(f *form.DialogForm)
}

func GetEditDialog(g GuiController, button base.RenderBase) (*html.Dialog, *form.DialogForm) {
	ctx := g.Gin()
	formRender := GetForm(g)
	dia := formRender.Dialog
	_, ok := ctx.Get("__gui_edit__")
	if !ok {
		button.AddRender(dia)
		ctx.Set("__gui_edit__", true)
	}

	return dia, formRender
}

func GetForm(g GuiController) *form.DialogForm {
	var f *form.DialogForm
	i, ok := g.Gin().Get("__gui_form__")
	if !ok {
		f = form.NewDialogForm()
		g.Form(f)
		g.Gin().Set("__gui_form__", f)
	} else {
		f = i.(*form.DialogForm)
	}
	return f
}

func NewTable(controller GuiController) *View {
	t := &View{
		View:        base.NewView("table.vue"),
		ViewDB:      &gui.ViewDB{},
		ViewPrimary: &gui.ViewPrimary{},
		gin:         controller.Gin(),
		Controller:  controller,
		columns:     make([]*Column, 0),
		uri:         "",
	}

	return t
}

func GetInt(ctx *gin.Context, k string, def int) int {
	v := ctx.Query(k)
	if v == "" {
		return def
	}
	i, err := strconv.Atoi(v)
	if err != nil {
		logrus.Error("GetInt", err)
		return 0
	}
	return i
}

func (g *View) GetFormItems() []*gui.FormItems {
	return GetForm(g.Controller).GetFormItems()
}

// Paginate 列表数据，分页获取
func (g *View) Paginate() (interface{}, int64) {
	g.Controller.Grid(g)

	var total int64
	list := make([]map[string]interface{}, 0)
	g.GetDB().Count(&total)
	if total > 0 {
		Page := GetInt(g.Controller.Gin(), "page", 1)
		PageSize := GetInt(g.Controller.Gin(), "pageSize", 20)
		tx := g.GetDB().Offset((Page - 1) * PageSize).Limit(PageSize).Find(&list)
		if tx.Error != nil {
			logrus.Error(tx.Error)
		}
		for _, column := range g.columns {
			for _, f := range column.funcs {
				for i, row := range list {
					row[column.Prop] = f(row[column.Prop], row)
					list[i] = row
				}
			}
		}
	}

	return list, total
}

func (g *View) ToResponse() *gui.IndexResponse {
	g.Controller.Grid(g)
	if app.IsDebug() {
		defer toVueFile(*g)
	}

	return &gui.IndexResponse{
		Template: g.GetTemplate(),
		Data:     g.GetData(),
		Methods:  g.GetMethods(),
		Mounted:  g.GetMounted(),
	}
}

func toVueFile(g View) {
	// 在本地生成一个调试模版
	vueStr := `
<template>__template__</template>
<script>
export default {
  data() {
    return __data__
  },
  mounted() {
    __mounted__
  },
  methods: __methods__
}
</script>
`
	d, _ := json.Marshal(g.GetData())
	vueStr = base.ReplaceAll(vueStr, []string{
		"__template__", (*g.Render).GetTemplate(),
		"__data__", string(d),
		"__mounted__", (*g.Render).GetMounted(),
		"__methods__", (*g.Render).GetMethods(),
	})
	_ = os.MkdirAll("resources/gui/", 0755)
	err := ioutil.WriteFile("resources/gui/"+uuid.NewV4().String()+".vue", []byte(vueStr), 0755)
	if err != nil {
		logrus.Error(err)
	}
}

// NewAction 列操作
func (g *View) NewAction() *RowAction {
	action := &RowAction{
		Render: base.NewRender(`
        <el-table-column label="操作" fixed="right" align="right" width="__column__width">
          <template #default="{row}">
            <slot id="actions"/>
          </template>
        </el-table-column>`),
		Context: g.Controller,
		t:       g,
	}
	action.Width("160")
	g.AddRender(action, "el-table-column")
	return action
}

// NewSearch 返回搜索栏
func (g *View) NewSearch() *Search {
	g.AddData("filterShow", true)
	f := NewTableSearch(g.gin)
	g.View.AddRender(f, "search")
	return f
}

func (g *View) NewHeader() *Header {
	r := NewHeader(g.Controller)
	g.View.AddRender(r, "header")
	return r
}

func (g *View) Column(label string, prop string) *Column {
	c := NewColumn()
	c.Column.Label = label
	c.Column.Prop = prop
	g.columns = append(g.columns, c)
	return c
}

func (g *View) GetTemplate(pr ...base.RenderBase) string {
	for _, column := range g.columns {
		if column.Render.Template != "" {
			g.AddRender(column, "el-table-column")
		}
	}

	return g.Render.GetTemplate(pr...)
}

func (g *View) GetColumn() []*Column {
	return g.columns
}

func (g *View) SetUri(v string) {
	g.uri = app.Config("app.url", "http:127.0.0.1:8080") + v
}

func (g *View) GetUri() string {
	if g.uri == "" {
		g.uri = app.Config("app.url", "http:127.0.0.1:8080") + g.Controller.Gin().Request.URL.Path + "/list"
	}
	return g.uri
}

func (g *View) GetData() map[string]interface{} {
	g.View.AddData("columns", g.GetColumn())
	g.View.AddData("url", g.GetUri())
	g.View.AddData("primary", g.GetPrimary())

	return g.View.GetData()
}
