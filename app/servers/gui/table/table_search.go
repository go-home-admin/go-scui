package table

import (
	"github.com/gin-gonic/gin"
	"github.com/go-home-admin/go-admin/app/servers/gui"
	"github.com/go-home-admin/go-admin/app/servers/gui/form"
	"gorm.io/gorm"
)

type Search struct {
	*gin.Context
	*form.Form
	where []*SearchWhere
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

func (s *Search) GetData() map[string]interface{} {
	config := s.Form
	if len(config.FormItems) == 0 {
		for _, where := range s.where {
			config.FormItems = append(config.FormItems, where.ToFormItems())
		}
	}

	data := map[string]interface{}{
		"config":  config,
		"form":    map[string]interface{}{},
		"loading": false,
	}

	return data
}

type SearchWhere struct {
	*gui.FormItems
	// 对应的数据库字段
	Field string

	ormWhere func(db *gorm.DB)
}

func (s *SearchWhere) ToFormItems() *gui.FormItems {
	return s.FormItems
}

// Input 普通组件
func (s *Search) Input(prop, label string) *FilterInput {
	w := &SearchWhere{
		Field: prop,
		FormItems: &gui.FormItems{
			Label:     label,
			Name:      prop,
			Component: "input",
			Options: map[string]interface{}{
				"placeholder": prop,
			},
		},
	}
	s.where = append(s.where, w)
	return &FilterInput{SearchWhere: w}
}
