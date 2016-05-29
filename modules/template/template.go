package template

import (
	"bytes"
	"html/template"

	"github.com/gernest/valeria/modules/util"
	"github.com/robertkrimen/otto"
)

type Template struct {
	tpl *template.Template
}

func New(name string) *Template {
	return &Template{
		tpl: template.New(name),
	}
}

func NewObject(vm *otto.Otto, name string) otto.Value {
	tpl, _ := vm.Object(`({})`)
	t := New(name)
	tpl.Set("definedTemplates", t.DefinedTemplates)
	tpl.Set("name", t.Name)
	tpl.Set("delims", t.Delimgs)
	tpl.Set("execute", t.Execute)
	tpl.Set("executeTemplate", t.ExecuteTemplate)
	tpl.Set("parse", t.Parse)
	tpl.Set("option", t.Option)
	return util.ToValue(tpl)
}

func (t *Template) DefinedTemplates(call otto.FunctionCall) otto.Value {
	return util.ToValue(t.tpl.DefinedTemplates())
}

func (t *Template) Name(call otto.FunctionCall) otto.Value {
	return util.ToValue(t.tpl.Name())
}

func (t *Template) Delimgs(call otto.FunctionCall) otto.Value {
	left, _ := call.Argument(1).ToString()
	right, _ := call.Argument(2).ToString()
	tpl := t.tpl.Delims(left, right)
	t.tpl = tpl
	return util.ToValue(t)
}

func (t *Template) Execute(call otto.FunctionCall) otto.Value {
	data, _ := call.Argument(0).Export()
	rst := &bytes.Buffer{}
	err := t.tpl.Execute(rst, data)
	if err != nil {
		util.Panic(err)
	}
	return util.ToValue(rst.String())
}

func (t *Template) ExecuteTemplate(call otto.FunctionCall) otto.Value {
	name, _ := call.Argument(0).ToString()
	data, _ := call.Argument(1).Export()
	rst := &bytes.Buffer{}
	err := t.tpl.ExecuteTemplate(rst, name, data)
	if err != nil {
		util.Panic(err)
	}
	return util.ToValue(rst.String())
}

func (t *Template) Parse(call otto.FunctionCall) otto.Value {
	src, _ := call.Argument(0).ToString()
	tpl, err := t.tpl.Parse(src)
	if err != nil {
		util.Panic(err)
	}
	t.tpl = tpl
	return util.ToValue(t)
}

func (t *Template) Option(call otto.FunctionCall) otto.Value {
	var opts []string
	if len(call.ArgumentList) > 0 {
		for _, args := range call.ArgumentList {
			v, err := args.ToString()
			if err != nil {
				util.Panic(err)
			}
			opts = append(opts, v)
		}
		t.tpl.Option(opts...)
	}
	return util.ToValue(t)
}
