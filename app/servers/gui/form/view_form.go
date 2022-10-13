package form

import (
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
)

type Form struct {
	*base.View

	LabelWidth    string       `json:"labelWidth"`
	LabelPosition string       `json:"labelPosition"`
	Size          string       `json:"size"`
	FormItems     []*FormItems `json:"formItems"`

	rules    map[string]interface{}
	formData map[string]interface{}
}

func (f *Form) AddItems(base base.RenderBase) {
	f.AddRender(base, "form-item")
}

func (f *Form) AddFormData(k string, v interface{}) {
	f.formData[k] = v
}

// OnSubmit 提交函数
func (f *Form) OnSubmit(v string) {
	f.AddMethods("__ID__onSubmit", v)
}

type FormItems struct {
	// 对应的数据库字段
	Field string `json:"-"`

	// 输出到前端的字段
	Label          string                 `json:"label,omitempty"`
	Name           string                 `json:"name,omitempty"`
	Value          string                 `json:"value,omitempty"`
	Component      string                 `json:"component,omitempty"`
	Span           int                    `json:"span"`
	Options        map[string]interface{} `json:"options,omitempty"`
	Rules          map[string]string      `json:"rules,omitempty"`
	RequiredHandle string                 `json:"required_handle,omitempty"`
	HideHandle     string                 `json:"hide_handle,omitempty"`
}

func (i *FormItems) GetOpt() string {
	s := ""
	for s2, i2 := range i.Options {
		switch i2.(type) {
		case string:
			s = s + " " + s2 + `="` + i2.(string) + `"`
		}
	}

	return s
}

type DialogForm struct {
	*Form
}

func NewForm() *DialogForm {
	return &DialogForm{
		Form: &Form{
			View:          base.NewView("form.vue"),
			LabelPosition: "right",
			FormItems:     make([]*FormItems, 0),
			rules:         map[string]interface{}{},
			formData:      map[string]interface{}{},
		},
	}
}

func (f *DialogForm) GetData() map[string]interface{} {
	data := map[string]interface{}{
		"form":    f.Form.formData,
		"rules":   false,
		"loading": false,
	}

	return map[string]interface{}{
		f.GetID(): data,
	}
}

func NewTableForm() *Form {
	view := base.NewView("table_form.vue")
	return &Form{
		View:          view,
		LabelPosition: "right",
	}
}
