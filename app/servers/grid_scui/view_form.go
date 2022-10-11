package grid_scui

type Form struct {
	*View

	LabelWidth    string       `json:"labelWidth"`
	LabelPosition string       `json:"labelPosition"`
	Size          string       `json:"size"`
	FormItems     []*FormItems `json:"formItems"`
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

type DialogForm struct {
	*Form

	formData map[string]interface{}
}

func NewForm() *DialogForm {
	view := NewView("form.vue")

	view.AddMethods("__ID__onSubmit", `async function(){
		alert("你点击了提交")
	}`)

	return &DialogForm{
		Form: &Form{
			View:          view,
			LabelWidth:    "130px",
			LabelPosition: "right",
		},
		formData: map[string]interface{}{},
	}
}

func (f *DialogForm) GetData() map[string]interface{} {
	data := map[string]interface{}{
		"formItems": f.FormItems,
		"form":      f.formData,
		"loading":   false,
	}

	return map[string]interface{}{
		f.GetID(): data,
	}
}

// Input 普通组件
func (f *Form) Input(prop, label string) *InputFormItems {
	item := &InputFormItems{
		FormItems: &FormItems{
			Label:     label,
			Name:      prop,
			Component: "input",
			Options: map[string]interface{}{
				"placeholder": prop,
			},
		},
	}
	f.FormItems = append(f.FormItems, item.FormItems)

	return item
}

type InputFormItems struct {
	*FormItems
}

func NewTableForm() *Form {
	view := NewView("table_form.vue")
	return &Form{
		View:          view,
		LabelWidth:    "130px",
		LabelPosition: "right",
	}
}
