package gui

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/generate/proto/common/grid"
	"github.com/go-home-admin/home/protobuf"
)

type Index interface {
	Index(ctx *gin.Context)
}

type ToResponse interface {
	ToResponse() *grid.IndexResponse
}

type Paginate interface {
	Paginate() ([]*protobuf.Any, int64)
}
