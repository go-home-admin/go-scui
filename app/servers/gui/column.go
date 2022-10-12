package gui

import (
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
	"github.com/go-home-admin/go-admin/generate/proto/common/grid"
)

type Column struct {
	*grid.Column
	*base.Render
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

func (c *Column) Filters(v []*grid.Filter) *Column {
	c.Column.Filters = v
	return c
}

func (c *Column) Template(v string) *Column {
	c.Column.Template = v
	return c
}
