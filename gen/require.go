package gen

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

//Require  implements a simple node.js like require mechanism. It loads jsavascript
// files from the source and exposes them as modules in the runtime.
//
// This follows node.js convention by using the exports object to attach the
// functions or objects exposed by the module.
//
// NOTE: cyclic dependencies are not taken care yet, so this will break in case
// of cyclic dependency.
type Require struct {
	Cache *cache2go.CacheTable
	Paths []string
	Fs    afero.Fs
}

//NewRequire returns an initialized Require object
func NewRequire(fs afero.Fs, paths ...string) *Require {
	return &Require{
		Cache: cache2go.Cache("require"),
		Paths: paths,
		Fs:    fs,
	}
}

//Load  loads module into the otto runtime.
func (r *Require) Load(call otto.FunctionCall) otto.Value {
	id, err := call.Argument(0).ToString()
	if err != nil {
		panicOtto(err)
	}
	if cached, ok := r.checkCache(id); ok {
		return cached
	}
	newID, err := r.resolve(id)
	if err != nil {
		panicOtto(err.Error())
	}
	if cached, ok := r.checkCache(newID); ok {
		return cached
	}
	return r.loadFromFile(newID, call.Otto)
}

func (r *Require) checkCache(id string) (otto.Value, bool) {
	if !r.Cache.Exists(id) {
		return otto.Value{}, false
	}
	res, err := r.Cache.Value(id)
	if err != nil {
		return otto.Value{}, false
	}
	return res.Data().(otto.Value), true
}

func (r *Require) addToCache(id string, v otto.Value) {
	_ = r.Cache.Add(id, time.Minute, v)
}

func (r *Require) resolve(id string) (string, error) {
	if id == "" {
		return "", errors.New("empty module name")
	}
	if filepath.IsAbs(id) {
		return id, nil
	}
	if strings.HasPrefix(id, ".") {
		_, err := r.Fs.Stat(id)
		if err != nil {
			return "", err
		}
	}
	ext := filepath.Ext(id)
	opts := []string{".js", ".json"}
	for i := 0; i < len(r.Paths); i++ {
		fullPath := filepath.Join(r.Paths[i], id)
		if ext != "" {
			_, err := r.Fs.Stat(fullPath)
			if err != nil {
				if i != len(r.Paths)-1 {
					continue
				}
				return "", err
			}
			return fullPath, nil
		}

		for _, e := range opts {
			_, err := r.Fs.Stat(fullPath + e)
			if err != nil {
				if i != len(r.Paths)-1 {
					continue
				}
				return "", err
			}
			return fullPath + e, nil
		}
	}
	return "", fmt.Errorf(msgModduleNotFound, id)
}

func (r *Require) loadFromFile(path string, vm *otto.Otto) otto.Value {
	f, err := r.Fs.Open(path)
	if err != nil {
		panicOtto(err.Error())
	}
	defer func() { _ = f.Close() }()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panicOtto(err.Error())
	}
	if filepath.Ext(path) == ".json" {
		v, err := vm.Call("JSON.parse", nil, string(data))
		if err != nil {
			panicOtto(err.Error())
		}
		return v
	}
	return r.loadFromSource(string(data), path, vm)
}

func (r *Require) loadFromSource(source string, path string, vm *otto.Otto) otto.Value {
	source = "(function(module) {var require = module.require;var exports = module.exports;\n" + source + "\n})"

	jsModule, _ := vm.Object(`({exports: {}})`)
	_ = jsModule.Set("require", r.Load)
	jsExports, _ := jsModule.Get("exports")

	moduleReturn, err := vm.Call(source, jsExports, jsModule)
	if err != nil {
		panicOtto(err.Error())
	}
	var moduleValue otto.Value
	if !moduleReturn.IsUndefined() {
		moduleValue = moduleReturn
		_ = jsModule.Set("exports", moduleValue)
	} else {
		moduleValue, _ = jsModule.Get("exports")
	}
	r.addToCache(path, moduleValue)
	return moduleValue
}
