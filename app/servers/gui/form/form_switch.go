package form

import (
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
)

type SwitchFormItems struct {
	formID string
	*base.Render
	formItems *FormItems
}

// Switch 普通组件
func (f *Form) Switch(prop, label string) *SwitchFormItems {
	item := &SwitchFormItems{
		formID: f.GetID(),
		Render: base.NewRender(input),
		formItems: &FormItems{
			Label:     label,
			Name:      prop,
			Component: "el-switch",
			Options: map[string]interface{}{
				"active-color":   "#13ce66",
				"inactive-color": "#ff4949",
			},
		},
	}
	f.AddFormData(prop, "")
	f.AddItems(item)
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
