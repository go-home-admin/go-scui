package table

import (
	_ "embed"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
	"github.com/go-home-admin/go-admin/app/servers/gui/form"
	"github.com/go-home-admin/go-admin/generate/proto/common/grid"
	"github.com/go-home-admin/home/app"
	"github.com/go-home-admin/home/app/http"
	"github.com/go-home-admin/home/protobuf"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io/ioutil"
	"strconv"
	"strings"
)

type Table struct {
	*base.View
	Context GuiContext
	db      *gorm.DB
	// 存储列表信息
	columns []*grid.Column
	uri     string

	action *RowAction
}

type GuiContext interface {
	http.Context
	Form(f *form.DialogForm)
}

func GetForm(g GuiContext) *form.DialogForm {
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

func NewTable(ctx GuiContext, db *gorm.DB) *Table {
	t := &Table{
		View:    base.NewView("table.vue"),
		Context: ctx,
		db:      db,
		columns: make([]*grid.Column, 0),
		uri:     "",
	}
	iniTable(t)
	return t
}

func iniTable(t *Table) {

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

// Paginate 列表数据，分页获取
func (g *Table) Paginate() ([]*protobuf.Any, int64) {
	var total int64
	list := make([]map[string]interface{}, 0)
	g.db.Count(&total)
	if total > 0 {
		Page := GetInt(g.Context.Gin(), "page", 1)
		PageSize := GetInt(g.Context.Gin(), "pageSize", 20)
		offset := (Page - 1) * PageSize
		tx := g.db.Offset(int(offset)).Limit(PageSize).Find(&list)
		if tx.Error != nil {
			logrus.Error(tx.Error)
		}
	}
	got := make([]*protobuf.Any, 0)
	for _, m := range list {
		got = append(got, protobuf.NewAny(m))
	}
	return got, total
}

func (g *Table) ToResponse() *grid.IndexResponse {
	if app.IsDebug() {
		defer toVueFile(*g)
	}

	return &grid.IndexResponse{
		Template: g.GetTemplate(),
		Data:     protobuf.NewAny(g.GetData()),
		Methods:  g.GetMethods(),
		Mounted:  g.GetMounted(),
	}
}

func toVueFile(g Table) {
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
	err := ioutil.WriteFile("resources/gui/"+strings.ReplaceAll(g.GetUri(), "/", "")+".vue", []byte(vueStr), 0755)
	if err != nil {
		logrus.Error(err)
	}
}

// NewAction 列操作
func (g *Table) NewAction() *RowAction {
	action := &RowAction{Context: g.Context, t: g}
	return action
}

// NewSearch 返回搜索栏
func (g *Table) NewSearch() *Search {
	f := NewTableSearch(g.Context)
	g.View.AddRender(f, "search", "filter")
	return f
}

func (g *Table) NewHeader() *Header {
	r := NewHeader(g.Context)
	g.View.AddRender(r, "header")
	return r
}

func (g *Table) Column(label string, prop string) *Column {
	c := Column{
		Column: &grid.Column{
			Label:   label,
			Prop:    prop,
			Filters: make([]*grid.Filter, 0),
		},
		Render: base.NewRender(),
	}
	g.columns = append(g.columns, c.Column)
	return &c
}

func (g *Table) GetColumn() []*grid.Column {
	for _, column := range g.columns {
		if column.Template == "" {
			// column.Template = column.R
		}
	}

	return g.columns
}

func (g *Table) SetUri(v string) {
	g.uri = app.Config("app.url", "http:127.0.0.1:8080") + v
}

func (g *Table) GetUri() string {
	if g.uri == "" {
		g.uri = app.Config("app.url", "http:127.0.0.1:8080") + g.Context.Gin().Request.URL.Path + "/list"
	}
	return g.uri
}

func (g *Table) GetData() map[string]interface{} {
	g.View.AddData("columns", g.GetColumn())
	g.View.AddData("url", g.GetUri())

	return g.View.GetData()
}
