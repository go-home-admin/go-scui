package grid_scui

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"strings"
)

// Render 组件功能, 只能由视图装载
type Render struct {
	id string
	// template 里面的内容如果存在#__ID__, 那么会替换成成组件id
	// 如果需要插槽, 可以使用<name/>
	template string
	// 子组件
	append []*Slot
	// 组件拥有的值, 会直接追加到父级data, 所以key最好加前奏, 如果允许多个，那么应该拼接__ID__
	data map[string]interface{}
	// [函数名称]代码
	methodsRender map[string]string
	mountedRender map[string]string
}

func (r *Render) SetID(id string) {
	r.id = id
}

// GetID 例如同时渲染多个表格，是需要一个key区分
func (r *Render) GetID() string {
	if r.id == "" {
		r.id = uuid.NewV4().String()
	}

	return r.id
}

func (r *Render) repID(t string) string {
	return strings.ReplaceAll(t, "__ID__", r.id)
}

// Slot 模拟插槽
type Slot struct {
	Name string // 插入父级的标签， <name>
	*Render
}

func (r *Render) forAppend() []*Slot {
	if r.append == nil {
		r.append = make([]*Slot, 0)
	}
	return r.append
}

func (r *Render) AddRender(render *Render, slots ...string) *Render {
	if r.append == nil {
		r.append = make([]*Slot, 0)
	}
	if len(slots) != 0 {
		for _, slot := range slots {
			r.append = append(r.append, &Slot{
				Name:   "<slot #id=\"" + slot + "\"/>",
				Render: render,
			})
		}
	} else {
		r.append = append(r.append, &Slot{
			Name:   "",
			Render: render,
		})
	}

	return render
}

func (r *Render) GetTemplate() string {
	template := r.template
	for _, slot := range r.forAppend() {
		if slot.Name == "" {
			// 没有插槽, 追加到模版未
			template = template + slot.GetTemplate()
		} else {
			// 放入插槽, 但是不删除插槽, 防止有多个
			template = strings.ReplaceAll(template, slot.Name, slot.GetTemplate()+slot.Name)
		}
	}

	return r.repID(template)
}

func (r *Render) GetData() map[string]interface{} {
	data := r.data
	for _, slot := range r.forAppend() {
		for k, v := range slot.data {
			data[k] = v
		}
	}
	v, _ := json.Marshal(data)
	s := r.repID(string(v))
	nData := make(map[string]interface{})
	_ = json.Unmarshal([]byte(s), &nData)
	return nData
}

func (r *Render) forMethods() map[string]string {
	if r.methodsRender == nil {
		r.methodsRender = make(map[string]string)
	}
	return r.methodsRender
}
func (r *Render) GetMethods() string {
	methodsRender := r.forMethods()
	for _, slot := range r.forAppend() {
		for k, v := range slot.forMethods() {
			methodsRender[k] = v
		}
	}
	str := ""
	for funName, funStr := range methodsRender {
		str = str + "\n" + funName + ":" + funStr + ",\n"
	}
	return "{" + r.repID(str) + "}"
}
func (r *Render) forMounted() map[string]string {
	if r.mountedRender == nil {
		r.mountedRender = make(map[string]string)
	}
	return r.mountedRender
}
func (r *Render) GetMounted() string {
	mountedRender := r.forMounted()
	for _, slot := range r.forAppend() {
		for k, v := range slot.forMounted() {
			mountedRender[k] = v
		}
	}
	str := ""
	for _, funStr := range mountedRender {
		str = str + "\n" + funStr + ",\n"
	}
	return r.repID(str)
}

// View 视图
// 只能由前端defineAsyncComponent动态加载
type View struct {
	Render

	// 这里是同平常一样的 Vue 组件选项
	Template string `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty" form:"template"`
	// 组件数据
	Data interface{} `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty" form:"data"`
	// js 函数 {}
	Methods string `protobuf:"bytes,3,opt,name=methods,proto3" json:"methods,omitempty" form:"methods"`
	// js 执行代码
	Mounted string `protobuf:"bytes,4,opt,name=mounted,proto3" json:"mounted,omitempty" form:"mounted"`
}

// Rendering 返回符合defineAsyncComponent要求
func (v *View) Rendering() View {
	return View{
		Template: v.GetTemplate(),
		Data:     v.GetData(),
		Methods:  v.GetMethods(),
		Mounted:  v.GetMounted(),
	}
}
