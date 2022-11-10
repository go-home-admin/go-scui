package system

import (
	gin "github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/app/entity/mysql"
	admin "github.com/go-home-admin/go-admin/generate/proto/admin"
	http "github.com/go-home-admin/home/app/http"
	"github.com/go-home-admin/home/bootstrap/services/database"
)

// MenuSave 菜单保存
func (receiver *Controller) MenuSave(req *admin.MenuSaveRequest, ctx http.Context) (*admin.MenuSaveResponse, error) {
	id := req.Id
	apiList := string(database.NewJSON(req.ApiList))
	if req.Id == 0 {
		id = mysql.NewOrmAdminMenu().InsertGetId(&mysql.AdminMenu{
			ParentId:  req.ParentId,
			Order:     10000,
			Name:      req.Name,
			Component: req.Component,
			Path:      &req.Path,
			Meta:      string(database.NewJSON(req.Meta)),
			ApiList:   &apiList,
			CreatedAt: database.NowPointer(),
			UpdatedAt: database.NowPointer(),
		})
	} else {
		mysql.NewOrmAdminMenu().WhereId(id).Updates(&mysql.AdminMenu{
			ParentId:  req.ParentId,
			Name:      req.Name,
			Component: req.Component,
			Path:      &req.Path,
			Meta:      string(database.NewJSON(req.Meta)),
			ApiList:   &apiList,
			UpdatedAt: database.NowPointer(),
		})
	}
	return &admin.MenuSaveResponse{
		Id: id,
	}, nil
}

// GinHandleMenuSave gin原始路由处理
// http.Post(/system/menu/save)
func (receiver *Controller) GinHandleMenuSave(ctx *gin.Context) {
	req := &admin.MenuSaveRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := receiver.MenuSave(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
