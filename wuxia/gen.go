package wuxia

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gocraft/health"
	"github.com/robertkrimen/otto"
	"github.com/spf13/afero"
)

const (
	scriptsDir = "scripts"
	initDir    = "init"
	planDir    = "plan"
	configFile = "config.json"
	indexFile  = "index.js"
)

//BuildStage the step in the generation process life cycle.
type BuildStage int

// available stages offered by the generator
const (
	StageInit BuildStage = iota
	StageConfig
	StagePlan
	StageExec
)

func (s BuildStage) String() string {
	var rst string
	switch s {
	case StageInit:
		rst = "init"
	case StageConfig:
		rst = "config"
	case StagePlan:
		rst = "plan"
	case StageExec:
		rst = "exec"
	default:
		rst = "unkown stage"
	}
	return rst
}

//buildError error returned when building the static website. The error string
//returned is a json string that encodes the build stage and the message.
type buildError struct {
	Stage   string `json:"stage"`
	Message string `json:"msg"`
}

func buildErr(stage BuildStage, msg string) error {
	return &buildError{Stage: stage.String(), Message: msg}
}

func (b *buildError) Error() string {
	o, err := json.Marshal(b)
	if err != nil {
		return err.Error()
	}
	return string(o)
}

//Generator provides the static website generation capabilities.This is heavily
//integrated with the otto javascript runtime.
type Generator struct {
	vm  *otto.Otto
	sys *System
	fs  afero.Fs

	// This is the absolute path to the root of the project from which the
	// Generator will be operating.
	workDir string
	job     *health.Job

	Verbose bool

	// when the verbose option is set to true. Here is where the ;og information
	// will be written to.
	Out io.Writer
}

//NewGenerator retrunes a new  Generator.
func NewGenerator(vm *otto.Otto, sys *System, fs afero.Fs) *Generator {
	return &Generator{
		vm:  vm,
		sys: sys,
		fs:  fs,
	}
}

//Init initializes the build process. Any stages after this will have the generator
//already bootstraped.
//
// It is possible to bootstrap the generator from the project( User's side) by
// providing an entry javascript file in the default path of
// scripts/init/index.js which will be executed and you can overide the default
// entry script which is evaluated internally
//
// Initialzation is offloaded to the javascript runtine of the generator..Any
// error returned is a build error.
func (g *Generator) Init() error {
	if g.Verbose {
		g.job.Event("initializing_generator")
	}
	_ = g.vm.Set("sys", func(call otto.FunctionCall) otto.Value {
		data, err := json.Marshal(g.sys)
		if err != nil {
			panicOtto(err)
		}
		val, err := call.Otto.Call("JSON.parse", nil, string(data))
		if err != nil {
			panicOtto(err)
		}
		return val
	})
	_, err := g.vm.Eval(entryScript())
	if err != nil {
		return buildErr(StageInit, err.Error())
	}

	_ = g.vm.Set("fileTree", fileTree(g.fs, g.workDir))

	// evaluate project provided entry script if provided. We ignore if the file
	// is not provided but any errors arsing from evaluating a provided script is
	// a built error.
	entryFile := filepath.Join(scriptsDir, initDir, indexFile)
	err = g.evaluateFile(entryFile)
	if err != nil {
		if !os.IsNotExist(err) {
			return buildErr(StageInit, err.Error())
		}
	}
	if g.Verbose {
		g.job.EventKv("initializing_generator.complete",
			health.Kvs{
				"project": g.sys.Config.ProjectName,
			})
	}
	return nil
}

//fileTree provides an array of all files found in the root( which is suppose to
//be the working directory.
//
//The file index is built only once and evaluated once too, then it is cached.
//This operates in the otto runtime.
func fileTree(fs afero.Fs, root string) func(otto.FunctionCall) otto.Value {
	var tree []string
	var v otto.Value
	return func(call otto.FunctionCall) otto.Value {
		if tree != nil {
			return v
		}
		ferr := afero.Walk(fs, "/", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				tree = append(tree, path)
			}
			return nil
		})
		if ferr != nil {
			tree = nil
			panicOtto(ferr.Error())
		}
		v, _ = call.Otto.ToValue(tree)
		return v
	}
}

