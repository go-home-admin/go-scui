package form

import (
	"github.com/go-home-admin/go-admin/app/servers/gui"
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
	"strconv"
)

type DatePickerFormItems struct {
	formID string
	*base.Render
	formItems *gui.FormItems
}

// DatePicker 时间插件
func (f *Form) DatePicker(prop, label string) *DatePickerFormItems {
	item := &DatePickerFormItems{
		formID:    f.GetID(),
		Render:    base.LoadVue("form/date_picker.vue"),
		formItems: gui.NewItems(prop, label, "el-date-picker"),
	}
	f.AddFormData(prop, "")
	f.AddItems(item)
	f.FormItems = append(f.FormItems, item.formItems)
	item.AddRep("__SPAN__", strconv.Itoa(item.formItems.Span))
	item.AddRep("__FORM__", f.GetID())
	item.AddRep("__PROP__", prop)
	item.AddRep("__LABEL__", label)
	item.AddRep("__KV__", "")
	return item
}

func (i *DatePickerFormItems) Ymd() *DatePickerFormItems {
	i.AddAttr("format", "YYYY-MM-DD HH:mm:ss")
	i.AddAttr("value-format", "YYYY-MM-DD HH:mm:ss")
	return i
}

// AddAttr 拼接成 key=""
func (i *DatePickerFormItems) AddAttr(k, v string) *DatePickerFormItems {
	i.AddRep("__KV__", k, v)

	return i
}

func (i *DatePickerFormItems) Span(v string) *DatePickerFormItems {
	i.AddRep("__SPAN__", v)
	return i
}
