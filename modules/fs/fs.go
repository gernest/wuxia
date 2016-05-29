package fs

import (
	"os"

	"github.com/robertkrimen/otto"
)

type FS interface {
	Pwd() (Dir, error)
	OpenFile(string) (File, error)
}

type File interface {
	Read() (string, error)
}

type Dir interface {
	Open() ([]File, error)
}

func FileInfo(vm *otto.Otto, info os.FileInfo) otto.Value {
	o, _ := vm.Object(`({})`)
	o.Set("name", func(call otto.FunctionCall) otto.Value {
		return util.ToValue(info.Name())
	})
	o.Set("isDir", func(call otto.FunctionCall) otto.Value {
		return util.ToValue(info.IsDir())
	})
	o.Set("modTime", func(call otto.FunctionCall) otto.Value {
		return util.ToValue(info.ModTime())
	})
	o.Set("size", func(call otto.FunctionCall) otto.Value {
		return util.ToValue(info.Size())
	})
}
