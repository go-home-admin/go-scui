package form

import (
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
)

const inputRadio = `
<el-form-item  prop="__prop__">
	<__component__ v-model="__FORM__.__prop__" __options__ >
		__label__
	</__component__>
</el-form-item>
`

type RadioPickerFormItems struct {
	formID string
	*base.Render
	formItems *FormItems
}

// Radio 普通组件
func (f *Form) Radio(prop, label string) *RadioPickerFormItems {
	item := &RadioPickerFormItems{
		formID: f.GetID(),
		Render: base.NewRender(inputRadio),
		formItems: &FormItems{
			Label:     label,
			Name:      prop,
			Component: "el-radio",
			Options:   map[string]interface{}{},
		},
	}
	f.AddFormData(prop, "")
	f.AddItems(item)
	return item
}

func (i *RadioPickerFormItems) Options(key string, value string) *RadioPickerFormItems {
	i.formItems.Options[key] = value
	return i
}

func (i *RadioPickerFormItems) GetTemplate(pr ...base.RenderBase) string {
	i.Template = base.ReplaceAll(i.Template, []string{
		"__FORM__", i.formID + ".form",
		"__depend_data__", i.formID + ".dependData",
		"__component__", i.formItems.Component,
		"__label__", i.formItems.Label,
		"__prop__", i.formItems.Name,
		"__options__", i.formItems.GetOpt(),
	})
	return i.Template
}
