package gui

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/generate/proto/common/grid"
	"github.com/go-home-admin/home/protobuf"
	"gorm.io/gorm"
)

type Index interface {
	Index(ctx *gin.Context)
}
type Create interface {
	Create(ctx *gin.Context)
}
type Update interface {
	Update(ctx *gin.Context)
}

type ToResponse interface {
	ToResponse() *grid.IndexResponse
}

type Paginate interface {
	Paginate() ([]*protobuf.Any, int64)
}

type GetDB interface {
	GetDB() *gorm.DB
}

type GetPrimary interface {
	GetPrimary() string
}

type ViewDB struct {
	db *gorm.DB
}

func (v *ViewDB) SetDb(db *gorm.DB) {
	v.db = db
}

func (v *ViewDB) GetDB() *gorm.DB {
	return v.db
}

type ViewPrimary struct {
	primary string
}

func (v *ViewPrimary) SetPrimary(p string) {
	v.primary = p
}

func (v *ViewPrimary) GetPrimary() string {
	if v.primary == "" {
		return "id"
	}
	return v.primary
}
