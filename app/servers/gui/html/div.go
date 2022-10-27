package html

import "github.com/go-home-admin/go-admin/app/servers/gui/base"

type Div struct {
	*base.Render
}

func NewDiv() *Div {
	d := &Div{Render: base.NewRender("<div></div>")}
	return d
}
