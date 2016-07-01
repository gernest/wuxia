package gen

import (
	"io/ioutil"
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
	ff, _ := fs.Create("hello.txt")
	ff.WriteString("hello")
	ff.Close()

	script, err := ioutil.ReadFile("fixture/test/open.js")
	if err != nil {
		t.Fatal(err)
	}
	_, err = vm.Run(script)
	if err != nil {
		t.Error(err)
	}

}
