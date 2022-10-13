package user_grid

import (
	gin "github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/generate/proto/common/grid"
	http "github.com/go-home-admin/home/app/http"
)

// List   列表数据
func (receiver *Controller) List(req *grid.ListRequest, ctx http.Context) (*grid.ListResponse, error) {
	data, total := NewGuiContext(ctx).Paginate()

	return &grid.ListResponse{
		Page:     req.Page,
		PageSize: req.PageSize,
		Rows:     data,
		Total:    uint64(total),
		Summary:  0,
	}, nil
}

// GinHandleList gin原始路由处理
// http.Get(/user/list)
func (receiver *Controller) GinHandleList(ctx *gin.Context) {
	req := &grid.ListRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := receiver.List(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
