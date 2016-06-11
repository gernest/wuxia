package vm

import "github.com/robertkrimen/otto"

type VM struct {
	o *otto.Otto
}

func NewVM(req *Require) *VM {
	o := otto.New()
	o.Set("require", req.ToValue())
	return &VM{o: o}
}
