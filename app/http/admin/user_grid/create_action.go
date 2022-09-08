package user_grid

import (
	gin "github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/generate/proto/common/grid"
	http "github.com/go-home-admin/home/app/http"
)

// Create   创建数据的页面
func (receiver *Controller) Create(req *grid.CreateRequest, ctx http.Context) (*grid.CreateResponse, error) {
	// TODO 这里写业务
	return &grid.CreateResponse{}, nil
}

// GinHandleCreate gin原始路由处理
// http.Get(/user/create)
func (receiver *Controller) GinHandleCreate(ctx *gin.Context) {
	req := &grid.CreateRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := receiver.Create(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
