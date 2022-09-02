package system

import (
	gin "github.com/gin-gonic/gin"
	admin "github.com/go-home-admin/go-admin/generate/proto/admin"
	http "github.com/go-home-admin/home/app/http"
)

// MenuSave   菜单保存
func (receiver *Controller) MenuSave(req *admin.MenuSaveRequest, ctx http.Context) (*admin.MenuSaveResponse, error) {
	// TODO 这里写业务
	return &admin.MenuSaveResponse{}, nil
}

// GinHandleMenuSave gin原始路由处理
// http.Post(/system/menu/save)
func (receiver *Controller) GinHandleMenuSave(ctx *gin.Context) {
	req := &admin.MenuSaveRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := receiver.MenuSave(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
