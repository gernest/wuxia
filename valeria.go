package valeria

import (
	"errors"

	"github.com/ddliu/motto"
	"github.com/robertkrimen/otto"
)

type Valeria struct {
	vm      *motto.Motto
	builtIn map[string]otto.Value
}

func (v *Valeria) AddBuiltinModule(name string, value otto.Value) {
	v.vm.AddModule(name, v.loader(name))
}

func (v *Valeria) loader(name string) func(*motto.Motto) (otto.Value, error) {
	return func(vm *motto.Motto) (otto.Value, error) {
		if m, ok := v.builtIn[name]; ok {
			return m, nil
		}
		return otto.Value{}, errors.New("module not found")
	}
}
