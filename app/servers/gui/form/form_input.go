package form

import (
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
)

const input = `
<el-form-item label="__label__" prop="__prop__">
	<el-input v-model="__FORM__.__prop__" __options__ clearable></el-input>
</el-form-item>
`

// Input 普通组件
func (f *Form) Input(prop, label string) *InputFormItems {
	item := &InputFormItems{
		formID: f.GetID(),
		Render: base.NewRender(input),
		FormItems: &FormItems{
			Label:     label,
			Name:      prop,
			Component: "input",
			Options: map[string]interface{}{
				"placeholder": prop,
			},
		},
	}
	f.AddFormData(prop, "")
	f.AddItems(item)
	return item
}

type InputFormItems struct {
	formID string
	*base.Render
	*FormItems
}

func (i *InputFormItems) GetTemplate(pr ...base.RenderBase) string {
	i.Template = base.ReplaceAll(i.Template, []string{
		"__FORM__", i.formID + ".form",
		"__label__", i.FormItems.Label,
		"__prop__", i.FormItems.Name,
		"__options__", i.FormItems.GetOpt(),
	})
	return i.Template
}
