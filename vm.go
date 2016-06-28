package gen

import (
	"errors"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

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
	modCreate = "c"
	modAppend = "a"
	modTrucc  = "t"
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

func buildFlags(src string) (int, error) {
	parts := strings.Split(src, "+")
	if len(parts) > 0 {
		var f int
		for i := 0; i < len(parts); i++ {
			switch parts[i] {
			case modeRead:
				f = f | os.O_RDONLY
			case modeWrite:
				f = f | os.O_WRONLY
			case modCreate:
				f = f | os.O_CREATE
			case modTrucc:
				f = f | os.O_TRUNC
			default:
				return f, errors.New("unknown flag " + parts[i])
			}
		}
	}
	return 0, errors.New("no flags found")
}

func (fs fileSys) openFile(call otto.FunctionCall) otto.Value {
	name, err := call.Argument(0).ToString()
	if err != nil {
		Panic(err)
	}
	flag, err := call.Argument(1).ToString()
	if err != nil {
		Panic(err)
	}
	uflag, err := buildFlags(flag)
	if err != nil {
		Panic(err)
	}
	mode, err := call.Argument(2).ToString()
	if err != nil {
		Panic(err)
	}
	umode, err := strconv.ParseUint(mode, 10, 32)
	if err != nil {
		Panic(err)
	}
	f, err := fs.OpenFile(name, uflag, os.FileMode(umode))
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

func (f *file) write(call otto.FunctionCall) otto.Value {
	c, err := call.Argument(0).ToString()
	if err != nil {
		Panic(err)
	}
	n, err := f.o.WriteString(c)
	if err != nil {
		Panic(err)
	}
	return ToValue(n)
}
