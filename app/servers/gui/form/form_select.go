package form

import (
	"github.com/go-home-admin/go-admin/app/servers/gui"
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
	"strconv"
)

type SelectFormItems struct {
	formID string
	*base.Render
	formItems *gui.FormItems
}

type SelectOptionItem struct {
	Value interface{} `json:"value"`
	Label string      `json:"label"`
}

// Select 普通组件
func (f *Form) Select(prop, label string) *SelectFormItems {
	item := &SelectFormItems{
		formID:    f.GetID(),
		Render:    base.LoadVue("form/select.vue"),
		formItems: gui.NewItems(prop, label, "el-select"),
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

func (i *SelectFormItems) SelectOptions(list interface{}) {
	i.AddData("__ID__options", list)
}

func (i *SelectFormItems) Options(k, v string) *SelectFormItems {
	i.formItems.Options[k] = v
	return i
}

func (i *SelectFormItems) Span(v string) *SelectFormItems {
	i.AddRep("__SPAN__", v)
	return i
}
