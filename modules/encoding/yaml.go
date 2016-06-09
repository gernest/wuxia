package encoding

import (
	"github.com/robertkrimen/otto"
	"github.com/valor-pw/backend/modules/util"
	"gopkg.in/yaml.v2"
)

type YAML struct {
}

func NewYAML(vm *otto.Otto) otto.Value {
	o := &YAML{}
	y, _ := vm.Object(`({})`)
	y.Set("encode", o.Marshal)
	y.Set("decode", o.Unmarshal)
	return util.ToValue(y)
}

func (y *YAML) Marshal(call otto.FunctionCall) otto.Value {
	arg, err := call.Argument(0).Export()
	if err != nil {
		util.Panic(err)
	}
	rst, err := yaml.Marshal(arg)
	if err != nil {
		util.Panic(err)
	}
	return util.ToValue(string(rst))
}

func (y *YAML) Unmarshal(call otto.FunctionCall) otto.Value {
	arg, err := call.Argument(0).ToString()
	if err != nil {
		util.Panic(err)
	}
	rst := make(map[string]interface{})
	err = yaml.Unmarshal([]byte(arg), rst)
	if err != nil {
		util.Panic(err)
	}
	v, err := call.Otto.ToValue(rst)
	if err != nil {
		util.Panic(err)
	}
	return v
}
