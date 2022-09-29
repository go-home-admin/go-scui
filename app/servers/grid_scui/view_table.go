package grid_scui

import (
	_ "embed"
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/generate/proto/common/grid"
	"github.com/go-home-admin/home/app"
	"github.com/go-home-admin/home/app/http"
	"github.com/go-home-admin/home/protobuf"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strconv"
)

type Table struct {
	*View
	http.Context

	db *gorm.DB
	// 存储列表信息
	columns []*grid.Column
	uri     string

	action *RowAction
}

func NewTable(ctx http.Context, db *gorm.DB) *Table {
	view := NewView("table.vue")
	view.AddMethods("getData", "async function(params){\n\t\treturn await this.$HTTP.get(this.url, params);\n\t}")
	return &Table{
		View:    view,
		Context: ctx,
		db:      db,
		columns: make([]*grid.Column, 0),
		uri:     "",
	}
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
		Page := GetInt(g.Gin(), "page", 1)
		PageSize := GetInt(g.Gin(), "pageSize", 20)
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
	return &grid.IndexResponse{
		Template: g.GetTemplate(),
		Data:     protobuf.NewAny(g.GetData()),
		Methods:  g.GetMethods(),
		Mounted:  g.GetMounted(),
	}
}

// NewAction 列操作
func (g *Table) NewAction() *RowAction {
	action := &RowAction{Context: g.Context, t: g}
	return action
}

// NewSearch 返回搜索栏
func (g *Table) NewSearch() *Search {
	form := NewTableSearch(g.Context)
	g.View.AddRender(form, "search", "filter")
	return form
}

func (g *Table) Column(label string, prop string) *Column {
	c := Column{
		Column: &grid.Column{
			Label:   label,
			Prop:    prop,
			Filters: make([]*grid.Filter, 0),
		},
		Render: NewRender(),
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
