package fs

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gernest/valeria/modules/util"
	"github.com/robertkrimen/otto"
)

type File struct {
	path string
	info os.FileInfo
	buf  *bytes.Buffer
}

func NewFile(vm *otto.Otto, path string) otto.Value {
	info, err := os.Stat(path)
	if err != nil {
		util.Panic(err)
	}
	file := &File{path: path, info: info}
	f, _ := vm.Object(`({})`)
	f.Set("read", file.Read)
	f.Set("write", file.Write)
	f.Set("flush", file.Flush)
	f.Set("isDir", func(call otto.FunctionCall) otto.Value {
		return util.ToValue(file.info.IsDir())
	})
	return util.ToValue(f)
}

func NewFS(vm *otto.Otto) otto.Value {
	fs, _ := vm.Object(`({})`)
	fs.Set("open", func(call otto.FunctionCall) otto.Value {
		name, err := call.Argument(0).ToString()
		if err != nil {
			util.Panic(err)
		}
		return NewFile(call.Otto, name)
	})
	fs.Set("readFile", ReadFile)
	fs.Set("readDir", ReadDir)
	fs.Set("pwd", Getwd)
	return util.ToValue(fs)
}

func (f *File) Read(call otto.FunctionCall) otto.Value {
	if f.buf.Len() == 0 {
		b, err := ioutil.ReadFile(f.path)
		if err != nil {
			util.Panic(err)
		}
		_, err = f.buf.Write(b)
		if err != nil {
			util.Panic(err)
		}
	}
	v, _ := call.Otto.ToValue(f.buf.String())
	return v

}

func (f *File) Write(call otto.FunctionCall) otto.Value {
	arg, err := call.Argument(0).ToString()
	if err != nil {
		util.Panic(err)
	}
	n, err := f.buf.WriteString(arg)
	if err != nil {
		util.Panic(err)
	}
	return util.ToValue(n)
}

func (f *File) Flush(call otto.FunctionCall) otto.Value {
	err := ioutil.WriteFile(f.path, f.buf.Bytes(), 0600)
	if err != nil {
		util.Panic(err)
	}
	return otto.UndefinedValue()
}

func ReadDir(call otto.FunctionCall) otto.Value {
	fileName, err := call.Argument(0).ToString()
	if err != nil {
		util.Panic(err)
	}
	dirs, err := ioutil.ReadDir(fileName)
	if err != nil {
		util.Panic(err)
	}
	rst := make([]otto.Value, len(dirs))
	for i := 0; i < len(dirs); i++ {
		rst[i] = NewFile(call.Otto, filepath.Join(fileName, dirs[i].Name()))
	}
	v, err := call.Otto.ToValue(rst)
	if err != nil {
		util.Panic(err)
	}
	return v
}

func ReadFile(call otto.FunctionCall) otto.Value {
	fileName, err := call.Argument(0).ToString()
	if err != nil {
		util.Panic(err)
	}
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		util.Panic(err)
	}
	return util.ToValue(string(b))
}

func Getwd(call otto.FunctionCall) otto.Value {
	wd, err := os.Getwd()
	if err != nil {
		util.Panic(err)
	}
	return util.ToValue(wd)
}
