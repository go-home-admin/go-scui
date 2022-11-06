package form

import (
	"github.com/go-home-admin/go-admin/app/servers/gui"
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
	"strconv"
)

type FileFormItems struct {
	formID string
	*base.Render
	formItems *gui.FormItems
}

// File 文件上传
func (f *Form) File(prop, label string) *FileFormItems {
	item := &FileFormItems{
		formID:    f.GetID(),
		Render:    base.LoadVue("form/file.vue"),
		formItems: gui.NewItems(prop, label, "el-upload"),
	}
	f.AddFormData(prop, "")
	f.AddItems(item)
	f.FormItems = append(f.FormItems, item.formItems)
	item.AddRep("__SPAN__", strconv.Itoa(item.formItems.Span))
	item.AddRep("__LABEL__", label)
	item.AddRep("__PROP__", prop)
	return item
}

func (i *FileFormItems) Options(list interface{}) {
	i.AddData("__ID__options", list)
}

func (i *FileFormItems) Span(v string) *FileFormItems {
	i.AddRep("__SPAN__", v)
	return i
}
