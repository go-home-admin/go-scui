package form

import (
	"fmt"
	"github.com/go-home-admin/go-admin/app/servers/gui"
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
	"strconv"
)

const inputNumber = `
<el-col span="__SPAN__">
<el-form-item label="__label__" prop="__prop__">
	<el-input-number v-model="__FORM__.__prop__" __options__ clearable />
</el-form-item>
</el-col>
`

type InputNumber struct {
	formID string
	*base.Render
	*gui.FormItems
	*Form
}

// InputNumber 普通组件
func (f *Form) InputNumber(prop, label string) *InputNumber {
	item := &InputNumber{
		formID: f.GetID(),
		Render: base.NewRender(inputNumber),
		FormItems: &gui.FormItems{
			Label:     label,
			Name:      prop,
			Span:      24,
			Component: "input-number",
			Options: map[string]interface{}{
				"placeholder": prop,
				":controls":   "true",
			},
		},
		Form: f,
	}
	item.AddRep("__SPAN__", strconv.Itoa(item.FormItems.Span))
	f.AddFormData(prop, "")
	f.AddItems(item)
	f.FormItems = append(f.FormItems, item.FormItems)
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

func (i *InputNumber) Span(v string) *InputNumber {
	i.AddRep("__SPAN__", v)
	return i
}
