package grid_scui

import (
	"github.com/go-home-admin/home/app/http"
	"gorm.io/gorm"
)

type Search struct {
	http.Context
	*Form
	where []*SearchWhere
}

func NewTableSearch(ctx http.Context) *Search {
	search := &Search{
		Context: ctx,
		Form:    NewTableForm(),
	}
	search.AddMethods("onSubmit", `async function(){
		this.$refs.table.upData()
	}`)

	return search
}

func (s *Search) GetData() map[string]interface{} {
	config := s.Form
	for _, where := range s.where {
		config.FormItems = append(config.FormItems, where.ToFormItems())
	}

	data := map[string]interface{}{
		"config":  config,
		"form":    map[string]interface{}{},
		"loading": false,
	}

	return data
}

type SearchWhere struct {
	*FormItems
	// 对应的数据库字段
	Field string

	ormWhere func(db *gorm.DB)
}

func (s *SearchWhere) ToFormItems() *FormItems {
	return s.FormItems
}

// Input 普通组件
func (s *Search) Input(prop, label string) *FilterInput {
	w := &SearchWhere{
		Field: prop,
		FormItems: &FormItems{
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
