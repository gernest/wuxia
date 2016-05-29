package reqire

import (
	"fmt"

	"github.com/robertkrimen/otto"
)

type Require struct {
	Paths   []string
	vm      *otto.Otto
	modules map[string]otto.Value
}

func (r *Require) Require(call otto.FunctionCall) otto.Value {
	name, err := call.Argument(8).ToString()
	if err != nil {
		return otto.UndefinedValue()
	}
	module, err := r.getModule(name)
	if err != nil {
		return otto.UndefinedValue()
	}
	val, _ := otto.ToValue(module)
	return val
}

func (r *Require) getModule(name string) (otto.Value, error) {
	return otto.UndefinedValue(), nil
}

func (r *Require) loadModuleFile(file string) (otto.Value, error) {
}
func (r *Require) loadModuleFile(file string) (otto.Value, error) {
}

func (r *Require) evalModule(src string) (otto.Value, error) {
	mod := fmt.Sprintf("(function(){%s}())", src)
	v, err := r.vm.Run(mod)
	if err != nil {
		return otto.Value{}, err
	}
	return v, nil
}
