package gen

import (
	"errors"
	"fmt"
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

func (r *require) load(call *otto.FunctionCall) otto.Value {
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
	return otto.Value{}
}
