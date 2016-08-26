package gen

import (
	"errors"
	"fmt"
	"html/template"

	"github.com/robertkrimen/otto"
)

//Template extends the *template.Template struct by supporting template
//functions defined in the javascript programming language. This depends heavily
//on the otto VM , and the functions are extracted from the otto Virtual
//Machine.
type Template struct {
	*otto.Otto
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

// returns a JS function that can be executed within Go template. The defined
// javascript function should accept one argument and return a string. Any
// exceptions raised by the javascript function will we returned when the
// templates are executed( Effectively halting rendering process).
//
// The functions are the one registered on the Tpl global pbject exposed in the
// Javascript runtine.
//
// You can register a function by attaching it to  Tpl.funcs
//   // Example
//   Tpl.funcs.world=function(hello){return hello+",world"}
// The xample function adds ",world" string to the passed argument. The above
// function can then be used like any other Go template functions like this
//   {{"hello"|world}}
//
// The type of the argument is not enforced so the template function
// implementations should be careful on what type of objects they are operating
// on and also great care should be taken on the conext object passed to these
// functions within the templates.
func (t *Template) jsTplFunc(name string) func(interface{}) (string, error) {
	return func(arg interface{}) (string, error) {
		call := fmt.Sprintf("Tpl.funcs.%s", name)
		rst, err := t.Call(call, nil, arg)
		if err != nil {
			return "", err
		}
		if !rst.IsString() {
			return "", errors.New("non string retrun value from " + name + " template func")
		}
		s, _ := rst.ToString()
		return s, nil
	}
}

//New created a new *Template. The returned *Template supports template fuctions
//defined in javascript.
func (t *Template) New() *Template {
	if t.jsFuncs == nil || len(t.jsFuncs) == 0 {
		if t.Otto != nil {
			rst, err := t.Call("Tpl.getTplFuncs", nil)
			if err == nil {
				v, _ := rst.Export()
				if va, ok := v.([]string); ok {
					t.jsFuncs = va
				}
			}
		}
	}
	tpl := template.New("base").Funcs(t.funcMap())
	return &Template{
		jsFuncs:  t.jsFuncs,
		Otto:     t.Otto,
		funcs:    t.funcs,
		Template: tpl,
	}
}
