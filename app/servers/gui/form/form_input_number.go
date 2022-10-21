package form

import (
	"fmt"
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
)

const inputNumber = `
<el-form-item label="__label__" prop="__prop__">
	<el-input-number v-model="__FORM__.__prop__" __options__ clearable />
</el-form-item>
`

type InputNumber struct {
	formID string
	*base.Render
	*FormItems
	*Form
}

// InputNumber 普通组件
func (f *Form) InputNumber(prop, label string) *InputNumber {
	item := &InputNumber{
		formID: f.GetID(),
		Render: base.NewRender(inputNumber),
		FormItems: &FormItems{
			Label:     label,
			Name:      prop,
			Component: "input-number",
			Options: map[string]interface{}{
				"placeholder": prop,
				":controls":   "true",
			},
		},
		Form: f,
	}
	f.AddFormData(prop, "")
	f.AddItems(item)
	return item
}

func (i *InputNumber) GetTemplate(pr ...base.RenderBase) string {
	i.Template = base.ReplaceAll(i.Template, []string{
		"__FORM__", i.formID + ".form",
		"__label__", i.FormItems.Label,
		"__prop__", i.FormItems.Name,
		"__options__", i.FormItems.GetOpt(),
	})
	return i.Template
}

func (i *InputNumber) Default(num int) *InputNumber {
	i.AddFormData(i.Name, num)
	return i
}

func (i *InputNumber) Min(num int) *InputNumber {
	i.Options[":min"] = fmt.Sprint(num)
	return i
}

func (i *InputNumber) Max(num int) *InputNumber {
	i.Options[":max"] = fmt.Sprint(num)
	return i
}

func (i *InputNumber) Step(num int) *InputNumber {
	i.Options[":step"] = fmt.Sprint(num)
	return i
}

func (i *InputNumber) Controls(b bool) *InputNumber {
	i.Options[":controls"] = fmt.Sprint(b)
	return i
}
