package wuxia

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/gocraft/health"
	"github.com/robertkrimen/otto"
	"github.com/robertkrimen/otto/underscore"
	"github.com/spf13/afero"
	// load underscore for otto runtime.
)

var stream *health.Stream

func init() {
	stream = health.NewStream()
}

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

//g.buildError error returned when building the static website. The error string
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
	g.log(StageInit)
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
		return g.buildErr(StageInit, err)
	}

	_ = g.vm.Set("fileTree", fileTree(g.fs, g.workDir))

	// evaluate project provided entry script if provided. We ignore if the file
	// is not provided but any errors arsing from evaluating a provided script is
	// a built error.
	entryFile := filepath.Join(scriptsDir, initDir, indexFile)
	err = g.evaluateFile(entryFile)
	if err != nil {
		if !os.IsNotExist(err) {
			return g.buildErr(StageInit, err)
		}
	}
	g.log(StageInit, "complete")
	return nil
}

func (g *Generator) log(s BuildStage, args ...interface{}) {
	if g.Verbose {
		g.job.Event(s.String() + ":" + fmt.Sprint(args...))
	}
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
			panicOtto(ferr)
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
	// Properly set working directory/
	if g.workDir == "" {
		wd, err := os.Getwd()
		if err != nil {
			return g.buildErr(StageConfig, err)
		}
		g.workDir = wd
	}

	if g.Verbose {
		if g.Out == nil {
			g.Out = os.Stdout
		}
		if stream.Sinks == nil {
			stream.AddSink(&health.WriterSink{Writer: g.Out})
		}
		g.job = stream.NewJob("generate:" + g.workDir)
	}
	g.log(StageConfig)
	if g.sys == nil {
		g.sys = defaultSystem()
	}
	if g.vm == nil {
		g.vm = defaultVM(g.sys)
	}

	// ensure everything is relative to the working directory
	// The working directory is where the  directory in which the project to be
	// built lives.
	//
	// All file operations, happens within the diretory. So to access the
	// configuration file for example it is located in /config.json.
	g.fs = afero.NewBasePathFs(g.fs, g.workDir)
	af := afero.Afero{Fs: g.fs}
	data, err := af.ReadFile(configFile)
	if err != nil {
		return g.buildErr(StageInit, err)
	}
	cfg := &Config{}
	err = json.Unmarshal(data, cfg)
	if err != nil {
		return g.buildErr(StageInit, err)
	}
	g.sys.Config = cfg

	// Add reuire
	req := newRequire(g.fs, scriptsDir)
	err = g.registerBuiltin(req)
	if err != nil {
		return g.buildErr(StageInit, err)
	}
	_ = g.vm.Set("require", req.load)
	g.log(StageConfig, "complete")
	return nil
}

//Plan is the planning phase of the generator. First a project specific plan
//file is evaluated, the file is loacetd in /scripts_dir/plan_fir/index.js for
//example /scripts/plan/index.js. Second the existing system plan is merged with
//default plan.
//
//TODO: Use user defined plan only when it is set, and use default plan only
//when there is no user defined plan.
//
// This is executed to prepare the Plan object, which is the blueprint on how
// the whole execution is going to take place.
func (g *Generator) Plan() error {
	g.log(StagePlan)

	pFile := filepath.Join(scriptsDir, planDir, indexFile)
	err := g.evaluateFile(pFile)
	if err != nil {
		if !os.IsNotExist(err) {
			return g.buildErr(StagePlan, err)
		}
	}
	v, err := g.vm.Call("getCurrentSys", nil)
	if err != nil {
		return g.buildErr(StagePlan, err)
	}
	str, err := v.ToString()
	if err != nil {
		return g.buildErr(StagePlan, err)
	}
	sys := &System{}
	err = json.Unmarshal([]byte(str), sys)
	if err != nil {
		return g.buildErr(StagePlan, err)
	}
	g.sys = sys
	if g.sys.Plan == nil {
		g.sys.Plan = defaultPlan()
	}
	g.log(StagePlan, "complete")
	return nil
}

//Exec executes the plans that are outilend by the Plan method. The Plan method
//should be called before this.
//
// Execution is done in two phases. First the files are evaluated using the
// matching plans. Plans are per file. Secondly the rendering of the site is
// done.
//
// Evaluation of files is done concurrently, then the results are aggregated
// back together.
//
// Rendering is done synchronously. As there is a lot of cross refences and
// stuffs to be considered, the rendering function is provided by the plan.
func (g *Generator) Exec() error {
	g.log(StageExec)
	o, err := g.vm.Call("fileTree", nil)
	if err != nil {
		return g.buildErr(StageExec, err)
	}
	ov, err := o.Export()
	if err != nil {
		return g.buildErr(StageExec, err)
	}
	if ov == nil {
		// No files to mess with
	}
	files, ok := ov.([]string)
	if !ok {
		// Some fish
		return g.buildErr(StageExec, errors.New("no files to build"))
	}

	var wg sync.WaitGroup
	var list FileList
	var errlist []error
	p := g.sys.Plan
	for i := range files {
		wg.Add(1)
		go func(fn string) {
			defer wg.Done()
			if g.Verbose {
			}
			g.log(StageExec, " evaluating "+fn)
			s, err := p.FindStrategy(fn)
			if err != nil {
				errlist = append(errlist, err)
				if g.Verbose {
					g.job.EventErr("exec:evaluating "+fn, err)
				}
				return
			}
			f, err := g.execStrategy(fn, s)
			if err != nil {
				errlist = append(errlist, err)
				if g.Verbose {
					g.job.EventErr("exec:evaluating "+fn, err)
				}
				return
			}
			list = append(list, f)
			if g.Verbose {
				g.job.Event("exec:   complete evaluating " + fn)
			}
		}(files[i])

	}
	wg.Wait()
	return g.execPlan(list, p)
}

func (g *Generator) execStrategy(filePath string, s *Strategy) (*File, error) {
	return nil, nil
}

func (g *Generator) execPlan(a FileList, p *Plan) error {
	if g.Verbose {
		g.job.Event("exec: render")
	}
	if g.Verbose {
		g.job.Event("exec: done")
	}
	return nil
}

func (g *Generator) down() error {
	return nil
}

// adds modules that are shipped with the Generator to require.
func (g *Generator) registerBuiltin(r *require) error {
	f := &fileSys{}
	f.Fs = g.fs
	v := f.export().ToValue(g.vm)
	r.addToCache("fs", v)
	r.addToCache("markdown", markdown().ToValue(g.vm))
	v, err := g.vm.Run(underscore.Source())
	if err != nil {
		return err
	}
	r.addToCache("underscore", v)
	return nil
}

func (g *Generator) buildErr(s BuildStage, err error) error {
	if g.Verbose {
		return g.job.EventErr(s.String(), err)
	}
	return buildErr(s, err.Error())
}
