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
	b.attr = append(b.attr, "@click=\""+funName+"\"")

	// 添加跳转函数, 内容固定, 连贯操作替换固定中文
	url = strings.ReplaceAll(url, "{{", "\"+")
	url = strings.ReplaceAll(url, "}}", "+\"")
	code = strings.ReplaceAll(code, "__url__", "\""+ToUrl(url)+"\"")
	b.AddMethods(funName, code)
	return &ConfirmButton{
		Button: b, funName: funName,
	}
}

func (b *Button) GetTemplate() string {
	btnStr := ""
	for _, s := range b.attr {
		btnStr = btnStr + " " + s
	}
	b.Render.template = strings.ReplaceAll(b.Render.template, "__BUTTON__", btnStr)
	return b.repID(b.Render.GetTemplate())
}

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
