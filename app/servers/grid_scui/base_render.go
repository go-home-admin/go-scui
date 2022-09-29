package grid_scui

import (
	"embed"
	"encoding/json"
	"github.com/go-home-admin/home/app"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"
)

const RenderID = "__ID__"

type RenderBase interface {
	SetID(id string)
	GetID() string
	AddRender(render RenderBase, slots ...string) RenderBase
	GetTemplate() string
	AddData(k string, v interface{})
	GetData() map[string]interface{}
	GetMethods() string
	MethodsRender() map[string]string
	GetMounted() string
	MountedRender() map[string]string
}

// Render 组件功能, 只能由视图装载
type Render struct {
	id string
	// template 里面的内容如果存在#__ID__, 那么会替换成成组件id
	// 如果需要插槽, 可以使用<name/>
	template     string
	templateData map[string]string
	// 子组件
	append []*Slot
	// 组件拥有的值, 会直接追加到父级data, 所以key最好加前奏, 如果允许多个，那么应该拼接__ID__
	data map[string]interface{}
	// [函数名称]代码
	methodsRender map[string]string
	mountedRender map[string]string
}

func (r *Render) MethodsRender() map[string]string {
	return r.methodsRender
}

func (r *Render) MountedRender() map[string]string {
	return r.mountedRender
}

func NewRender(templates ...string) *Render {
	template := ""
	if len(templates) != 0 {
		template = templates[0]
	}
	return &Render{
		template:      template,
		templateData:  make(map[string]string),
		append:        make([]*Slot, 0),
		data:          make(map[string]interface{}),
		methodsRender: make(map[string]string),
		mountedRender: make(map[string]string),
	}
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
	return strings.ReplaceAll(t, RenderID, r.id)
}

// Slot 模拟插槽
type Slot struct {
	PreData string // data放到父级key下
	Name    string // 插入父级的标签， <name>
	RenderBase
}

func (r *Render) AddTemplateData(k, v string) {
	r.templateData[k] = v
}

// AddRender （要添加的组件，插槽，data前缀)
func (r *Render) AddRender(render RenderBase, slots ...string) RenderBase {
	slot := &Slot{
		RenderBase: render,
	}

	switch len(slots) {
	case 2:
		slot.PreData = slots[1]
		slot.Name = "<slot id=\"" + slots[0] + "\"/>"
	case 1:
		slot.Name = "<slot id=\"" + slots[0] + "\"/>"
	}
	r.append = append(r.append, slot)
	return slot
}

func (r *Render) GetTemplate() string {
	template := r.template
	// 处理需要替换的值
	for s, i := range r.templateData {
		template = strings.ReplaceAll(template, s, i)
	}
	for i, slot := range r.append {
		if slot.Name == "" {
			// 没有插槽, 追加到模版未
			template = template + slot.GetTemplate()
		} else if (i + 1) == len(r.append) {
			// 放入插槽, 最后一个
			template = strings.ReplaceAll(template, slot.Name, slot.GetTemplate())
		} else {
			// 放入插槽, 但是不删除插槽, 防止有多个
			template = strings.ReplaceAll(template, slot.Name, slot.GetTemplate()+slot.Name)
		}
	}

	return r.repID(template)
}

func (r *Render) AddData(k string, v interface{}) {
	r.data[k] = v
}

func (r *Render) GetData() map[string]interface{} {
	data := r.data
	for _, slot := range r.append {
		for k, v := range slot.GetData() {
			if slot.PreData != "" {
				v2, ok := data[slot.PreData]
				if !ok {
					v2 = make(map[string]interface{})
				}
				PreDataMap, ok2 := v2.(map[string]interface{})
				if !ok2 {
					logrus.Error("前缀key类型错误")
					continue
				}
				PreDataMap[k] = v
				data[slot.PreData] = PreDataMap
			} else {
				data[k] = v
			}
		}
	}
	v, err := json.Marshal(data)
	if err != nil {
		logrus.Error(err)
	}
	s := r.repID(string(v))
	nData := make(map[string]interface{})
	_ = json.Unmarshal([]byte(s), &nData)
	return nData
}

func (r *Render) AddMethods(funName string, code string) {
	r.methodsRender[funName] = code
}

func (r *Render) GetMethods() string {
	methodsRender := r.methodsRender
	for _, slot := range r.append {
		for k, v := range slot.MethodsRender() {
			methodsRender[k] = v
		}
	}
	str := ""
	for funName, funStr := range methodsRender {
		str = str + "\n" + funName + ":" + funStr + ",\n"
	}
	return "{" + r.repID(str) + "}"
}

func (r *Render) GetMounted() string {
	mountedRender := r.mountedRender
	for _, slot := range r.append {
		for k, v := range slot.MountedRender() {
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

	file string

	// 这里是同平常一样的 Vue 组件选项
	Template string `protobuf:"bytes,1,opt,name=template,proto3" json:"template,omitempty" form:"template"`
	// 组件数据
	Data interface{} `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty" form:"data"`
	// js 函数 {}
	Methods string `protobuf:"bytes,3,opt,name=methods,proto3" json:"methods,omitempty" form:"methods"`
	// js 执行代码
	Mounted string `protobuf:"bytes,4,opt,name=mounted,proto3" json:"mounted,omitempty" form:"mounted"`
}

func NewView(view string) *View {
	v := &View{
		file: view,
		Render: Render{
			data:          map[string]interface{}{},
			methodsRender: map[string]string{},
			mountedRender: make(map[string]string),
		},
	}
	return v
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

func (v *View) GetTemplate() string {
	view := loadView(v.file)
	v.Render.template = view

	return v.Render.GetTemplate()
}

//go:embed views
var views embed.FS

func loadView(file string) string {
	vTable := ""
	if app.IsDebug() {
		if _, err := os.Stat("app/servers/grid_scui/views/" + file); err == nil {
			v, _ := ioutil.ReadFile("app/servers/grid_scui/views/" + file)
			vTable = string(v)
		}
	} else {
		fileContext, _ := views.ReadFile("views/" + file)
		vTable = string(fileContext)
	}

	return vTable
}

// name, code
func loadJsFunction(file string) (string, string) {
	s := loadView(file)
	index := strings.Index(s, "(")

	name := strings.TrimSpace(s[8:index])
	return name, "function (" + s[index+1:]
}

func ToUrl(uri string) string {
	if strings.Index(uri, "http") != 0 {
		uri = app.Config("app.url", "http://127.0.0.1:8080") + uri
	}
	return uri
}
