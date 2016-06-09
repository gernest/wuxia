package export

import "github.com/robertkrimen/otto"

func New() Export {
	return make(Export)
}

type Export map[string]func(otto.FunctionCall) otto.Value

func (e Export) Set(key string, value func(otto.FunctionCall) otto.Value) {
	e[key] = value
}
