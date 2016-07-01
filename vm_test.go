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
		return f.export().ToValue(call.Otto)
	})
	ff, _ := fs.Create("hello.txt")
	ff.WriteString("hello")
	ff.Close()

	scrippts := []struct {
		name string
		path string
	}{
		{"open", "fixture/test/open.js"},
		{"open_file", "fixture/test/open_file.js"},
	}

	for i := 0; i < len(scrippts); i++ {
		script, err := ioutil.ReadFile(scrippts[i].path)
		if err != nil {
			t.Fatal(err)
		}
		_, err = vm.Run(script)
		if err != nil {
			t.Errorf("%s : %v", scrippts[i].name, err)
		}
	}
}
