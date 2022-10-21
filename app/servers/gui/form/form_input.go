package form

import (
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
)

const input = `
<el-form-item label="__label__" prop="__prop__">
	<__component__ v-model="__FORM__.__prop__" __options__ clearable></__component__>
</el-form-item>
`

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
const inputRadio = `
<el-form-item  prop="__prop__">
	<__component__ v-model="__FORM__.__prop__" __options__ >
		__label__
	</__component__>
</el-form-item>
`

type InputFormItems struct {
	formID string
	*base.Render
	formItems *FormItems
}

type SelectOptionItem struct {
	Value interface{} `json:"value"`
	Label string      `json:"label"`
}

// Input 普通组件
func (f *Form) Input(prop, label string) *InputFormItems {
	item := &InputFormItems{
		formID: f.GetID(),
		Render: base.NewRender(input),
		formItems: &FormItems{
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
	return item
}

// Select 普通组件
func (f *Form) Select(prop, label string) *InputFormItems {
	item := &InputFormItems{
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
func (i *InputFormItems) AddSelectOptions(f *Form, optionsConfig []*SelectOptionItem) *InputFormItems {
	f.AddDependData(i.formItems.Name+"Options", optionsConfig)
	return i
}

// Switch 普通组件
func (f *Form) Switch(prop, label string) *InputFormItems {
	item := &InputFormItems{
		formID: f.GetID(),
		Render: base.NewRender(input),
		formItems: &FormItems{
			Label:     label,
			Name:      prop,
			Component: "el-switch",
			Options: map[string]interface{}{
				"active-color":   "#13ce66",
				"inactive-color": "#ff4949",
			},
		},
	}
	f.AddFormData(prop, "")
	f.AddItems(item)
	return item
}

// Radio 普通组件
func (f *Form) Radio(prop, label string) *InputFormItems {
	item := &InputFormItems{
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

// DatePicker 时间插件
func (f *Form) DatePicker(prop, label string) *InputFormItems {
	item := &InputFormItems{
		formID: f.GetID(),
		Render: base.NewRender(input),
		formItems: &FormItems{
			Label:     label,
			Name:      prop,
			Component: "el-date-picker",
			Options: map[string]interface{}{
				"placeholder": label,
			},
		},
	}
	f.AddFormData(prop, "")
	f.AddItems(item)
	return item
}

// ColorPicker 颜色选择器
func (f *Form) ColorPicker(prop, label string) *InputFormItems {
	item := &InputFormItems{
		formID: f.GetID(),
		Render: base.NewRender(input),
		formItems: &FormItems{
			Label:     label,
			Name:      prop,
			Component: "el-color-picker",
			Options:   map[string]interface{}{},
		},
	}
	f.AddFormData(prop, "#409EFF")
	f.AddItems(item)
	return item
}

func (i *InputFormItems) Options(key string, value string) *InputFormItems {
	i.formItems.Options[key] = value
	return i
}

func (i *InputFormItems) GetTemplate(pr ...base.RenderBase) string {
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
