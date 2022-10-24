package form

import (
	"github.com/go-home-admin/go-admin/app/servers/gui"
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
)

type Form struct {
	*base.View

	LabelWidthOpt string           `json:"labelWidth"`
	LabelPosition string           `json:"labelPosition"`
	Size          string           `json:"size"`
	FormItems     []*gui.FormItems `json:"formItems"`

	rules    map[string]interface{}
	formData map[string]interface{}
}

func (f *Form) GetFormItems() []*gui.FormItems {
	return f.FormItems
}

func (f *Form) LabelWidth(v string) {
	f.AddRep("__FORM_OPT__", "label-width", v)
}

func (f *Form) AddItems(base base.RenderBase) {
	f.AddRender(base, "form-item")
}

func (f *Form) AddFormData(k string, v interface{}) {
	f.AddData("__ID__.form."+k, v)
}

// OnSubmit 提交函数
func (f *Form) OnSubmit(v string) {
	f.AddMethods("__ID__onSubmit", v)
}

type DialogForm struct {
	*Form
}

func NewForm() *DialogForm {
	return &DialogForm{
		Form: &Form{
			View:          base.NewView("form.vue"),
			LabelPosition: "right",
			FormItems:     make([]*gui.FormItems, 0),
			rules:         map[string]interface{}{},
			formData:      map[string]interface{}{},
		},
	}
}

func NewTableForm() *Form {
	view := base.NewView("table_form.vue")
	return &Form{
		View:          view,
		LabelPosition: "right",
	}
}
