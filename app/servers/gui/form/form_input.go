package form

import (
	"github.com/go-home-admin/go-admin/app/servers/gui"
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
	"strconv"
)

const input = `
<el-col :span="__SPAN__">
<el-form-item label="__label__" prop="__prop__">
	<__component__ v-model="__FORM__.__prop__" __options__ clearable></__component__>
</el-form-item>
</el-col>
`

type InputFormItems struct {
	formID string
	*base.Render
	formItems *gui.FormItems
}

// Input 普通组件
func (f *Form) Input(prop, label string) *InputFormItems {
	item := &InputFormItems{
		formID:    f.GetID(),
		Render:    base.NewRender(input),
		formItems: gui.NewItems(prop, label, "el-input"),
	}
	f.AddFormData(prop, "")
	f.AddItems(item)
	f.FormItems = append(f.FormItems, item.formItems)
	item.AddRep("__SPAN__", strconv.Itoa(item.formItems.Span))
	item.AddRep("__FORM__", item.formID+".form")
	item.AddRep("__component__", item.formItems.Component)
	item.AddRep("__label__", item.formItems.Label)
	item.AddRep("__prop__", item.formItems.Name)
	return item
}

func (i *InputFormItems) GetTemplate(pr ...base.RenderBase) string {
	i.Template = base.ReplaceAll(i.Template, []string{
		"__options__", i.formItems.GetOpt(),
	})
	return i.Render.GetTemplate(pr...)
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

func (i *InputFormItems) Span(v string) *InputFormItems {
	i.AddRep("__SPAN__", v)
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
