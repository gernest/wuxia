package gen

import "github.com/robertkrimen/otto"

// Plan is the execution planner object. It states the steps and stages on which
// the execution process should take.
type Plan struct {
	Title string `json:"title"`

	// Modules that are supposed to be loaded before the execution starts. The
	// execution process wont start if one of the dependencies is missing.
	Dependency []string `json:"dependencies"`
	Before     []string `json:"before"`
	Exec       []string `json:"exec"`
	After      []string `json:"after"`
}

//GetPlan retrieves Plan from a javascript src.
func GetPlan(vm *otto.Otto, src string) (*Plan, error) {
	p := &Plan{}
	err := getToJSON(p, vm, src)
	if err != nil {
		return nil, err
	}
	return p, nil
}
