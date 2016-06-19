package gen

import (
	"encoding/json"

	"github.com/gernest/valeria/gen"
	"github.com/gernest/valeria/vm"
	"github.com/spf13/afero"
)

const (
	scriptsDir = "_scripts"
	initDri    = "init"
	planDir    = "plan"
	configFile = "config.json"
)

type BuildError struct {
	Stage   string `json:"stage"`
	Message string `json:"msg"`
}

func (b BuildError) Error() string {
	o, err := json.Marshal(b)
	if err != nil {
		return err.Error()
	}
	return string(o)
}

type Generator struct {
	vm  *vm.VM
	sys *gen.System
	fs  afero.Fs
}

func (g *Generator) Build() error {
	return evaluate(g.init, g.config, g.plan, g.exec, g.down)
}
func (g *Generator) init() error {
	return nil
}
func (g *Generator) config() error {
	return nil
}
func (g *Generator) plan() error {
	return nil
}
func (g *Generator) exec() error {
	return nil
}
func (g *Generator) down() error {
	return nil
}

func evaluate(fn ...func() error) error {
	var err error
	for _, f := range fn {
		err = f()
		if err != nil {
			return err
		}
	}
	return err
}
