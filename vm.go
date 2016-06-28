package gen

import (
	"io/ioutil"

	"github.com/robertkrimen/otto"
	"github.com/spf13/afero"
)

type VM struct {
	*otto.Otto
}

type Export map[string]interface{}

func (e Export) Set(key string, value interface{}) {
	e[key] = value
}

func (e Export) ToValue(vm *otto.Otto) otto.Value {
	o, err := vm.Object(`({})`)
	if err != nil {
		panic(err)
	}
	for key, value := range e {
		o.Set(key, value)
	}
	return o.Value()
}

const (
	modeRead  = "r"
	modeWrite = "w"
)

type fileSys struct {
	afero.Fs
	vm *otto.Otto
}

func (fs fileSys) export() Export {
	e := make(Export)
	e.Set("open", fs.open)
	return e
}

func (fs fileSys) open(call otto.FunctionCall) otto.Value {
	name, err := call.Argument(0).ToString()
	if err != nil {
		Panic(err)
	}
	f, err := fs.Open(name)
	if err != nil {
		Panic(err)
	}
	af := &file{o: f}
	return af.export().ToValue(call.Otto)
}

type file struct {
	o afero.File
}

func (f *file) export() Export {
	e := make(Export)
	e.Set("close", f.close)
	e.Set("read", f.read)
	return e
}

func (f *file) close(call otto.FunctionCall) otto.Value {
	err := f.o.Close()
	if err != nil {
		Panic(err)
	}
	return otto.Value{}
}

func (f *file) read(call otto.FunctionCall) otto.Value {
	b, err := ioutil.ReadAll(f.o)
	if err != nil {
		Panic(err)
	}
	return ToValue(string(b))
}
