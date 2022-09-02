package admin_user

import (
	gin "github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/app/servers/guid_scui"
	"github.com/go-home-admin/go-admin/generate/proto/common/guid"
	http "github.com/go-home-admin/home/app/http"
)

// User   用户列表
func (receiver *Controller) User(req *guid.TableRequest, ctx http.Context) (*guid.TableResponse, error) {
	table := guid_scui.NewTable()
	table.Column("姓名", "name").Width("150")
	table.Column("性别", "sex").Width("150").Filters([]*guid.Filter{{Text: "男", Value: "男"}, {Text: "女", Value: "女"}})
	table.Column("邮箱", "email").Width("250")
	table.Column("进度", "progress").Width("250").Sortable(true)
	table.Column("注册时间", "datetime").Width("150").Sortable(true)

	return &guid.TableResponse{
		Api: "",
		Table: &guid.TableData{
			Page:     0,
			PageSize: 0,
			Rows:     "",
			Total:    0,
			Summary:  0,
		},
		Columns: table.GetColumn(),
	}, nil
}

// GinHandleUser gin原始路由处理
// http.Get(/admin/user)
func (receiver *Controller) GinHandleUser(ctx *gin.Context) {
	req := &guid.TableRequest{}
	err := ctx.ShouldBind(req)
	context := http.NewContext(ctx)
	if err != nil {
		context.Fail(err)
		return
	}

	resp, err := receiver.User(req, context)
	if err != nil {
		context.Fail(err)
		return
	}

	context.Success(resp)
}
