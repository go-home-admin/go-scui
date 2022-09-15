package public

import (
	gin "github.com/gin-gonic/gin"
	admin "github.com/go-home-admin/go-admin/generate/proto/admin"
	http "github.com/go-home-admin/home/app/http"
)

// Login   登陆
func (receiver *Controller) Login(req *admin.LoginRequest, ctx http.Context) (*admin.LoginResponse, error) {
	return &admin.LoginResponse{
		Token: "123",
		UserInfo: &admin.Userinfo{
			UserId:    "admin",
			UserName:  "admin",
			Dashboard: "admin",
			Role:      nil,
		},
	}, nil
}

// GinHandleLogin gin原始路由处理
// http.Post(/login)
func (receiver *Controller) GinHandleLogin(ctx *gin.Context) {
	req := &admin.LoginRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := receiver.Login(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
