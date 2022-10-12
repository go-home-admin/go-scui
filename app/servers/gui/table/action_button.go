package table

import (
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
	"github.com/go-home-admin/go-admin/app/servers/gui/form"
	"github.com/go-home-admin/home/app/http"
	"strings"
)

type Button struct {
	http.Context
	*base.Render

	attr []string
}

func NewButton(ctx http.Context, text string) *Button {
	return &Button{
		Context: ctx,
		Render:  base.NewRender(`<el-button __BUTTON__>` + text + `</el-button>`),
		attr:    make([]string, 0),
	}
}

// Confirm 确认后请求, url 中的 {{ row.id }} 会作为代码执行
func (b *Button) Confirm(url string) *ConfirmButton {
	funName, code := base.LoadJsFunction("button_confirm.js")
	funName = base.RenderID + funName
	b.attr = append(b.attr, "@click=\""+funName+"(row)\"")

	// 添加跳转函数, 内容固定, 连贯操作替换固定中文
	url = strings.ReplaceAll(url, "{{", "\"+")
	url = strings.ReplaceAll(url, "}}", "+\"")
	code = strings.ReplaceAll(code, "__url__", "\""+base.ToUrl(url)+"\"")
	b.AddMethods(funName, code)
	return &ConfirmButton{
		Button: b, funName: funName,
	}
}

func (b *Button) Edit(render *form.DialogForm) *Dialog {
	render.OnSubmit(`function() {
		console.log(this.__ID__.form)
		alert("待请求")
	}`)
	dialog := &Dialog{Button: b}
	b.attr = append(b.attr, `@click="__ID__DialogOpen(row)"`)
	dialog.AddData("__ID__visible", false)
	dialog.AddMethods(`__ID__DialogOpen`, strings.ReplaceAll(`function (row) {
		this.__ID__visible = true
		this.__FORM_ID__SetData(row)
	}
	`, "__FORM_ID__", render.GetID()))
	// 按钮操作的关联组件
	dia := base.NewRender(`<el-dialog v-model="` + b.GetID() + `visible" :width="500""><slot id="form"/></el-dialog>`)
	dia.AddRender(render, "form")
	b.AddRender(dia)

	return dialog
}

func (b *Button) Dialog(render base.RenderBase) *Dialog {
	dialog := &Dialog{
		Button: b,
	}

	b.attr = append(b.attr, `@click="__ID__DialogOpen(row)"`)
	dialog.AddData("__ID__visible", false)
	dialog.AddMethods(`__ID__DialogOpen`, `function (row) {
	this.__ID__visible = true
}
`)
	// 按钮操作的关联组件
	dia := base.NewRender(`<el-dialog v-model="` + b.GetID() + `visible" :width="500""><slot id="form"/></el-dialog>`)
	dia.AddRender(render.(base.RenderBase), "form")
	b.AddRender(dia)
	return dialog
}

// Dialog 弹窗
type Dialog struct {
	*Button
}

func (b *Button) GetTemplate(pr ...base.RenderBase) string {
	btnStr := ""
	for _, s := range b.attr {
		btnStr = btnStr + " " + s
	}
	b.Render.Template = strings.ReplaceAll(b.Render.Template, "__BUTTON__", btnStr)
	return b.RepID(b.Render.GetTemplate(pr...))
}

// ConfirmButton 确认按钮
type ConfirmButton struct {
	*Button
	funName string
}

func (b *ConfirmButton) Title(title string) *ConfirmButton {
	b.Button.MountedRender[b.funName] = strings.ReplaceAll(b.Button.MountedRender[b.funName], "敏感操作提示", title)
	return b
}

func (b *ConfirmButton) Text(text string) *ConfirmButton {
	b.Button.MountedRender[b.funName] = strings.ReplaceAll(b.Button.MountedRender[b.funName], "这是一个敏感的操作，是否继续?", text)
	return b
}