func defaultSystem() *System {
	return &System{
		Boot: &Boot{
			ConfigiFile: configFile,
			PlanFile:    "index.js",
		},
	}
}

func defaultPlan() *Plan {
	return &Plan{}
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

func defaultVM(sys *System) *otto.Otto {
	return otto.New()
}

//Config configures the generator.
func (g *Generator) Config() error {
	if g.Verbose {
		g.job.Event("configuring_generator")
	}
	if g.sys == nil {
		g.sys = defaultSystem()
	}
	if g.vm == nil {
		g.vm = defaultVM(g.sys)
	}
	// Properly set working directory/
	if g.workDir == "" {
		wd, err := os.Getwd()
		if err != nil {
			return buildErr(StageInit, err.Error())
		}
		g.workDir = wd
	}

	// ensure everything is relative to the working directory
	g.fs = afero.NewBasePathFs(g.fs, g.workDir)
	af := afero.Afero{Fs: g.fs}
	data, err := af.ReadFile(configFile)
	if err != nil {
		return buildErr(StageInit, err.Error())
	}
	cfg := &Config{}
	err = json.Unmarshal(data, cfg)
	if err != nil {
		return buildErr(StageInit, err.Error())
	}
	g.sys.Config = cfg

	// Add reuire
	req := newRequire(g.fs, scriptsDir)
	err = g.registerBuiltin(req)
	if err != nil {
		return buildErr(StageInit, err.Error())
	}
	_ = g.vm.Set("require", req.load)
	if g.Verbose {
		g.job.EventKv("configuring__generator.complete",
			health.Kvs{
				"project": g.sys.Config.ProjectName,
			})
	}
	return nil
}

func (g *Generator) Plan() error {
	pFile := filepath.Join(scriptsDir, planDir, indexFile)
	err := g.evaluateFile(pFile)
	if err != nil {
		if !os.IsNotExist(err) {
			return buildErr(StagePlan, err.Error())
		}
	}
	v, err := g.vm.Call("getCurrentSys", nil)
	if err != nil {
		return buildErr(StagePlan, err.Error())
	}
	str, err := v.ToString()
	if err != nil {
		return buildErr(StagePlan, err.Error())
	}
	sys := &System{}
	err = json.Unmarshal([]byte(str), sys)
	if err != nil {
		return buildErr(StagePlan, err.Error())
	}
	g.sys = sys
	if g.sys.Plan == nil {
		g.sys.Plan = defaultPlan()
	}
	return nil
}

//Exec executes the plans that are outilend by the Plan method. The Plan method
//should be called before this.
func (g *Generator) Exec() error {
	o, err := g.vm.Call("fileTree", nil)
	if err != nil {
		return buildErr(StageExec, err.Error())
	}
	ov, err := o.Export()
	if err != nil {
		return buildErr(StageExec, err.Error())
	}
	if ov == nil {
		// No files to mess with
	}
	files, ok := ov.([]string)
	if !ok {
		// Some fish
	}
	var list FileList
	p := g.sys.Plan
	for i := range files {
		s, err := p.FindStrategy(files[i])
		if err != nil {
			continue
		}
		f, err := g.execStrategy(files[i], s)
		if err != nil {
			return buildErr(StageExec, err.Error())
		}
		list = append(list, f)
	}
	return g.execPlan(list, p)
}

func (g *Generator) execStrategy(filePath string, s *Strategy) (*File, error) {
	return nil, nil
}

func (g *Generator) execPlan(a FileList, p *Plan) error {
	return nil
}

func (g *Generator) down() error {
	return nil
}

func (g *Generator) registerBuiltin(r *require) error {
	f := &fileSys{}
	f.Fs = g.fs
	v := f.export().ToValue(g.vm)
	r.addToCache("fs", v)
	r.addToCache("markdown", markdown().ToValue(g.vm))
	return nil
}
