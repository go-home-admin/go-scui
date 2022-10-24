package form

import (
	"github.com/go-home-admin/go-admin/app/servers/gui"
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
)

type DatePickerFormItems struct {
	formID string
	*base.Render
	formItems *gui.FormItems
}

// DatePicker 时间插件
func (f *Form) DatePicker(prop, label string) *DatePickerFormItems {
	item := &DatePickerFormItems{
		formID: f.GetID(),
		Render: base.NewRender(input),
		formItems: &gui.FormItems{
			Label:     label,
			Name:      prop,
			Component: "el-date-picker",
			Options: map[string]interface{}{
				"placeholder": label,
			},
		},
	}
	f.AddFormData(prop, "")
	f.AddItems(item)
	f.FormItems = append(f.FormItems, item.formItems)
	return item
}

func (i *DatePickerFormItems) Options(key string, value string) *DatePickerFormItems {
	i.formItems.Options[key] = value
	return i
}

func (i *DatePickerFormItems) YmdHis() *DatePickerFormItems {
	i.formItems.Options["format"] = "YYYY-MM-DD HH:mm:ss"
	i.formItems.Options["value-format"] = "YYYY-MM-DD HH:mm:ss"
	return i
}

func (i *DatePickerFormItems) GetTemplate(pr ...base.RenderBase) string {
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
