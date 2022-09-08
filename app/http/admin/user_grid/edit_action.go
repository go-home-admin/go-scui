package user_grid

import (
	gin "github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/generate/proto/common/grid"
	http "github.com/go-home-admin/home/app/http"
)

// Edit   编辑数据的页面
func (receiver *Controller) Edit(req *grid.EditRequest, ctx http.Context) (*grid.EditResponse, error) {
	// TODO 这里写业务
	return &grid.EditResponse{}, nil
}

// GinHandleEdit gin原始路由处理
// http.Get(/user/edit)
func (receiver *Controller) GinHandleEdit(ctx *gin.Context) {
	req := &grid.EditRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := receiver.Edit(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
