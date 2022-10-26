package table

import (
	"fmt"
	"github.com/go-home-admin/go-admin/app/servers/gui"
	"github.com/go-home-admin/go-admin/app/servers/gui/base"
	"github.com/sirupsen/logrus"
	time2 "time"
)

type Row map[string]interface{}

type Column struct {
	*gui.Column
	*base.Render
	funcs []func(interface{}, Row) interface{}
	// k=>v (k 数字类型只会是int64)
	mapVal map[interface{}]interface{}
}

func NewColumn() *Column {
	return &Column{
		Column: &gui.Column{
			Filters: make([]gui.Filter, 0),
		},
		Render: base.NewRender(),
		funcs:  make([]func(interface{}, Row) interface{}, 0),
		mapVal: make(map[interface{}]interface{}),
	}
}

func (c *Column) Width(v string) *Column {
	c.Column.Width = v
	return c
}

func (c *Column) Hide(v bool) *Column {
	c.Column.Hide = v
	return c
}

func (c *Column) Sortable(v bool) *Column {
	c.Column.Sortable = v
	return c
}

// Date https://go-zh.org/pkg/time/#Time.Format
func (c *Column) Date(format ...string) *Column {
	if len(format) == 0 {
		format = append(format, "2006-01-02 15:04:05")
	}

	c.Display(func(val interface{}, row Row) interface{} {
		switch val.(type) {
		case time2.Time:
			t := val.(time2.Time)
			return t.Format(format[0])
		default:
			return val
		}
	})
	return c
}

func (c *Column) Filters(v []gui.Filter) *Column {
	c.Column.Filters = v
	c.Using(v)
	return c
}

// Using 不同值不同显示
func (c *Column) Using(v []gui.Filter) *Column {
	newKey := "__" + c.Prop + "__"
	c.Display(func(val interface{}, row Row) interface{} {
		for _, filter := range v {
			row["__"+c.Prop+"__"] = val
			if fmt.Sprint(filter.Value) == fmt.Sprint(val) {
				row[newKey] = fmt.Sprint(filter.Text)
				break
			}
		}
		return val
	})

	c.Template("<template #" + c.Prop + "=\"scope\">{{ scope.row." + newKey + "}}</template>")
	return c
}

// Template 放到插槽 el-table-column
func (c *Column) Template(v string) *Column {
	c.Render.Template = v
	return c
}

// Display 自定义处理值 val 当前值 row 行
func (c *Column) Display(f func(val interface{}, row Row) interface{}) *Column {
	c.funcs = append(c.funcs, f)
	return c
}

// Translate 对列数据翻译
func (c *Column) Translate(k, v interface{}) *Column {
	switch k.(type) {
	case string:
		c.mapVal[k] = v
	case int, int32, uint32, int64, uint64, float32, float64:
		c.mapVal[gui.TryNumberToInt64(k)] = v
	default:
		c.mapVal[k] = v
	}
	if len(c.funcs) == 0 {
		c.addMapVal()
	}
	return c
}

func (c *Column) addMapVal() *Column {
	c.Display(func(val interface{}, row Row) interface{} {
		if kk, ok := c.mapVal[val]; ok {
			return kk
		}
		// 数字类型太多, 自动识别类似类型, 作为转义值应该不会超过int64
		if kk, ok := c.mapVal[gui.TryNumberToInt64(val)]; ok {
			return kk
		}

		return val
	})
	return c
}

// Map 替换值, 不想被替换应该使用Using
func (c *Column) Map(v interface{}) *Column {
	switch v.(type) {
	case map[string]string:
		for i, i2 := range v.(map[string]string) {
			c.mapVal[i] = i2
		}
	case map[string]int:
		for i, i2 := range v.(map[string]int) {
			c.mapVal[i] = int64(i2)
		}
	case map[int]string, map[int32]string, map[uint32]string, map[int64]string, map[uint64]string:
		for i, i2 := range switchInt64(v) {
			c.mapVal[i] = i2
		}
	case map[interface{}]interface{}:
		c.mapVal = v.(map[interface{}]interface{})
	default:
		logrus.Error("不支持的map映射类型, map[interface{}]interface{}...", v)
	}
	c.addMapVal()
	return c
}

func switchInt64(i interface{}) map[int64]string {
	m := map[int64]string{}
	switch i.(type) {
	case map[int]string:
		for i2, i3 := range i.(map[int]string) {
			m[int64(i2)] = i3
		}
	case map[int32]string:
		for i2, i3 := range i.(map[int32]string) {
			m[int64(i2)] = i3
		}
	case map[uint32]string:
		for i2, i3 := range i.(map[uint32]string) {
			m[int64(i2)] = i3
		}
	case map[int64]string:
		m = i.(map[int64]string)
	case map[uint64]string:
		for i2, i3 := range i.(map[uint64]string) {
			m[int64(i2)] = i3
		}
	case map[float32]string:
		for i2, i3 := range i.(map[float32]string) {
			m[int64(i2)] = i3
		}
	}
	return m
}
