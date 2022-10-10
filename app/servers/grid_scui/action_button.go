package grid_scui

import "strings"

type Button struct {
	*Render

	attr []string
}

func NewButton(text string) *Button {
	return &Button{
		Render: NewRender(`<el-button __BUTTON__>` + text + `</el-button>`),
		attr:   make([]string, 0),
	}
}

// Confirm 确认后请求, url 中的 {{ row.id }} 会作为代码执行
func (b *Button) Confirm(url string) *ConfirmButton {
	funName, code := loadJsFunction("button_confirm.js")
	funName = RenderID + funName
	b.attr = append(b.attr, "@click=\""+funName+"(row)\"")

	// 添加跳转函数, 内容固定, 连贯操作替换固定中文
	url = strings.ReplaceAll(url, "{{", "\"+")
	url = strings.ReplaceAll(url, "}}", "+\"")
	code = strings.ReplaceAll(code, "__url__", "\""+ToUrl(url)+"\"")
	b.AddMethods(funName, code)
	return &ConfirmButton{
		Button: b, funName: funName,
	}
}

func (b *Button) Dialog(render RenderBase) *Dialog {
	dialog := &Dialog{
		Button: b,
	}

	b.attr = append(b.attr, `@click="__ID__DialogOpen"`)
	dialog.AddData("__ID__visible", false)
	dialog.AddMethods(`__ID__DialogOpen`, `
function () {
	this.__ID__visible = true
}
`)
	// 按钮操作的关联组件
	dia := NewRender(`<el-dialog v-model="` + b.id + `visible" :width="500""><slot id="form"/></el-dialog>`)
	dia.AddRender(render.(RenderBase), "form")
	b.AddRender(dia)
	return dialog
}

// Dialog 弹窗
type Dialog struct {
	*Button
}

func (b *Button) GetTemplate(pr ...RenderBase) string {
	btnStr := ""
	for _, s := range b.attr {
		btnStr = btnStr + " " + s
	}
	b.Render.template = strings.ReplaceAll(b.Render.template, "__BUTTON__", btnStr)
	return b.repID(b.Render.GetTemplate(pr...))
}

// ConfirmButton 确认按钮
type ConfirmButton struct {
	*Button
	funName string
}

func (b *ConfirmButton) Title(title string) *ConfirmButton {
	b.Button.mountedRender[b.funName] = strings.ReplaceAll(b.Button.mountedRender[b.funName], "敏感操作提示", title)
	return b
}

func (b *ConfirmButton) Text(text string) *ConfirmButton {
	b.Button.mountedRender[b.funName] = strings.ReplaceAll(b.Button.mountedRender[b.funName], "这是一个敏感的操作，是否继续?", text)
	return b
}
