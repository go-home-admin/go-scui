package grid_scui

import (
	_ "embed"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/generate/proto/common/grid"
	"github.com/go-home-admin/home/app"
	"github.com/go-home-admin/home/app/http"
	"github.com/go-home-admin/home/protobuf"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"strconv"
)

type Table struct {
	http.Context

	db *gorm.DB
	// 存储列表信息
	columns []*grid.Column
	uri     string

	filter *Search
	action *RowAction
	// [函数名称]代码
	MethodsRender map[string]func(*Table) string
	MountedRender map[string]func(*Table) string
}

func NewTable(ctx http.Context, db *gorm.DB) *Table {
	return &Table{
		Context: ctx,
		db:      db,
		columns: make([]*grid.Column, 0),
		uri:     "",
		MethodsRender: map[string]func(*Table) string{
			"get": func(*Table) string {
				return "async function(params){\n\t\treturn await this.http.get(this.url, params);\n\t}"
			},
			"post": func(*Table) string {
				return "async function(params){\n\t\treturn await this.http.post(this.url, params);\n\t}"
			},
		},
		MountedRender: nil,
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

func (g *Table) Paginate() ([]*protobuf.Any, int64) {
	var total int64
	list := make([]map[string]interface{}, 0)
	g.db.Count(&total)
	if total > 0 {
		Page := GetInt(g.Gin(), "page", 1)
		PageSize := GetInt(g.Gin(), "page", 20)
		offset := (Page - 1) * PageSize
		tx := g.db.Offset(int(offset)).Limit(PageSize).Find(&list)
		if tx.Error != nil {
			logrus.Error(tx.Error)
		}
	}
	got := make([]*protobuf.Any, 0)
	for _, m := range list {
		s := protobuf.NewAny(m)
		got = append(got, s)
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

func (g *Table) RowAction() *RowAction {
	g.action = &RowAction{}
	return g.action
}

// Search 返回搜索栏
func (g *Table) Search() *Search {
	g.filter = &Search{}
	return g.filter
}

func (g *Table) Column(label string, prop string) *Column {
	c := Column{
		Column: &grid.Column{
			Label:   label,
			Prop:    prop,
			Filters: make([]*grid.Filter, 0),
		},
		Render: &Render{},
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

//go:embed views/table.vue
var vTable string

func (g *Table) GetTemplate() string {
	if app.IsDebug() {
		if _, err := os.Stat("app/servers/grid_scui/views/table.vue"); err == nil {
			v, _ := ioutil.ReadFile("app/servers/grid_scui/views/table.vue")
			vTable = string(v)
		}
	}

	return vTable
}

// AddMethods js函数定义
func (g *Table) AddMethods(name string, fun func(*Table) string) {
	if _, ok := g.MethodsRender[name]; !ok {
		g.MethodsRender[name] = fun
	}
}

// AddMethodsCode js函数定义
func (g *Table) AddMethodsCode(name string, code string) {
	if _, ok := g.MethodsRender[name]; !ok {
		g.MethodsRender[name] = func(table *Table) string {
			return code
		}
	}
}

// GetMethods 获取所有js函数定义
func (g *Table) GetMethods() string {
	str := "{\n"
	for s, f := range g.MethodsRender {
		str = str + "\n" + s + ":" + f(g) + ",\n"
	}
	return str + "\n}"
}

func (g *Table) GetMounted() string {
	return ""
}

func (g *Table) SetUri(v string) {
	g.uri = app.Config("app.url", "http:127.0.0.1:8080") + v
}

func (g *Table) GetData() string {
	data := map[string]interface{}{
		"columns": g.GetColumn(),
		"url":     g.uri,
	}

	if g.filter != nil {
		data["filter"] = g.filter.ToString()
	}

	str, err := json.Marshal(data)
	if err != nil {
		logrus.Error(err)
	}
	return string(str)
}
