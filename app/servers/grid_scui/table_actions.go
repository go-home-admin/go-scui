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

func (r *RowAction) CreateAction() *Button {
	btn := NewButton("创建")
	btn.template = `<el-button type="primary" plain size="small" @click="table_edit(row)">编辑</el-button>`
	btn.AddMethods("table_edit", `
function(row){
	this.dialog.save = true
	this.$nextTick(() => {
		this.$refs.saveDialog.open('edit').setData(row)
	})
}
`)
	btn.AddData("dialog", map[string]bool{"save": false})
	r.t.AddRender(btn, "actions")
	r.t.AddRender(NewRender(loadView("table_create.vue")))
	return btn
}
