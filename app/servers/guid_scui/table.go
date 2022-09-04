package guid_scui

import (
	"github.com/go-home-admin/go-admin/generate/proto/common/guid"
)

type Table struct {
	columns []*guid.Column
}

func NewTable() *Table {
	return &Table{
		columns: make([]*guid.Column, 0),
	}
}

func (g *Table) Column(label string, prop string) *Column {
	c := Column{
		&guid.Column{
			Label:   label,
			Prop:    prop,
			Filters: make([]*guid.Filter, 0),
		},
	}
	g.columns = append(g.columns, c.Column)
	return &c
}

func (g *Table) GetColumn() []*guid.Column {
	return g.columns
}

type Column struct {
	*guid.Column
}

func (c *Column) Width(v string) *Column {
	c.Column.Width = v
	return c
}

func (c *Column) Hide(v bool) *Column {
	c.Column.Hide = v
	return c
}

func (c *Column) Sortable(v bool) *Column {
	c.Column.Sortable = v
	return c
}

func (c *Column) Filters(v []*guid.Filter) *Column {
	c.Column.Filters = v
	return c
}

func (c *Column) Template(v string) *Column {
	c.Column.Template = v
	return c
}
