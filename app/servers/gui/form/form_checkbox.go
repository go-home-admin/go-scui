package form

import (
	"github.com/go-home-admin/go-admin/app/servers/gui"
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
)

type CheckBoxPickerFormItems struct {
	formID string
	*base.Render
	formItems *gui.FormItems
}

// CheckBox 普通组件
func (f *Form) CheckBox(prop, label string) *CheckBoxPickerFormItems {
	item := &CheckBoxPickerFormItems{
		formID: f.GetID(),
		Render: base.LoadVue("form/checkbox.vue"),
		formItems: &gui.FormItems{
			Label:     label,
			Name:      prop,
			Component: "el-checkbox-group",
			Options:   map[string]interface{}{},
		},
	}
	f.AddFormData(prop, "")
	f.AddItems(item)
	f.FormItems = append(f.FormItems, item.formItems)
	item.AddRep("__FORM__", item.formID)
	item.AddRep("__LABEL__", label)
	item.AddRep("__PROP__", prop)
	item.AddRep("__options__", item.formItems.GetOpt())
	return item
}

func (i *CheckBoxPickerFormItems) Options(list interface{}) {
	i.AddData("__ID__options", list)
}
