package encoding

import (
	"testing"

	"github.com/gernest/valeria/modules/util"
	"github.com/robertkrimen/otto"
)

func TestYAML_Marhsal(t *testing.T) {
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
		console.log(o);
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
