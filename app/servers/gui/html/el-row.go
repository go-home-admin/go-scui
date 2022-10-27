package html

import "github.com/go-home-admin/go-admin/app/servers/gui/base"

type ElRow struct {
	*base.Render
}

func NewElRow() *ElRow {
	e := &ElRow{Render: base.NewRender()}
	return e
}
