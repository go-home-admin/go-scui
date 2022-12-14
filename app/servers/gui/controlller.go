package gui

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TryNumberToInt64 尝试对所有数字类型转int64
func TryNumberToInt64(i interface{}) int64 {
	switch i.(type) {
	case int:
		return int64(i.(int))
	case int32:
		return int64(i.(int32))
	case uint32:
		return int64(i.(uint32))
	case int64:
		return i.(int64)
	case uint64:
		return int64(i.(uint64))
	case float32:
		return int64(i.(float32))
	case float64:
		return int64(i.(float64))
	}
	return 0
}

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
type Delete interface {
	Delete(ctx *gin.Context)
}

type ToResponse interface {
	ToResponse() *IndexResponse
}

type Paginate interface {
	Paginate() (interface{}, int64)
}

type GetDB interface {
	TableName() string
	GetDB() *gorm.DB
}

type GetPrimary interface {
	GetPrimary() string
}

type ViewDB struct {
	model Model
}

type Model interface {
	TableName() string
	GetDB() *gorm.DB
}

func (v *ViewDB) SetDb(model Model) {
	v.model = model
}

func (v *ViewDB) GetDB() *gorm.DB {
	return v.model.GetDB()
}

func (v *ViewDB) TableName() string {
	return v.model.TableName()
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

func NewItems(prop, label, component string) *FormItems {
	return &FormItems{
		Label:     label,
		Name:      prop,
		Span:      24,
		Component: component,
		Options: map[string]interface{}{
			"placeholder": label,
		},
	}
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
