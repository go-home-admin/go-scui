package form

import (
	"github.com/go-home-admin/go-admin/app/servers/gui"
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
	"strconv"
)

type CheckBoxPickerFormItems struct {
	formID string
	*base.Render
	formItems *gui.FormItems
}

// CheckBox 普通组件
func (f *Form) CheckBox(prop, label string) *CheckBoxPickerFormItems {
	item := &CheckBoxPickerFormItems{
		formID:    f.GetID(),
		Render:    base.LoadVue("form/checkbox.vue"),
		formItems: gui.NewItems(prop, label, "el-checkbox-group"),
	}
	f.AddFormData(prop, "")
	f.AddItems(item)
	f.FormItems = append(f.FormItems, item.formItems)
	item.AddRep("__SPAN__", strconv.Itoa(item.formItems.Span))
	item.AddRep("__FORM__", item.formID)
	item.AddRep("__LABEL__", label)
	item.AddRep("__PROP__", prop)
	item.AddRep("__options__", item.formItems.GetOpt())
	return item
}

func (i *CheckBoxPickerFormItems) Options(list interface{}) {
	i.AddData("__ID__options", list)
}

func (i *CheckBoxPickerFormItems) Span(v string) *CheckBoxPickerFormItems {
	i.AddRep("__SPAN__", v)
	return i
}
