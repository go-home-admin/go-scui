package base

import (
	"encoding/json"
	"github.com/go-home-admin/home/app"
	"github.com/sirupsen/logrus"
	"math/rand"
	"strings"
	"time"
)

const RenderID = "__ID__"

type RenderBase interface {
	SetID(id string)
	GetID() string
	AddRender(render RenderBase, slots ...string) RenderBase
	AppendTemplate(s string)
	GetTemplate(pr ...RenderBase) string
	AddData(k string, v interface{})
	GetData() map[string]interface{}
	GetMethods() string
	GetMethodsRender() map[string]string
	GetMounted() string
	GetMountedRender() map[string]string
}

// Render 组件功能, 只能由视图装载
type Render struct {
	id string
	// Template 里面的内容如果存在#__ID__, 那么会替换成成组件id
	// 如果需要插槽, 可以使用<name/>
	Template string `json:"-"`
	// 追加的模版内容, 直接拼接
	templateAppend string
	templateData   map[string]string
	// 模版自定义替换的内容
	templateReplace map[string]interface{}
	// 子组件
	append []*Slot
	// 组件拥有的值, 会直接追加到父级data, 所以key最好加前奏, 如果允许多个，那么应该拼接__ID__
	data map[string]interface{}
	// [函数名称]代码
	MethodsRender map[string]string `json:"-"`
	MountedRender map[string]string `json:"-"`
	// 防止反复渲染
	toTemplate bool
}

func (r *Render) AppendTemplate(s string) {
	r.templateAppend = r.templateAppend + s
}

func (r *Render) GetMethodsRender() map[string]string {
	methodsRender := r.MethodsRender
	for _, slot := range r.append {
		for k, v := range slot.GetMethodsRender() {
			methodsRender[k] = v
		}
	}

	return methodsRender
}

func (r *Render) GetMountedRender() map[string]string {
	mountedRender := r.MountedRender
	for _, slot := range r.append {
		for k, v := range slot.GetMountedRender() {
			mountedRender[k] = v
		}
	}
	return r.MountedRender
}

func NewRender(templates ...string) *Render {
	template := ""
	if len(templates) != 0 {
		template = templates[0]
	}
	return &Render{
		id:              Krand(10, KC_RAND_KIND_UPPER),
		Template:        template,
		templateData:    make(map[string]string),
		append:          make([]*Slot, 0),
		data:            make(map[string]interface{}),
		MethodsRender:   make(map[string]string),
		MountedRender:   make(map[string]string),
		templateReplace: make(map[string]interface{}),
	}
}

func (r *Render) SetID(id string) {
	r.id = id
}

const (
	KC_RAND_KIND_NUM   = 0 // 纯数字
	KC_RAND_KIND_LOWER = 1 // 小写字母
	KC_RAND_KIND_UPPER = 2 // 大写字母
	KC_RAND_KIND_ALL   = 3 // 数字、大小写字母
)

// 随机字符串
func Krand(size int, kind int) string {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)
	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}

// GetID 例如同时渲染多个表格，是需要一个key区分
func (r *Render) GetID() string {
	return r.id
}

// AddRep 替换模版内容, 把key替换成
// 1, (key, v1, "s1", "ss2")		把key替换成v1="s1:ss2; s3:ss3"
// 2, (key, v1, "s1") 			把key替换成v1="s1" v2="s2"
// 3, (key, v1) 				把key替换成v1
func (r *Render) AddRep(k string, v ...string) string {
	switch len(v) {
	case 1: // 把key替换成v1
		r.templateReplace[k] = v[0]
	case 2: // 把key替换成v1="s1" v2="s2"
		if _, ok := r.templateReplace[k]; !ok {
			r.templateReplace[k] = map[string]interface{}{v[0]: v[1]}
		} else {
			switch r.templateReplace[k].(type) {
			case string:
				r.templateReplace[k] = map[string]interface{}{v[0]: v[1]}
			}
			m := r.templateReplace[k].(map[string]interface{})
			m[v[0]] = v[1]
			r.templateReplace[k] = m
		}
	case 3: // (key, v1, "s1", ";")		把key替换成v1="s1; s2"
		if _, ok := r.templateReplace[k]; !ok {
			r.templateReplace[k] = map[string]interface{}{v[0]: map[string]string{v[1]: v[2]}}
		} else {
			switch r.templateReplace[k].(type) {
			case string:
				r.templateReplace[k] = map[string]interface{}{v[0]: map[string]string{v[1]: v[2]}}
			}
			m := r.templateReplace[k].(map[string]interface{})
			mm := m[v[0]].(map[string]string)
			mm[v[1]] = v[2]
			m[v[0]] = mm
			r.templateReplace[k] = m
		}
	}

	return r.id
}

