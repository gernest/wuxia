package wuxia

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"

	"github.com/muesli/cache2go"
	"github.com/robertkrimen/otto"
	"github.com/spf13/afero"
)

const (
	msgModduleNotFound = "module %s not found"
)

// implements a simple node.js like require mechanism. It loads jsavascript
// files from the source and exposes them as modules in the runtime.
//
// This follows node.js convention by using the exports object to attach the
// functions or objects exposed by the module.
//
// NOTE: cyclic dependencies are not taken care yet, so this will break in case
// of cyclic dependency.
type require struct {
	cache *cache2go.CacheTable
	paths []string
	fs    afero.Fs
}

func newRequire(fs afero.Fs, paths ...string) *require {
	return &require{
		cache: cache2go.Cache("require"),
		paths: paths,
		fs:    fs,
	}
}

func (r *require) load(call otto.FunctionCall) otto.Value {
	id, err := call.Argument(0).ToString()
	if err != nil {
		Panic(err)
	}
	newID, err := r.resolve(id)
	if err != nil {
		Panic(err.Error())
	}
	if cached, ok := r.checkCache(newID); ok {
		return cached
	}
	return r.loadFromFile(newID, call.Otto)
}

func (r *require) checkCache(id string) (otto.Value, bool) {
	if !r.cache.Exists(id) {
		return otto.Value{}, false
	}
	res, err := r.cache.Value(id)
	if err != nil {
		return otto.Value{}, false
	}
	return res.Data().(otto.Value), true
}

func (r *require) addToCache(id string, v otto.Value) {
	_ = r.cache.Add(id, time.Minute, v)
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
				if i != len(r.paths)-1 {
					continue
				}
				return "", err
			}
			return fullPath, nil
		}

		for _, e := range opts {
			_, err := r.fs.Stat(fullPath + e)
			if err != nil {
				if i != len(r.paths)-1 {
					continue
				}
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
	return r.loadFromSource(string(data), path, vm)
}

func (r *require) loadFromSource(source string, path string, vm *otto.Otto) otto.Value {
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
	r.addToCache(path, moduleValue)
	return moduleValue
}
