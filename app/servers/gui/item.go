package gui

type SelectOptions []Options

type Options struct {
	Label string      `json:"label"`
	Value interface{} `json:"value"`
}

type FilterOptions []Filter

type Filter struct {
	Text  string      `json:"text,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

type Column struct {
	// 名称
	Label string `json:"label,omitempty" form:"label"`
	// 字段
	Prop string `json:"prop,omitempty" form:"prop"`
	// 宽度
	Width string `json:"width,omitempty" form:"width"`
	// 是否隐藏
	Hide bool `json:"hide,omitempty" form:"hide"`
	// 是否开启排序
	Sortable bool `json:"sortable,omitempty" form:"sortable"`
	// 状态翻译
	Filters []Filter `json:"filters,omitempty" form:"filters"`
	// Vue 模板
	Template string `json:"template,omitempty" form:"-"`
}

// 这里应该输出一个完整的组件格式
type IndexResponse struct {
	// 这里是同平常一样的 Vue 组件选项
	Template string `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty" form:"template"`
	// 组件数据
	Data interface{} `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty" form:"data"`
	// js 函数 {}
	Methods string `protobuf:"bytes,3,opt,name=methods,proto3" json:"methods,omitempty" form:"methods"`
	// js 执行代码
	Mounted string `protobuf:"bytes,4,opt,name=mounted,proto3" json:"mounted,omitempty" form:"mounted"`
}

type ListResponse struct {
	// 页码
	Page int64 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty" form:"page"`
	// 分页大小
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty" form:"pageSize"`
	//分析行数据字段结构
	Rows interface{} `protobuf:"bytes,3,rep,name=rows,proto3" json:"rows,omitempty" form:"rows"`
	//分析总数字段结构
	Total uint64 `protobuf:"varint,4,opt,name=total,proto3" json:"total,omitempty" form:"total"`
	//分析合计行字段结构
	Summary uint64 `protobuf:"varint,5,opt,name=summary,proto3" json:"summary,omitempty" form:"summary"`
}
