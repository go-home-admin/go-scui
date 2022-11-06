package form

import (
	"github.com/go-home-admin/go-admin/app/servers/gui"
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
	"strconv"
)

type DateTimePickerFormItems struct {
	formID string
	*base.Render
	formItems *gui.FormItems
}

// DateTimePicker 时间插件
func (f *Form) DateTimePicker(prop, label string) *DateTimePickerFormItems {
	item := &DateTimePickerFormItems{
		formID:    f.GetID(),
		Render:    base.LoadVue("form/date_time_picker.vue"),
		formItems: gui.NewItems(prop, label, "el-date-picker"),
	}
	f.AddFormData(prop, "")
	f.AddItems(item)
	f.FormItems = append(f.FormItems, item.formItems)
	item.AddRep("__SPAN__", strconv.Itoa(item.formItems.Span))
	item.AddRep("__FORM__", f.GetID())
	item.AddRep("__PROP__", prop)
	item.AddRep("__LABEL__", label)
	item.AddAttr("value-format", "YYYY-MM-DD HH:mm:ss")
	return item
}

func (f *Form) DateTime(prop, label string) *DateTimePickerFormItems {
	return f.DateTimePicker(prop, label).YmdHis()
}

func (i *DateTimePickerFormItems) AddAttr(k, v string) *DateTimePickerFormItems {
	i.AddRep("__KV__", k, v)

	return i
}

func (i *DateTimePickerFormItems) YmdHis() *DateTimePickerFormItems {
	i.AddAttr("format", "YYYY-MM-DD HH:mm:ss")
	i.AddAttr("value-format", "YYYY-MM-DD HH:mm:ss")
	return i
}

func (i *DateTimePickerFormItems) Ymd() *DateTimePickerFormItems {
	i.AddAttr("format", "YYYY-MM-DD")
	i.AddAttr("value-format", "YYYY-MM-DD")
	return i
}

func (i *DateTimePickerFormItems) Span(v string) *DateTimePickerFormItems {
	i.AddRep("__SPAN__", v)
	return i
}
