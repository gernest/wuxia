package wuxia

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/robertkrimen/otto"
	"github.com/spf13/afero"
)

const (
	scriptsDir = "_scripts"
	initDir    = "init"
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

//initializes the build process. Any stages after this will have the generator
//already bootstraped.
//
// It is possible to bootstrap the generator from the project( User's side) by
// providing an entry javascript file in the default path of
// scripts/init/index.js which will be executed and you can overide the default
// entry script which is evaluated internally
//
// Initialzation is offloaded to the javascript runtine of the generator..Any
// error returned is a build error.
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

	// evaluate project provided entry script if provided. We ignore if the file
	// is not provided but any errors arsing from evaluating a provided script is
	// a built error.
	entryFile := fmt.Sprintf("%s/%s/index.js", scriptsDir, initDir)
	err = g.evaluateFile(entryFile)
	if err != nil {
		if !os.IsNotExist(err) {
			return buildErr("init", err.Error())
		}
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

// opens the file in the specified path and evaluates it withing the context of
// the javascript runtine.
//
// The evaluated javascript code can mutate the global state. Use execFile to
// execute the javascript without mutating the state of the generato'r
// javascript runtime.
//
// TODO: (gernest) implement callFile if necessary
func (g *Generator) evaluateFile(path string) error {
	f, err := g.fs.Open(path)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	_, err = g.vm.Eval(data)
	if err != nil {
		return err
	}
	return nil
}

func defaultVM(sys *System) *VM {
	return &VM{otto.New()}
}

func (g *Generator) config() error {
	v, err := g.vm.Call("getCurrentSys", nil)
	if err != nil {
		return buildErr("config", err.Error())
	}
	cSys := &System{}
	str, _ := v.ToString()
	err = json.Unmarshal([]byte(str), cSys)
	if err != nil {
		return buildErr("config", err.Error())
	}
	cfgFile := cSys.Boot.ConfigiFile
	cf, err := g.fs.Open(cfgFile)
	if err != nil {
		return buildErr("config", err.Error())
	}
	defer func() { _ = cf.Close() }()
	data, err := ioutil.ReadAll(cf)
	if err != nil {
		return buildErr("config", err.Error())
	}
	c := &Config{}
	err = json.Unmarshal(data, c)
	if err != nil {
		return buildErr("config", err.Error())
	}
	if cSys.Config == nil {
		cSys.Config = c
	} else {
		updateConfig(cSys.Config, c)
	}
	g.sys = cSys
	return nil
}

func updateConfig(old, new *Config) {
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
	for i := 0; i < len(fn); i++ {
		err = fn[i]()
		if err != nil {
			return err
		}
	}
	return err
}
