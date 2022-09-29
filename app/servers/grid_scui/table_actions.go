package grid_scui

import "github.com/go-home-admin/home/app/http"

type RowAction struct {
	http.Context

	t *Table
}

func (r *RowAction) AddButton(text string) *Button {
	btn := NewButton(text)
	r.t.AddRender(btn, "actions")
	return btn
}
