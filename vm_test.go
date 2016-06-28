package gen

import (
	"testing"

	"github.com/robertkrimen/otto"
	"github.com/spf13/afero"
)

func TestFs(t *testing.T) {
	vm := otto.New()
	fs := afero.NewMemMapFs()
	vm.Set("newFS", func(call otto.FunctionCall) otto.Value {
		f := &fileSys{}
		f.Fs = fs
		f.vm = vm
		return f.export().ToValue(call.Otto)
	})

	var script = `
var fs=newFS();
function testOpen(){
	try{
		f=fs.open("hello.txt");
		f.write("hello");
		f.close();
	}catch(e){
		throw e;
	}
}
testOpen();

`
	_, err := vm.Run(script)
	if err != nil {
		t.Error(err)
	}

}
