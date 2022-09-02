package system

import (
	gin "github.com/gin-gonic/gin"
	admin "github.com/go-home-admin/go-admin/generate/proto/admin"
	http "github.com/go-home-admin/home/app/http"
)

// MenuDel   菜单删除
func (receiver *Controller) MenuDel(req *admin.MenuDelRequest, ctx http.Context) (*admin.MenuDelResponse, error) {
	// TODO 这里写业务
	return &admin.MenuDelResponse{}, nil
}

// GinHandleMenuDel gin原始路由处理
// http.Delete(/system/menu/del)
func (receiver *Controller) GinHandleMenuDel(ctx *gin.Context) {
	req := &admin.MenuDelRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := receiver.MenuDel(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
