package table

import (
	"github.com/go-home-admin/go-admin/app/servers/gui"
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
)

type Column struct {
	*gui.Column
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

func (c *Column) Filters(v []*gui.Filter) *Column {
	c.Column.Filters = v
	return c
}

func (c *Column) Template(v string) *Column {
	c.Column.Template = v
	return c
}
