package vm

import (
	"bytes"
	"errors"
	"io"
	"path/filepath"

	"github.com/robertkrimen/otto"
	"github.com/spf13/afero"
)

type File struct {
	cache   map[string]otto.Value
	vm      *otto.Otto
	require func(otto.FunctionCall) otto.Value
	isInit  bool
	paths   []string
	fs      afero.Fs
	ext     []string
}

func (f *File) IsInit() bool {
	return f.isInit
}

func (f *File) Init(vm *otto.Otto, require func(otto.FunctionCall) otto.Value) {
	if f.IsInit() {
		return
	}
	f.vm = vm
	f.require = require
	if f.cache == nil {
		f.cache = make(map[string]otto.Value)
	}
}

func (f *File) tryFind(pwd, name string) ([]byte, error) {
	if filepath.IsAbs(name) {
		return readFile(f.fs, name, f.ext...)
	}
	var lookup []string
	lookup = append(lookup, pwd)
	lookup = append(lookup, f.paths...)
	for _, dir := range lookup {
		p := filepath.Join(dir, name)
		data, err := readFile(f.fs, p, f.ext...)
		if err != nil {
			continue
		}
		return data, nil
	}
	return nil, errors.New("nothing found")
}
func readFile(fs afero.Fs, name string, ext ...string) ([]byte, error) {
	if filepath.Ext(name) == "" {
		found := false
		for _, e := range ext {
			n := name + "." + e
			_, err := fs.Stat(n)
			if err != nil {
				continue
			}
			found = true
			name = name + "." + e
			break
		}
		if !found {
			return nil, errors.New("not found")
		}
	}
	var b bytes.Buffer
	f, err := fs.Open(name)
	if err != nil {
		return nil, err
	}
	defer func() { f.Close() }()
	_, err = io.Copy(&b, f)
	if err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (f *File) Load(name, pwd string) (otto.Value, bool) {
	// The following code is addopted from
	//https://github.com/ddliu/motto
	if !filepath.IsAbs(name) {
		name = filepath.Clean(name)
	}
	if m, ok := f.cache[name]; ok {
		return m, ok
	}
	data, err := f.tryFind(pwd, name)
	if err != nil {
		Panic(err)
	}

	v, err := f.loadFromSource(string(data))
	if err != nil {
		Panic(err)
	}
	f.cache[name] = v
	return v, true
}

func (f *File) loadFromSource(src string) (otto.Value, error) {
	// The following code is addopted from
	//https://github.com/ddliu/motto
	source := "(function(module) {var require = module.require;var exports = module.exports;\n" + src + "\n})"
	module, err := f.vm.Object(`({exports: {}})`)
	if err != nil {
		return otto.UndefinedValue(), err
	}
	module.Set("require", f.require)
	exports, _ := module.Get("exports")
	val, err := f.vm.Call(source, exports, module)
	if err != nil {
		return otto.UndefinedValue(), err
	}
	if !val.IsUndefined() {
		return val, nil
	}
	expV, err := module.Get("exports")
	if err != nil {
		return otto.UndefinedValue(), err
	}
	return expV, nil
}
