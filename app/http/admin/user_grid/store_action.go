package user_grid

import (
	gin "github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/generate/proto/common/grid"
	http "github.com/go-home-admin/home/app/http"
)

// Store   提交创建的数据
func (receiver *Controller) Store(req *grid.StoreRequest, ctx http.Context) (*grid.StoreResponse, error) {
	// TODO 这里写业务
	return &grid.StoreResponse{}, nil
}

// GinHandleStore gin原始路由处理
// http.Post(/user/create)
func (receiver *Controller) GinHandleStore(ctx *gin.Context) {
	req := &grid.StoreRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := receiver.Store(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
