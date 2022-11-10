package system

import (
	"encoding/json"
	gin "github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/app/entity/mysql"
	admin "github.com/go-home-admin/go-admin/generate/proto/admin"
	http "github.com/go-home-admin/home/app/http"
)

// MenuList 菜单列表
func (receiver *Controller) MenuList(req *admin.MenuListRequest, ctx http.Context) (*admin.MenuListResponse, error) {
	got := make([]*admin.MenuInfo, 0)
	menus := make(map[uint32]*admin.MenuInfo)
	list := mysql.NewOrmAdminMenu().Order("sort asc").Order("id asc").Get()

	for _, menu := range list {
		meta := &admin.Meta{}
		api := make([]*admin.ApiList, 0)
		if menu.Meta != "" {
			json.Unmarshal([]byte(menu.Meta), meta)
		}
		if menu.ApiList != nil {
			json.Unmarshal([]byte(*menu.ApiList), &api)
		}
		menus[menu.Id] = &admin.MenuInfo{
			Id:        menu.Id,
			ParentId:  menu.ParentId,
			Name:      menu.Name,
			Path:      *menu.Path,
			Meta:      meta,
			Component: menu.Component,
			Children:  make([]*admin.MenuInfo, 0),
			ApiList:   api,
		}

		if menu.ParentId != 0 {
			m, mok := menus[menu.Id]
			if mok {
				p, pok := menus[menu.ParentId]
				if pok {
					p.Children = append(p.Children, m)
				}
			}
		}
	}
	for _, menu := range list {
		if menu.ParentId == 0 {
			got = append(got, menus[menu.Id])
		}
	}

	return &admin.MenuListResponse{
		List: got,
	}, nil
}

// GinHandleMenuList gin原始路由处理
// http.Get(/system/menu/list)
func (receiver *Controller) GinHandleMenuList(ctx *gin.Context) {
	req := &admin.MenuListRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := receiver.MenuList(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
