package vm

import "github.com/robertkrimen/otto"

type VM struct {
	o *otto.Otto
}

func (vm *VM) Exec(pwd string, src string) (otto.Value, error) {
	return otto.UndefinedValue(), nil
}

func (vm *VM) ExecFile(pwd string, src string) (otto.Value, error) {
	return otto.UndefinedValue(), nil
}

func (vm *VM) Eval(pwd string, src string) (otto.Value, error) {
	return otto.UndefinedValue(), nil
}

func (vm *VM) EvalFile(pwd string, src string) (otto.Value, error) {
	return otto.UndefinedValue(), nil
}

func NewVM(req *Require) *VM {
	o := otto.New()
	o.Set("require", req.ToValue())
	return &VM{o: o}
}
