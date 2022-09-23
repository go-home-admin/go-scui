package grid_scui

import "gorm.io/gorm"

type Search struct {
	Form
	where []*SearchWhere
}

func (s *Search) ToString() map[string]interface{} {
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
	FormItems
	// 对应的数据库字段
	Field string

	ormWhere func(db *gorm.DB)
}

func (s *SearchWhere) ToFormItems() FormItems {
	return s.FormItems
}

func (s *SearchWhere) Where() {

}

type Form struct {
	LabelWidth    string      `json:"labelWidth"`
	LabelPosition string      `json:"labelPosition"`
	Size          string      `json:"size"`
	FormItems     []FormItems `json:"formItems"`
}

type FormItems struct {
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

// Input 普通组件
func (s *Search) Input(prop, label string) *FilterInput {
	w := &SearchWhere{
		Field: prop,
		FormItems: FormItems{
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
