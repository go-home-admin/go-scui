package form

import (
	"github.com/go-home-admin/go-admin/app/servers/gui"
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
	"github.com/go-home-admin/go-admin/app/servers/gui/html"
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

// AddItems 添加渲染
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

func (f *Form) Style(k, v string) {
	f.AddRep("__FORM_OPT__", "style", k, v)
}

func (f *Form) Width(v string) {
	f.Style("width", v)
}

type DialogForm struct {
	*Form
	Dialog *html.Dialog
}

func (f *DialogForm) Width(v string) *DialogForm {
	f.Dialog.Width(v)
	return f
}

func (f *DialogForm) GetTemplate(pr ...base.RenderBase) string {
	return f.Dialog.GetTemplate(pr...)
}

func NewForm() *Form {
	form := &Form{
		View:          base.NewView("form.vue"),
		LabelPosition: "right",
		FormItems:     make([]*gui.FormItems, 0),
		rules:         map[string]interface{}{},
		formData:      map[string]interface{}{},
	}
	form.Width("100%")
	return form
}

func NewDialogForm() *DialogForm {
	d := &DialogForm{
		Form:   NewForm(),
		Dialog: html.NewDialog(),
	}
	d.Dialog.AddRender(d.Form, "context")
	return d
}

func NewTableForm() *Form {
	form := NewForm()
	return form
}
