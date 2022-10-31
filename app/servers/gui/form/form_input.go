package form

import (
	"github.com/go-home-admin/go-admin/app/servers/gui"
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
)

const input = `
<el-form-item label="__label__" prop="__prop__">
	<__component__ v-model="__FORM__.__prop__" __options__ clearable></__component__>
</el-form-item>
`

type InputFormItems struct {
	formID string
	*base.Render
	formItems *gui.FormItems
}

// Input 普通组件
func (f *Form) Input(prop, label string) *InputFormItems {
	item := &InputFormItems{
		formID: f.GetID(),
		Render: base.NewRender(input),
		formItems: &gui.FormItems{
			Label:     label,
			Name:      prop,
			Component: "el-input",
			Options: map[string]interface{}{
				"placeholder": label,
			},
		},
	}
	f.AddFormData(prop, "")
	f.AddItems(item)
	f.FormItems = append(f.FormItems, item.formItems)
	item.AddRep("__options__", item.formItems.GetOpt())
	return item
}

func (i *InputFormItems) GetTemplate(pr ...base.RenderBase) string {
	i.Template = base.ReplaceAll(i.Template, []string{
		"__FORM__", i.formID + ".form",
		"__component__", i.formItems.Component,
		"__label__", i.formItems.Label,
		"__prop__", i.formItems.Name,
		"__options__", i.formItems.GetOpt(),
	})
	return i.Template
}

func (i *InputFormItems) SaveToInt() *InputFormItems {
	i.formItems.SaveToInt()
	return i
}
func (i *InputFormItems) SaveToString() *InputFormItems {
	i.formItems.SaveToString()
	return i
}

func (i *InputFormItems) Options(k, v string) *InputFormItems {
	i.formItems.Options[k] = v
	return i
}

func (i *InputFormItems) Rules(b bool) *InputFormItems {
	if b {
		i.formItems.Rules = map[string]string{
			i.formItems.Name: "[{ required: true, message: '请填写内容' }]",
		}
	}
	return i
}
