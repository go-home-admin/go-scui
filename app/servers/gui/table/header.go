package table

import (
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
)

type Header struct {
	Context GuiController
	*base.Render
}

func NewHeader(ctx GuiController) *Header {
	return &Header{
		Context: ctx,
		Render:  base.NewRender(`<slot id="header"/>`),
	}
}

// 在指定位置新增组件
func (h *Header) add(r base.RenderBase) {
	h.AddRender(r, "header")
}

// Create 创建按钮
func (h *Header) Create() *DialogButton {
	dia, formRender := GetEditDialog(h.Context, h)

	button := NewButton(h.Context, "创建")
	h.add(button)
	dialogButton := NewDialogButton(button)
	dialogButton.Click("__ID__DialogOpen({})")
	dialogButton.AddMethods(`__ID__DialogOpen`, `function (row) {
		this.__DIALOG__ = true
		this.__FORM_ID__SetData(row)
		this.__FORM_ID__.url = "`+base.ToUrl(h.Context.Gin().Request.URL.Path)+`/create"
	}`,
		"__DIALOG__", dia.GetVisibleName(),
		"__FORM_ID__", formRender.GetID(),
	)

	return dialogButton
}
