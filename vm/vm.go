package vm

import "github.com/robertkrimen/otto"

type VM struct {
	*otto.Otto
}

func NewVM(req *Require) *VM {
	o := &VM{}
	o.Otto = otto.New()
	o.Set("require", req.ToValue())
	return o
}
