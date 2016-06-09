package base

import (
	"errors"

	"github.com/ddliu/motto"
	"github.com/robertkrimen/otto"
)

type VM struct {
	vm      *motto.Motto
	builtIn map[string]otto.Value
}

func (v *VM) AddBuiltinModule(name string, value otto.Value) {
	v.vm.AddModule(name, v.loader(name))
}

func (v *VM) loader(name string) func(*motto.Motto) (otto.Value, error) {
	return func(vm *motto.Motto) (otto.Value, error) {
		if m, ok := v.builtIn[name]; ok {
			return m, nil
		}
		return otto.Value{}, errors.New("module not found")
	}
}
