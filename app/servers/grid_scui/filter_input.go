package grid_scui

type FilterInput struct {
	*SearchWhere
}

func (f *FilterInput) Placeholder(v string) *FilterInput {
	f.Options["placeholder"] = v

	return f
}

func (f *FilterInput) Span(v int) *FilterInput {
	f.SearchWhere.Span = v
	return f
}
