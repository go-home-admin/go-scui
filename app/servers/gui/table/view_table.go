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
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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

func GetEditDialog(ctx *gin.Context, form, button base.RenderBase) *html.Dialog {
	var dia *html.Dialog
	i, ok := ctx.Get("__gui_edit__")
	if !ok {
		dia = html.NewDialog().SetContext(form)
		button.AddRender(dia)
		ctx.Set("__gui_edit__", dia)
	} else {
		dia = i.(*html.Dialog)
	}

	return dia
}

func GetForm(g GuiController) *form.DialogForm {
	var f *form.DialogForm
	i, ok := g.Gin().Get("__gui_form__")
	if !ok {
		f = form.NewForm()
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
	iniTable(t)
	return t
}

func iniTable(t *View) {

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
		offset := (Page - 1) * PageSize
		tx := g.GetDB().Offset(int(offset)).Limit(PageSize).Find(&list)
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
		"__template__", g.GetTemplate(),
		"__data__", string(d),
		"__mounted__", g.GetMounted(),
		"__methods__", g.GetMethods(),
	})
	_ = os.MkdirAll("resources/gui/", 0755)
	err := ioutil.WriteFile("resources/gui/"+strings.ReplaceAll(g.GetUri(), "/", "")+".vue", []byte(vueStr), 0755)
	if err != nil {
		logrus.Error(err)
	}
}

// NewAction 列操作
func (g *View) NewAction() *RowAction {
	action := &RowAction{Context: g.Controller, t: g}
	return action
}

// NewSearch 返回搜索栏
func (g *View) NewSearch() *Search {
	f := NewTableSearch(g.gin)
	g.View.AddRender(f, "search", "filter")
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
