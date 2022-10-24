package gui

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/generate/proto/common/grid"
	"github.com/go-home-admin/home/protobuf"
	"gorm.io/gorm"
)

type Gin interface {
	Gin() *gin.Context
}

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

type FormItems struct {
	// 对应的数据库字段, 参数类型(存入数据库可以格式化): string、int...
	Field     string `json:"-"`
	ValueType string `json:"-"`

	// 输出到前端的字段
	Label          string                 `json:"label,omitempty"`
	Name           string                 `json:"name,omitempty"`
	Value          string                 `json:"value,omitempty"`
	Component      string                 `json:"component,omitempty"`
	Span           int                    `json:"span"`
	Options        map[string]interface{} `json:"options,omitempty"`
	Rules          map[string]string      `json:"rules,omitempty"`
	RequiredHandle string                 `json:"required_handle,omitempty"`
	HideHandle     string                 `json:"hide_handle,omitempty"`
}

func (i *FormItems) GetOpt() string {
	s := ""
	for s2, i2 := range i.Options {
		switch i2.(type) {
		case string:
			s = s + " " + s2 + `="` + i2.(string) + `"`
		}
	}

	return s
}

func (i *FormItems) SaveToInt() *FormItems {
	i.ValueType = "int"
	return i
}
func (i *FormItems) SaveToString() *FormItems {
	i.ValueType = "string"
	return i
}

type GetFormItems interface {
	GetFormItems() []*FormItems
}
