package util

import "github.com/robertkrimen/otto"

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
