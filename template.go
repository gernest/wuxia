package gen

import (
	"fmt"
	"html/template"

	"github.com/robertkrimen/otto"
)

type Template struct {
	vm      *otto.Otto
	jsFuncs []string
	funcs   template.FuncMap
	*template.Template
}

func (t *Template) funcMap() template.FuncMap {
	rst := make(template.FuncMap)
	for _, name := range t.jsFuncs {
		rst[name] = t.jsTplFunc(name)
	}
	for k, v := range t.funcs {
		rst[k] = v
	}
	return rst
}

func (t *Template) jsTplFunc(name string) func(interface{}) string {
	return func(arg interface{}) string {
		call := fmt.Sprintf("Tpl.funcs.%s", name)
		rst, err := t.vm.Call(call, nil, arg)
		if err != nil {
			Panic(err.Error())
		}
		if !rst.IsString() {
			Panic("non string retrun value from " + name + " template func")
		}
		s, _ := rst.ToString()
		return s
	}
}

func (t *Template) New() *Template {
	tpl := template.New("base").Funcs(t.funcMap())
	return &Template{
		jsFuncs:  t.jsFuncs,
		vm:       t.vm,
		funcs:    t.funcs,
		Template: tpl,
	}
}
