package admin_user

import (
	gin "github.com/gin-gonic/gin"
	admin "github.com/go-home-admin/go-admin/generate/proto/admin"
	http "github.com/go-home-admin/home/app/http"
)

// UserAdd   用户创建
func (receiver *Controller) UserAdd(req *admin.UserAddRequest, ctx http.Context) (*admin.UserAddResponse, error) {
	// TODO 这里写业务
	return &admin.UserAddResponse{}, nil
}

// GinHandleUserAdd gin原始路由处理
// http.Post(/admin/user)
func (receiver *Controller) GinHandleUserAdd(ctx *gin.Context) {
	req := &admin.UserAddRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := receiver.UserAdd(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
