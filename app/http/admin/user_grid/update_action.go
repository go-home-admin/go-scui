package user_grid

import (
	gin "github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/generate/proto/common/grid"
	http "github.com/go-home-admin/home/app/http"
)

// Update   提交编辑的数据
func (receiver *Controller) Update(req *grid.UpdateRequest, ctx http.Context) (*grid.UpdateResponse, error) {
	// TODO 这里写业务
	return &grid.UpdateResponse{}, nil
}

// GinHandleUpdate gin原始路由处理
// http.Put(/user/edit)
func (receiver *Controller) GinHandleUpdate(ctx *gin.Context) {
	req := &grid.UpdateRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := receiver.Update(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
