package user_grid

import (
	gin "github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/generate/proto/common/grid"
	http "github.com/go-home-admin/home/app/http"
)

// Export   列表导出
func (receiver *Controller) Export(req *grid.ExportRequest, ctx http.Context) (*grid.ExportResponse, error) {
	// TODO 这里写业务
	return &grid.ExportResponse{}, nil
}

// GinHandleExport gin原始路由处理
// http.Get(/user/export)
func (receiver *Controller) GinHandleExport(ctx *gin.Context) {
	req := &grid.ExportRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := receiver.Export(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
