package system

import (
	gin "github.com/gin-gonic/gin"
	admin "github.com/go-home-admin/go-admin/generate/proto/admin"
	http "github.com/go-home-admin/home/app/http"
)

// MenuList   菜单列表
func (receiver *Controller) MenuList(req *admin.MenuListRequest, ctx http.Context) (*admin.MenuListResponse, error) {
	// TODO 这里写业务
	return &admin.MenuListResponse{}, nil
}

// GinHandleMenuList gin原始路由处理
// http.Get(/system/menu/list)
func (receiver *Controller) GinHandleMenuList(ctx *gin.Context) {
	req := &admin.MenuListRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := receiver.MenuList(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
