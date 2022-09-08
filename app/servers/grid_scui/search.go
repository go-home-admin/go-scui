package grid_scui

type Search struct {
	form  Form
	where []SearchWhere
}

type SearchWhere struct {
	// 对应的数据库字段
	Field string
	// 表单字段
	Prop string
	FormItems
}

func (s *SearchWhere) Where() {

}

type Form struct {
	LabelWidth    string
	LabelPosition string
	Size          string
	FormItems     []FormItems
}

type FormItems struct {
	Label          string
	Name           string
	Value          string
	Component      string
	Options        map[string]interface{}
	Rules          map[string]string
	RequiredHandle string
	hideHandle     string
}

// 普通组件
func (s *Search) Input(prop string) *SearchWhere {
	return nil
}

// 组件 select
func (s *Search) Select(prop string) *SearchWhere {
	return nil
}

func (s *Search) Between(prop string) *SearchWhere {
	return nil
}
