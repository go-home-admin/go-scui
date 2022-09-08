package grid_scui

import (
	_ "embed"
	"encoding/json"
	"github.com/go-home-admin/go-admin/generate/proto/common/grid"
	"github.com/go-home-admin/home/protobuf"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Table struct {
	db *gorm.DB
	// 存储列表信息
	columns []*grid.Column
	uri     string
	req     *grid.ListRequest
}

func NewTable(db *gorm.DB) *Table {
	return &Table{
		db:      db,
		columns: make([]*grid.Column, 0),
	}
}

func (g *Table) Paginate() ([]*protobuf.Any, int64) {
	var total int64
	list := make([]map[string]interface{}, 0)
	g.db.Count(&total)
	if total > 0 {
		if g.req.Page == 0 {
			g.req.Page = 1
		}

		offset := (g.req.Page - 1) * g.req.PageSize
		tx := g.db.Offset(int(offset)).Limit(int(g.req.PageSize)).Find(&list)
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
		Data:     g.GetData(),
		Methods:  g.GetMethods(),
		Mounted:  g.GetMounted(),
	}
}

func (g *Table) Search(req ...*grid.ListRequest) *Search {
	if len(req) != 0 {
		g.req = req[0]
	}
	return &Search{}
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
	return vTable
}

func (g *Table) GetMethods() string {
	return `{
	get: async function(params){
		console.log(this.url)
		return await this.http.get(this.url, params);
	},
}`
}

func (g *Table) GetMounted() string {
	return ""
}

func (g *Table) SetUri(v string) {
	g.uri = v
}

func (g *Table) GetData() string {
	str, err := json.Marshal(map[string]interface{}{
		"columns": g.GetColumn(),
		"url":     g.uri,
	})
	if err != nil {
		logrus.Error(err)
	}
	return string(str)
}
