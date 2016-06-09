package export

import "github.com/robertkrimen/otto"

func New() Export {
	return make(Export)
}

type Export map[string]interface{}

func (e Export) Set(key string, value interface{}) {
	e[key] = value
}

func (e Export) Register(vm *otto.Otto) error {
	o, err := vm.Object(`({})`)
	if err != nil {
		return err
	}
	for key, value := range e {
		o.Set(key, value)
	}
	return nil
}
