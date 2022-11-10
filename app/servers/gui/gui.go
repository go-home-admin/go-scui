package gui

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/home/app/http"
	"github.com/sirupsen/logrus"
	"strconv"
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
		if i, ok := g.Controller.(Create); ok {
			i.Create(g.Context)
		} else {
			g.Create(g.Context)
		}
	case "edit": // edit
		if i, ok := g.Controller.(Update); ok {
			i.Update(g.Context)
		} else {
			g.Update(g.Context)
		}
	case "del": // edit
		if i, ok := g.Controller.(Delete); ok {
			i.Delete(g.Context)
		} else {
			g.Delete(g.Context)
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
		context.Success(&ListResponse{
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
	by, _ := ctx.GetRawData()
	var data map[string]interface{}
	err := json.Unmarshal(by, &data)
	if err != nil {
		logrus.Error(err)
		return
	}

	td := g.Controller.(GetDB).GetDB().Create(&data)
	if td.Error != nil {
		logrus.Error(td.Error)
		http.NewContext(ctx).Fail(errors.New("创建失败"))
		return
	}
	http.NewContext(ctx).Success(nil)
}

func (g *GinHandle) Update(ctx *gin.Context) {
	data, all := g.getData(ctx)
	primary := g.Controller.(GetPrimary).GetPrimary()
	primaryValStringOrFloat64, ok := all[primary]
	if !ok {
		logrus.Error("必须要有主键数据才能更新, 当前的主建=" + primary)
		return
	}
	var primaryVal interface{}
	switch primaryValStringOrFloat64.(type) {
	case string:
		primaryVal = primaryValStringOrFloat64
	case float64:
		primaryVal = int(primaryValStringOrFloat64.(float64))
	}

	td := g.Controller.(GetDB).GetDB().Where(primary+" = ?", primaryVal).Updates(&data)
	if td.Error != nil {
		logrus.Error(td.Error)
		http.NewContext(ctx).Fail(errors.New("更新失败"))
		return
	}

	http.NewContext(ctx).Success(nil)
}

func (g *GinHandle) Delete(ctx *gin.Context) {
	primary := g.Controller.(GetPrimary).GetPrimary()
	var primaryVal interface{}
	primaryVal, ok := ctx.GetQuery(primary)
	if !ok {
		logrus.Error("必须要有主键数据才能删除, 当前的主建=" + primary)
		return
	}
	_ = primaryVal
	model := g.Controller.(GetDB)
	td := model.GetDB().Delete(model, primaryVal)
	if td.Error != nil {
		logrus.Error(td.Error)
		http.NewContext(ctx).Fail(errors.New("删除失败"))
		return
	}

	http.NewContext(ctx).Success(nil)
}
func (g *GinHandle) getData(ctx *gin.Context) (map[string]interface{}, map[string]interface{}) {
	by, _ := ctx.GetRawData()
	var m map[string]interface{}
	err := json.Unmarshal(by, &m)
	if err != nil {
		logrus.Error(err)
		return nil, nil
	}
	data := map[string]interface{}{}
	if c, ok := g.Controller.(GetFormItems); ok {
		for _, items := range c.GetFormItems() {
			if d, ok := m[items.Name]; ok {
				switch items.ValueType {
				case "int":
					d = toInt(d)
				}

				data[items.Name] = d
			}
		}
	}

	return data, m
}

func toInt(i interface{}) int {
	var d int
	switch i.(type) {
	case string:
		d, _ = strconv.Atoi(i.(string))
	default:
		d = int(i.(float64))
	}
	return d
}
