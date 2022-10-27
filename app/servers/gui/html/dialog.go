package html

import "github.com/go-home-admin/go-admin/app/servers/gui/base"

type Dialog struct {
	*base.Render
}

func NewDialog() *Dialog {
	d := &Dialog{
		Render: base.LoadVue("dialog.vue"),
	}
	d.Width("50%")
	return d
}

func (d *Dialog) GetVisibleName() string {
	return d.RepID("__ID__visible")
}

func (d *Dialog) Style(k, v string) {
	d.AddRep("__OPT__", "style", k, v)
}

func (d *Dialog) Width(v string) {
	d.AddRep("__OPT__", "width", v)
}