func (r *Render) RepAll(t string) string {
	r.AddRep(RenderID, r.GetID())
	for s, i := range r.templateReplace {
		switch i.(type) {
		case string:
			t = strings.ReplaceAll(t, s, i.(string))
		case map[string]interface{}:
			str := ""
			for k3, s3 := range i.(map[string]interface{}) {
				switch s3.(type) {
				case string:
					str += " " + k3 + `="` + s3.(string) + `"`
				case map[string]string:
					str2 := ""
					for k4, s4 := range s3.(map[string]string) {
						str2 += k4 + `:` + s4 + `; `
					}
					str += " " + k3 + `="` + str2 + `"`
				}
			}
			t = strings.ReplaceAll(t, s, str)
		}
	}

	return t
}

func (r *Render) RepID(t string) string {
	return strings.ReplaceAll(t, RenderID, r.GetID())
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

func (r *Render) GetTemplate(pr ...RenderBase) string {
	template := r.Template
	// 处理需要替换的值
	for s, i := range r.templateData {
		template = strings.ReplaceAll(template, s, i)
	}
	for i, slot := range r.append {
		if slot.Name == "" {
			// 没有插槽, 追加到模版未
			if len(pr) == 0 {
				template = template + slot.GetTemplate(r)
			} else if !r.toTemplate {
				r.toTemplate = true
				// 添加到上级
				for _, base := range pr {
					base.AppendTemplate(r.RepID(slot.GetTemplate(r)))
				}
			}
		} else if (i + 1) == len(r.append) {
			// 放入插槽, 最后一个
			template = strings.ReplaceAll(template, slot.Name, slot.GetTemplate(r))
		} else {
			// 放入插槽, 但是不删除插槽, 防止有多个
			template = strings.ReplaceAll(template, slot.Name, slot.GetTemplate(r)+slot.Name)
		}
	}

	return r.RepAll(template + r.templateAppend)
}

// AddData 添加组件的数据, ("k.k.1", "2")
func (r *Render) AddData(k string, v interface{}) {
	r.data = setData(k, r.data, v)
}

func setData(k string, o map[string]interface{}, v interface{}) map[string]interface{} {
	arr := strings.Split(k, ".")
	if len(arr) == 1 {
		// 最后一个赋值
		o[k] = v
	} else {
		kk := arr[0]
		nk := k[len(kk)+1:]
		if _, ok := o[kk]; !ok {
			o[kk] = make(map[string]interface{})
		} else if _, ok := o[kk].(map[string]interface{}); !ok {
			// 覆盖旧的值
			o[kk] = make(map[string]interface{})
		}
		o[kk] = setData(nk, o[kk].(map[string]interface{}), v)
	}

	return o
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
	s := r.RepID(string(v))
	nData := make(map[string]interface{})
	_ = json.Unmarshal([]byte(s), &nData)
	return nData
}

// AddMethods 函数名称, 函数js代码, 如果跨组件通常存在其他的ID需要替换
func (r *Render) AddMethods(funName string, code string, reps ...string) {
	if len(reps) != 0 {
		code = ReplaceAll(code, reps)
	}

	r.MethodsRender[r.RepID(funName)] = r.RepID(code)
}

// AddMounted 新增打开页面执行的代码, k作为唯一键, 防止重复新增
func (r *Render) AddMounted(k, code string) {
	r.MountedRender[k] = r.RepID(code)
}

func (r *Render) GetMethods() string {
	methodsRender := r.MethodsRender
	for _, slot := range r.append {
		for k, v := range slot.GetMethodsRender() {
			methodsRender[k] = v
		}
	}
	str := ""
	for funName, funStr := range methodsRender {
		str = str + "\n" + funName + ":" + funStr + ",\n"
	}
	return "{" + r.RepID(str) + "}"
}

func (r *Render) GetMounted() string {
	mountedRender := r.MountedRender
	for _, slot := range r.append {
		for k, v := range slot.GetMountedRender() {
			mountedRender[k] = v
		}
	}
	str := ""
	for _, funStr := range mountedRender {
		str = str + "\n" + funStr + "\n"
	}
	return r.RepID(str)
}

// View 视图
// 只能由前端defineAsyncComponent动态加载
type View struct {
	*Render

	file string

	// 这里是同平常一样的 Vue 组件选项
	Template string `protobuf:"bytes,1,opt,name=Template,proto3" json:"Template,omitempty" form:"Template"`
	// 组件数据
	Data interface{} `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty" form:"data"`
	// js 函数 {}
	Methods string `protobuf:"bytes,3,opt,name=methods,proto3" json:"methods,omitempty" form:"methods"`
	// js 执行代码
	Mounted string `protobuf:"bytes,4,opt,name=mounted,proto3" json:"mounted,omitempty" form:"mounted"`
}

func NewView(view string) *View {
	v := &View{
		file:   view,
		Render: LoadVue(view),
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

func (v *View) GetTemplate(pr ...RenderBase) string {
	return v.Render.GetTemplate(pr...)
}

// LoadJsFunction name, code
func LoadJsFunction(file string) (string, string) {
	s := LoadView(file)
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

func ReplaceAll(s string, r []string) string {
	for i := 0; i < len(r); i++ {
		s = strings.ReplaceAll(s, r[i], r[i+1])
		i++
	}
	return s
}
