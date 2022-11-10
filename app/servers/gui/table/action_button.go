package table

import (
	"github.com/go-home-admin/go-admin/app/servers/gui"
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
	"github.com/go-home-admin/go-admin/app/servers/gui/html"
	"strings"
)

type Button struct {
	Context GuiController
	*base.Render

	attr []string
}

func NewButton(ctx GuiController, text string) *Button {
	return &Button{
		Context: ctx,
		Render:  base.NewRender(`<el-button __BUTTON__>` + text + `</el-button>`),
		attr:    make([]string, 0),
	}
}

// Click 设置点击时间函数
func (b *Button) Click(fun string) *Button {
	b.attr = append(b.attr, `@click="`+fun+`"`)
	return b
}

// Confirm 确认后请求, url 中的 {{ row.id }} 会作为代码执行
func (b *Button) Confirm(url string) *ConfirmButton {
	funName, code := base.LoadJsFunction("button_confirm.js")
	funName = base.RenderID + funName
	b.attr = append(b.attr, "@click=\""+funName+"(row)\"")

	// 添加跳转函数, 内容固定, 连贯操作替换固定中文
	url = base.ReplaceAll(url, []string{
		"{{", "\"+",
		"}}", "+\"",
	})
	code = strings.ReplaceAll(code, "__url__", "\""+base.ToUrl(url)+"\"")
	b.AddMethods(funName, code)
	return &ConfirmButton{
		Button: b, funName: funName,
	}
}

func (b *Button) Delete() *ConfirmButton {
	url := b.Context.Gin().Request.URL.Path + "/del"

	if c, ok := b.Context.(gui.GetPrimary); ok {
		return b.Confirm(url + "?id={{ row." + c.GetPrimary() + " }}")
	}

	return b.Confirm(url + "?id={{ row.id }}")
}

func (b *Button) Edit() *DialogButton {
	dia, formRender := GetEditDialog(b.Context, b)

	dialogButton := &DialogButton{Button: b}
	b.attr = append(b.attr, `@click="__ID__DialogOpen(row)"`)
	dialogButton.AddMethods(`__ID__DialogOpen`, `function (row) {
		this.__DIALOG__ = true
		this.__FORM_ID__SetData(row)
		this.__FORM_ID__.url = "`+base.ToUrl(b.Context.Gin().Request.URL.Path)+`/edit"
	}
	`,
		"__DIALOG__", dia.GetVisibleName(),
		"__FORM_ID__", formRender.GetID(),
	)

	return dialogButton
}

func (b *Button) Options(k, v string) {
	str := k + "=" + "'" + v + "'"
	b.attr = append(b.attr, str)
}

func (b *Button) Dialog(render base.RenderBase) *DialogButton {
	dialog := &DialogButton{
		Button: b,
	}

	b.attr = append(b.attr, `@click="__ID__DialogOpen(row)"`)
	dialog.AddData("__ID__visible", false)
	dialog.AddMethods(`__ID__DialogOpen`, `function (row) {
	this.__ID__visible = true
}
`)
	// 按钮操作的关联组件
	dia := html.NewDialog()
	dia.AddRender(render.(base.RenderBase), "context")
	b.AddRender(dia)
	return dialog
}

// DialogButton 按钮弹窗
type DialogButton struct {
	*Button
}

func NewDialogButton(b *Button) *DialogButton {
	return &DialogButton{Button: b}
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
