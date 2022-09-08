package user_grid

import (
	gin "github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/generate/proto/common/grid"
	http "github.com/go-home-admin/home/app/http"
	"strings"
)

// Index   列表页面
func (receiver *Controller) Index(req *grid.IndexRequest, ctx http.Context) (*grid.IndexResponse, error) {
	table := receiver.NewGridTable()
	table.SetUri(strings.Replace(ctx.Gin().Request.RequestURI, "?", "/list?", 1))
	return table.ToResponse(), nil
}

// GinHandleIndex gin原始路由处理
// http.Get(/user)
func (receiver *Controller) GinHandleIndex(ctx *gin.Context) {
	req := &grid.IndexRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := receiver.Index(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
