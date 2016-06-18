package vm

import "github.com/robertkrimen/otto"

type ModuleLoader interface {
	Init(*otto.Otto, func(otto.FunctionCall) otto.Value)
	IsInit() bool
	Load(name, pwd string) (otto.Value, bool)
}

func NewRequre(loaders ...ModuleLoader) *Require {
	return &Require{l: loaders}
}

type Require struct {
	l   []ModuleLoader
	pwd string
}

func (r *Require) ToValue() func(otto.FunctionCall) otto.Value {
	return r.require
}

func (r *Require) require(call otto.FunctionCall) otto.Value {
	name, err := call.Argument(0).ToString()
	if err != nil {
		Panic(err)
	}
	return r.findModule(call.Otto, name)
}
func (r *Require) findModule(vm *otto.Otto, name string) otto.Value {
	for i := 0; i < len(r.l); i++ {
		if !r.l[i].IsInit() {
			r.l[i].Init(vm, r.require)
		}
		if m, ok := r.l[i].Load(name, r.pwd); ok {
			return m
		}
	}
	return otto.UndefinedValue()
}

func (r *Require) SetWorkingDir(dir string) {
	r.pwd = dir
}

type ValueLoader struct {
	cache  map[string]otto.Value
	isInit bool
}

func (v *ValueLoader) Init(*otto.Otto, func(otto.FunctionCall) otto.Value) {
	if !v.isInit {
		v.isInit = true
	}
}

func (v *ValueLoader) IsInit() bool {
	return v.isInit
}
func (v *ValueLoader) Load(name, pwd string) (otto.Value, bool) {
	val, ok := v.cache[name]
	return val, ok
}
