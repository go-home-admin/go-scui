package admin_user

import (
	gin "github.com/gin-gonic/gin"
	admin "github.com/go-home-admin/go-admin/generate/proto/admin"
	http "github.com/go-home-admin/home/app/http"
)

// UserDel   用户删除
func (receiver *Controller) UserDel(req *admin.UserDelRequest, ctx http.Context) (*admin.UserDelResponse, error) {
	// TODO 这里写业务
	return &admin.UserDelResponse{}, nil
}

// GinHandleUserDel gin原始路由处理
// http.Delete(/admin/user)
func (receiver *Controller) GinHandleUserDel(ctx *gin.Context) {
	req := &admin.UserDelRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := receiver.UserDel(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
