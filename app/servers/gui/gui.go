package gui

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/generate/proto/common/grid"
	"github.com/go-home-admin/home/app/http"
)

func NewGui(ctx *gin.Context) *GinHandle {
	return &GinHandle{Context: ctx}
}

type GinHandle struct {
	Context    *gin.Context
	Controller interface{}
}

func (g *GinHandle) SetController(controller interface{}) {
	g.Controller = controller
}

func (g *GinHandle) Gin() *gin.Context {
	return g.Context
}

func (g *GinHandle) ActionHandle() {
	action := g.Gin().Param("action")

	switch action {
	case "", "index": // 列表页面
		if i, ok := g.Controller.(Index); ok {
			i.Index(g.Context)
		} else {
			g.Index(g.Context)
		}
	case "list": // 列表数据
		g.List(g.Context)
	case "create": // 创建
		if i, ok := g.Controller.(Index); ok {
			i.Index(g.Context)
		} else {
			g.Index(g.Context)
		}
	default:
		http.NewContext(g.Context).Fail(errors.New("不支持的路由"))
	}
}

func (g *GinHandle) Index(ctx *gin.Context) {
	context := http.NewContext(ctx)
	if i, ok := g.Controller.(ToResponse); ok {
		context.Success(i.ToResponse())
	} else {
		context.Fail(errors.New("结构未实现ToResponse"))
	}
}

func (g *GinHandle) List(ctx *gin.Context) {
	context := http.NewContext(ctx)
	if i, ok := g.Controller.(Paginate); ok {
		data, total := i.Paginate()
		context.Success(&grid.ListResponse{
			Page:     1,
			PageSize: 15,
			Rows:     data,
			Total:    uint64(total),
			Summary:  0,
		})
	} else {
		context.Fail(errors.New("结构未实现ToResponse"))
	}
}

func (g *GinHandle) Create(ctx *gin.Context) {

}
