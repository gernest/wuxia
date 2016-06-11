package loaders

import (
	"io/ioutil"
	"path/filepath"

	"github.com/gernest/valeria/modules/util"
	"github.com/robertkrimen/otto"
)

type File struct {
	cache   map[string]otto.Value
	vm      *otto.Otto
	require func(otto.FunctionCall) otto.Value
	isInit  bool
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

func (f *File) Load(name, pwd string) (otto.Value, bool) {
	// The following code is addopted from
	//https://github.com/ddliu/motto
	if !filepath.IsAbs(name) {
		name = filepath.Clean(name)
		name = filepath.Join(pwd, name)
	}
	if m, ok := f.cache[name]; ok {
		return m, ok
	}
	data, err := ioutil.ReadFile(name)
	if err != nil {
		util.Panic(err)
	}

	v, err := f.loadFromSource(string(data))
	if err != nil {
		util.Panic(err)
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
	expVl, err := module.Get("exports")
	if err != nil {
		return otto.UndefinedValue(), err
	}
	return expVl, nil
}
