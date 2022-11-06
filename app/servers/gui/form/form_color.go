package form

import (
	"github.com/go-home-admin/go-admin/app/servers/gui"
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
	"strconv"
)

type ColorFormItems struct {
	formID string
	*base.Render
	formItems *gui.FormItems
}

// ColorPicker 颜色选择器
func (f *Form) ColorPicker(prop, label string) *ColorFormItems {
	item := &ColorFormItems{
		formID:    f.GetID(),
		Render:    base.NewRender(input),
		formItems: gui.NewItems(prop, label, "el-color-picker"),
	}
	f.AddFormData(prop, "#409EFF")
	f.AddItems(item)
	return item
}

func (i *ColorFormItems) GetTemplate(pr ...base.RenderBase) string {
	i.AddRep("__SPAN__", strconv.Itoa(i.formItems.Span))
	i.AddRep("__FORM__", i.formID+".form")
	i.AddRep("__depend_data__", i.formID+".dependData")
	i.AddRep("__component__", i.formItems.Component)
	i.AddRep("__label__", i.formItems.Label)
	i.AddRep("__prop__", i.formItems.Name)
	i.AddRep("__options__", i.formItems.GetOpt())
	return i.Render.GetTemplate(pr...)
}

func (i *ColorFormItems) Span(v string) *ColorFormItems {
	i.AddRep("__SPAN__", v)
	return i
}
