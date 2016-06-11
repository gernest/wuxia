package vm

import "github.com/robertkrimen/otto"

func NewExport() Export {
	return make(Export)
}

type Export map[string]interface{}

func (e Export) Set(key string, value interface{}) {
	e[key] = value
}

func (e Export) ToValue(vm *otto.Otto) otto.Value {
	o, err := vm.Object(`({})`)
	if err != nil {
		panic(err)
	}
	for key, value := range e {
		o.Set(key, value)
	}
	return o.Value()
}
