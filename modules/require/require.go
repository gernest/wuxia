package require

import (
	"github.com/gernest/valeria/modules/util"
	"github.com/robertkrimen/otto"
)

type ModuleLoader interface {
	Init(*otto.Otto, func(otto.FunctionCall) otto.Value)
	IsInit() bool
	Load(name string) (otto.Value, bool)
}

func New(loaders ...ModuleLoader) *Require {
	return &Require{l: loaders}
}

type Require struct {
	l []ModuleLoader
}

func (r *Require) ToValue() func(otto.FunctionCall) otto.Value {
	return r.require
}

func (r *Require) require(call otto.FunctionCall) otto.Value {
	name, err := call.Argument(0).ToString()
	if err != nil {
		util.Panic(err)
	}
	return r.findModule(call.Otto, name)
}
func (r *Require) findModule(vm *otto.Otto, name string) otto.Value {
	for i := 0; i < len(r.l); i++ {
		if !r.l[i].IsInit() {
			r.l[i].Init(vm, r.require)
		}
		if m, ok := r.l[i].Load(name); ok {
			return m
		}
	}
	return otto.UndefinedValue()
}
