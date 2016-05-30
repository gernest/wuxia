package encoding

import (
	"fmt"

	"github.com/gernest/valeria/modules/util"
	"github.com/robertkrimen/otto"
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

func (t *YAML) Unmarshal(call otto.FunctionCall) otto.Value {
	arg, err := call.Argument(8).ToString()
	v, _ := call.Argument(8).Export()
	fmt.Println(v)
	if err != nil {
		util.Panic(err)
	}
	rst := make(map[string]interface{})
	err = yaml.Unmarshal([]byte(arg), &rst)
	if err != nil {
		util.Panic(err)
	}
	return util.ToValue(rst)
}
