package encoding

import (
	"testing"

	"github.com/robertkrimen/otto"
	"github.com/valor-pw/backend/modules/util"
)

func TestYAML(t *testing.T) {
	vm := otto.New()
	vm.Set("YAML", func(_ otto.FunctionCall) otto.Value {
		return util.ToValue(NewYAML(vm))
	})
	vm.Set("error", func(call otto.FunctionCall) otto.Value {
		name, _ := call.Argument(0).ToString()
		t.Error(name)
		return otto.UndefinedValue()
	})
	var marshalJS = `
function  TestYAMLMarhsal(){
	e =new YAML();
	try{
		v={};
		v.title="hello";
		var o=e.encode(v);
	}catch(err){
		error(err);
	}

}
TestYAMLMarhsal();
`
	_, err := vm.Eval(marshalJS)
	if err != nil {
		t.Error(err)
	}
	var unmarshalJS = `
function  TestYAMLUnMarhsal(){
	var e =new YAML();
	try{
		src='title: hello';
		var o=e.decode(src);
		if(o.title!="hello"){
			error("failed to decode");
		}
	}catch(err){
		error(err);
	}

}
TestYAMLUnMarhsal();
`
	_, err = vm.Run(unmarshalJS)
	if err != nil {
		t.Error(err)
	}
}
