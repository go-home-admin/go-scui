package form

import (
	"github.com/go-home-admin/go-admin/app/servers/gui"
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
)

type ColorFormItems struct {
	formID string
	*base.Render
	formItems *gui.FormItems
}

// ColorPicker 颜色选择器
func (f *Form) ColorPicker(prop, label string) *ColorFormItems {
	item := &ColorFormItems{
		formID: f.GetID(),
		Render: base.NewRender(input),
		formItems: &gui.FormItems{
			Label:     label,
			Name:      prop,
			Component: "el-color-picker",
			Options:   map[string]interface{}{},
		},
	}
	f.AddFormData(prop, "#409EFF")
	f.AddItems(item)
	return item
}

func (i *ColorFormItems) GetTemplate(pr ...base.RenderBase) string {
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
