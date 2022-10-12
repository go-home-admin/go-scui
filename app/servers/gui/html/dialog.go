package html

import "github.com/go-home-admin/go-admin/app/servers/gui/base"

type Dialog struct {
	*base.Render
}

func NewDialog() *Dialog {
	return &Dialog{
		Render: base.LoadVue("dialog.vue"),
	}
}

func (d *Dialog) GetVisibleName() string {
	return d.RepID("__ID__visible")
}

func (d *Dialog) SetContext(b base.RenderBase) *Dialog {
	d.AddRender(b, "context")
	return d
}
