package system

import (
	gin "github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/app/entity/mysql"
	admin "github.com/go-home-admin/go-admin/generate/proto/admin"
	http "github.com/go-home-admin/home/app/http"
)

// MenuDel 菜单删除
func (receiver *Controller) MenuDel(req *admin.MenuDelRequest, ctx http.Context) (*admin.MenuDelResponse, error) {
	delChildren(req.Ids)

	return &admin.MenuDelResponse{}, nil
}

func delChildren(ids []uint32) {
	pis := mysql.NewOrmAdminMenu().WhereParentIdIn(ids).Get().GetIdList()
	if len(pis) != 0 {
		delChildren(pis)
	}
	mysql.NewOrmAdminMenu().WhereIdIn(ids).Delete()
}

// GinHandleMenuDel gin原始路由处理
// http.Delete(/system/menu/del)
func (receiver *Controller) GinHandleMenuDel(ctx *gin.Context) {
	req := &admin.MenuDelRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := receiver.MenuDel(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
