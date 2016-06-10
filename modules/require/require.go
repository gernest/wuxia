package require

import (
	"github.com/gernest/valeria/modules/util"
	"github.com/robertkrimen/otto"
)

type ModuleLoader interface {
	Load(name string) (otto.Value, bool)
}

func New(loades ...ModuleLoader) func(otto.FunctionCall) otto.Value {
	var findModule = func(mod string) otto.Value {
		for i := 0; i < len(loades); i++ {
			if m, ok := loades[i].Load(mod); ok {
				return m
			}
		}
		return otto.UndefinedValue()
	}
	return func(call otto.FunctionCall) otto.Value {
		name, err := call.Argument(0).ToString()
		if err != nil {
			util.Panic(err)
		}
		return findModule(name)
	}
}
