package html

import "github.com/go-home-admin/go-admin/app/servers/gui/base"

const el_table_column = `<el-table-column label="性别" prop="sex" width="150" :filters="sexFilters" :filter-method="filterHandler">
<template #default="scope">
<slot id="template"/>
</template>
</el-table-column>`

type elTableColumn struct {
	*base.Render
}

func NewElTableColumn() *elTableColumn {
	return &elTableColumn{
		Render: base.NewRender(el_table_column),
	}
}
