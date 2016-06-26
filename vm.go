package gen

import "github.com/robertkrimen/otto"

type VM struct {
	*otto.Otto
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
func Panic(o interface{}) {
	v, err := otto.ToValue(o)
	if err != nil {
		errV, _ := otto.ToValue(err)
		panic(errV)
	}
	panic(v)
}

func ToValue(o interface{}) otto.Value {
	v, _ := otto.ToValue(o)
	return v
}
