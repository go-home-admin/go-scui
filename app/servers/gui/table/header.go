package table

import (
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
	"github.com/go-home-admin/go-admin/app/servers/gui/form"
	"github.com/go-home-admin/go-admin/app/servers/gui/html"
	"github.com/go-home-admin/home/app/http"
)

type Header struct {
	http.Context
	*base.Render
}

func NewHeader(ctx http.Context) *Header {
	return &Header{
		Context: ctx,
		Render:  base.NewRender(`<slot id="header"/>`),
	}
}

func (h *Header) add(r base.RenderBase) {
	h.AddRender(r, "header")
}

func (h *Header) Create(formRender *form.DialogForm) *DialogButton {
	formRender.OnSubmit(`function() {
		console.log(this.__ID__.form)
		alert("待请求")
	}`)
	dia := html.NewDialog().SetContext(formRender)
	h.AddRender(dia)

	button := NewButton(h, "创建")
	h.add(button)
	dialogButton := &DialogButton{Button: button}
	button.attr = append(button.attr, `@click="__ID__DialogOpen({})"`)
	dialogButton.AddMethods(`__ID__DialogOpen`, base.ReplaceAll(`function (row) {
		this.__DIALOG__ = true
		this.__FORM_ID__SetData(row)
	}
	`, []string{
		"__DIALOG__", dia.GetVisibleName(),
		"__FORM_ID__", formRender.GetID(),
	}))

	return dialogButton
}
