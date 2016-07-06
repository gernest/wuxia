package gen

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/robertkrimen/otto"
	"github.com/spf13/afero"
)

const (
	msgModduleNotFound = "module %s not found"
)

type require struct {
	cache map[string]otto.Value
	paths []string
	fs    afero.Fs
}

type moduleInfo struct {
	id string
}

func (r *require) load(call otto.FunctionCall) otto.Value {
	id, err := call.Argument(0).ToString()
	if err != nil {
		Panic(err)
	}
	newID, err := r.resolve(id)
	if err != nil {
		Panic(err)
	}
	return r.loadFromFile(newID, call.Otto)
}

func (r *require) resolve(id string) (string, error) {
	if id == "" {
		return "", errors.New("empty module name")
	}
	if filepath.IsAbs(id) {
		return id, nil
	}
	if strings.HasPrefix(id, ".") {
		_, err := r.fs.Stat(id)
		if err != nil {
			return "", err
		}
	}
	ext := filepath.Ext(id)
	opts := []string{".js", ".json"}
	for i := 0; i < len(r.paths); i++ {
		fullPath := filepath.Join(r.paths[i], id)
		if ext != "" {
			_, err := r.fs.Stat(fullPath)
			if err != nil {
				return "", err
			}
			return fullPath, nil
		}

		for _, e := range opts {
			_, err := r.fs.Stat(fullPath + e)
			if err != nil {
				return "", err
			}
			return fullPath + e, nil
		}
	}
	return "", fmt.Errorf(msgModduleNotFound, id)
}

func (r *require) loadFromFile(path string, vm *otto.Otto) otto.Value {
	f, err := r.fs.Open(path)
	if err != nil {
		Panic(err.Error())
	}
	defer func() { _ = f.Close() }()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		Panic(err.Error())
	}
	if filepath.Ext(path) == ".json" {
		v, err := vm.Call("JSON.parse", nil, string(data))
		if err != nil {
			Panic(err.Error())
		}
		return v
	}
	return r.loadFromSource(string(data), vm)
}

func (r *require) loadFromSource(source string, vm *otto.Otto) otto.Value {
	source = "(function(module) {var require = module.require;var exports = module.exports;\n" + source + "\n})"

	jsModule, _ := vm.Object(`({exports: {}})`)
	jsModule.Set("require", r.load)
	jsExports, _ := jsModule.Get("exports")

	moduleReturn, err := vm.Call(source, jsExports, jsModule)
	if err != nil {
		Panic(err.Error())
	}
	var moduleValue otto.Value
	if !moduleReturn.IsUndefined() {
		moduleValue = moduleReturn
		jsModule.Set("exports", moduleValue)
	} else {
		moduleValue, _ = jsModule.Get("exports")
	}
	return moduleValue
}
