package form

import (
	"github.com/go-home-admin/go-admin/app/servers/gui"
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
	"strconv"
)

type SwitchFormItems struct {
	formID string
	*base.Render
	formItems *gui.FormItems
}

// Switch 普通组件
func (f *Form) Switch(prop, label string) *SwitchFormItems {
	item := &SwitchFormItems{
		formID:    f.GetID(),
		Render:    base.NewRender(input),
		formItems: gui.NewItems(prop, label, "el-switch"),
	}
	f.AddFormData(prop, "")
	f.AddItems(item)
	item.AddRep("__SPAN__", strconv.Itoa(item.formItems.Span))
	return item
}

func (i *SwitchFormItems) GetTemplate(pr ...base.RenderBase) string {
	i.Template = base.ReplaceAll(i.Template, []string{
		"__FORM__", i.formID + ".form",
		"__depend_data__", i.formID + ".dependData",
		"__component__", i.formItems.Component,
		"__label__", i.formItems.Label,
		"__prop__", i.formItems.Name,
		"__options__", i.formItems.GetOpt(),
	})
	return i.Template
}

func (i *SwitchFormItems) Span(v string) *SwitchFormItems {
	i.AddRep("__SPAN__", v)
	return i
}
