package table

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/app/servers/gui/form"
)

type Search struct {
	*gin.Context
	*form.Form
}

func NewTableSearch(ctx *gin.Context) *Search {
	search := &Search{
		Context: ctx,
		Form:    form.NewTableForm(),
	}
	search.AddMethods("onSubmit", `async function(){
		this.$refs.table.upData()
	}`)

	return search
}
