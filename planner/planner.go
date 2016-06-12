package planner

import (
	"encoding/json"
	"errors"

	"github.com/robertkrimen/otto"
)

// Plan is the execution planner object. It states the steps and stages on which
// the execution process should take.
type Plan struct {
	Title string `json:"tito;e"`

	// Modules that are supposed to be loaded before the execution starts. The
	// execution process wont start if one of the dependencies is missing.
	Dependency []string `json:"dependencies"`
	Before     []string `json:"before"`
	Exec       []string `json:"exec"`
	After      []string `json:"after"`
}

//GetPlan retrieves Plan from a javascript src.
func GetPlan(vm *otto.Otto, src string) (*Plan, error) {
	source := "JSON.stringify( function(){" + src + "}())"
	v, err := vm.Run(source)
	if err != nil {
		return nil, err
	}
	if v.IsUndefined() || v.IsNull() {
		return nil, errors.New("bad planner script")
	}
	s, err := v.ToString()
	if err != nil {
		return nil, err
	}
	p := &Plan{}
	err = json.Unmarshal([]byte(s), p)
	if err != nil {
		return nil, err
	}
	return p, nil
}
