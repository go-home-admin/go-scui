package form

import (
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
)

const inputSelect = `
<el-form-item label="__label__" prop="__prop__">
	<__component__ v-model="__FORM__.__prop__" __options__ clearable>
	   <el-option
		  v-for="item in __depend_data__.__prop__Options"
		  :key="item.value"
		  :label="item.label"
		  :value="item.value">
		</el-option>
	</__component__>
</el-form-item>
`

type SelectFormItems struct {
	formID string
	*base.Render
	formItems *FormItems
}

type SelectOptionItem struct {
	Value interface{} `json:"value"`
	Label string      `json:"label"`
}

// Select 普通组件
func (f *Form) Select(prop, label string) *SelectFormItems {
	item := &SelectFormItems{
		formID: f.GetID(),
		Render: base.NewRender(inputSelect),
		formItems: &FormItems{
			Label:     label,
			Name:      prop,
			Component: "el-select",
			Options: map[string]interface{}{
				"placeholder": label,
			},
		},
	}
	f.AddDependData(prop+"Options", make([]*SelectOptionItem, 0))
	f.AddFormData(prop, "")
	f.AddItems(item)
	return item
}

//AddSelectOptions 初始化select选择项
func (i *SelectFormItems) AddSelectOptions(f *Form, optionsConfig []*SelectOptionItem) *SelectFormItems {
	f.AddDependData(i.formItems.Name+"Options", optionsConfig)
	return i
}

func (i *SelectFormItems) GetTemplate(pr ...base.RenderBase) string {
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
