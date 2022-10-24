package base

import (
	"embed"
	"encoding/json"
	"github.com/go-home-admin/home/app"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
	"strings"
)

//go:embed views
var views embed.FS

func LoadView(file string) string {
	vTable := ""
	if app.IsDebug() {
		if _, err := os.Stat("app/servers/gui/base/views/" + file); err == nil {
			v, _ := ioutil.ReadFile("app/servers/gui/base/views/" + file)
			vTable = string(v)
		} else {
			fileContext, _ := views.ReadFile("views/" + file)
			vTable = string(fileContext)
		}
	} else {
		fileContext, _ := views.ReadFile("views/" + file)
		vTable = string(fileContext)
	}

	return vTable
}

func LoadVue(file string) *Render {
	fileContent := strings.TrimSpace(LoadView(file))

	tStop := strings.LastIndex(fileContent, "\n</template>")
	if tStop == -1 {
		panic("你的模版需要存在<template>")
	}
	template := fileContent[10:tStop]
	render := NewRender(template)

	next := strings.TrimSpace(fileContent[tStop+12:])
	parserScript(next, render)

	return render
}

func parserScript(s string, render *Render) {
	sS := strings.Index(s, "{")
	sE := strings.LastIndex(s, "}")
	if sS == -1 || sE == -1 {
		return
	}
	s = s[sS+1 : sE]

	wl := getWordsWitchFile(s)
	for i := 0; i < len(wl.list); i++ {
		w := wl.list[i]
		switch w.Str {
		case "data":
			var m map[string]interface{}
			m, i = parserData(wl.list[i:], i)
			for s2, i3 := range m {
				render.AddData(s2, i3)
			}
		case "mounted":
			var m string
			m, i = parserMounted(wl.list[i:], i)
			render.AddMounted(render.GetID(), m)
		case "methods":
			var m map[string]string
			m, i = parserMethods(wl.list[i:], i)
			for s2, i3 := range m {
				render.AddMethods(s2, i3)
			}
		}
	}
}

func parserMethods(wl []*word, i int) (map[string]string, int) {
	m := map[string]string{}

	sta, stp := GetBrackets(wl, "{", "}")
	i = i + stp
	nl := wl[sta+1 : stp-1]
	for i2 := 0; i2 < len(nl); i2++ {
		w := nl[i2]
		if w.Str == ":" {
			funName := nl[i2-1]
			nnl := nl[i2:]
			_, stp := GetBrackets(nnl, "{", "}")
			ns := WlToString(nl[i2+1 : i2+stp+1])
			ns = strings.TrimSpace(ns)
			i2 = i2 + stp
			m[funName.Str] = ns
		} else if w.Str == "(" {
			funName := nl[i2-1]
			nnl := nl[i2:]
			_, stp := GetBrackets(nnl, "{", "}")
			ns := WlToString(nl[i2 : i2+stp+1])
			ns = strings.TrimSpace(ns)
			i2 = i2 + stp
			m[funName.Str] = "function" + ns
		}
	}

	return m, i
}

func parserMounted(wl []*word, i int) (string, int) {
	sta, stp := GetBrackets(wl, "{", "}")
	i = i + stp

	nl := wl[sta+1 : stp-1]
	ns := WlToString(nl)
	ns = strings.TrimSpace(ns)

	return ns, i
}

func parserData(wl []*word, i int) (map[string]interface{}, int) {
	sta, stp := GetBrackets(wl, "{", "}")
	i = i + stp

	nl := wl[sta+1 : stp-1]
	ns := WlToString(nl)
	ns = strings.TrimSpace(ns)
	ns = strings.TrimSpace(ns[6:])
	m := map[string]interface{}{}
	err := json.Unmarshal([]byte(ns), &m)
	if err != nil {
		logrus.Error("检查json格式", ns)
		panic(err)
	}
	return m, i
}
