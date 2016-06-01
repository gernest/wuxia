package fs

import (
	"testing"

	"github.com/robertkrimen/otto"
)

func TestFS(t *testing.T) {
	vm := otto.New()
	vm.Set("error", func(call otto.FunctionCall) otto.Value {
		name, _ := call.Argument(0).ToString()
		t.Error(name)
		return otto.UndefinedValue()
	})
	vm.Set("FS", NewFS(vm))
	var fsTest = `
try{
// Open a new file
name="sample.txt";
FS.writeFile(name,"");
var f=FS.open("sample.txt");
var msg="hello";
f.write(msg);
f.flush();
}catch(e){
	error(e);
}
`
	_, err := vm.Eval(fsTest)
	if err != nil {
		t.Error(err)
	}
}
