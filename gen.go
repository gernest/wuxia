package wuxia

import (
	"encoding/json"
	"os"

	"github.com/robertkrimen/otto"
	"github.com/spf13/afero"
)

const (
	scriptsDir = "_scripts"
	initDri    = "init"
	planDir    = "plan"
	configFile = "config.json"
)

//buildError error returned when building the static website. The error string
//returned is a json string that encodes the build stage and the message.
type buildError struct {
	Stage   string `json:"stage"`
	Message string `json:"msg"`
}

func buildErr(stage, msg string) error {
	return &buildError{Stage: stage, Message: msg}
}

func (b *buildError) Error() string {
	o, err := json.Marshal(b)
	if err != nil {
		return err.Error()
	}
	return string(o)
}

type Generator struct {
	vm  *VM
	sys *System
	fs  afero.Fs
}

func (g *Generator) Build() error {
	return evaluate(g.init, g.config, g.plan, g.exec, g.down)
}
func (g *Generator) init() error {
	if g.sys == nil {
		g.sys = defaultSystem()
	}
	if g.vm == nil {
		g.vm = defaultVM(g.sys)
	}
	g.vm.Set("sys", func(call otto.FunctionCall) otto.Value {
		data, err := json.Marshal(g.sys)
		if err != nil {
			Panic(err)
		}
		val, err := call.Otto.Call("JSON.parse", nil, string(data))
		if err != nil {
			Panic(err)
		}
		return val
	})
	_, err := g.vm.Eval(entryScript())
	if err != nil {
		return buildErr("init", err.Error())
	}
	return nil
}

func defaultSystem() *System {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return &System{
		Boot: &Boot{
			ConfigiFile: configFile,
			PlanFile:    "index.js",
		},
		WorkDir: pwd,
	}
}

func defaultVM(sys *System) *VM {
	return &VM{otto.New()}
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
