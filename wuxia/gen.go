package wuxia

import (
	"encoding/json"
	"errors"
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

//Initilize initializes the build process. Any stages after this will have the generator
//already bootstraped.
//
// It is possible to bootstrap the generator from the project( User's side) by
// providing an entry javascript file in the default path of
// scripts/init/index.js which will be executed and you can overide the default
// entry script which is evaluated internally
//
// Initialzation is offloaded to the javascript runtine of the generator..Any
// error returned is a build error.
func Initilize(ctx *Context) error {
	_ = ctx.VM.Set("sys", func(call otto.FunctionCall) otto.Value {
		data, err := json.Marshal(ctx.Sys)
		if err != nil {
			panicOtto(err)
		}
		val, err := call.Otto.Call("JSON.parse", nil, string(data))
		if err != nil {
			panicOtto(err)
		}
		return val
	})
	_, err := ctx.VM.Eval(entryScript())
	if err != nil {
		return err
	}

	_ = ctx.VM.Set("fileTree", fileTree(ctx.FS, ctx.WorkDir))

	// evaluate project provided entry script if provided. We ignore if the file
	// is not provided but any errors arsing from evaluating a provided script is
	// a built error.
	entryFile := filepath.Join(scriptsDir, initDir, indexFile)
	err = EvaluateFile(ctx.FS, ctx.VM, entryFile)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
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

//EvaluateFile  opens the file in the specified path and evaluates it withing the context of
// the javascript runtine.
//
// The evaluated javascript code can mutate the global state. Use execFile to
// execute the javascript without mutating the state of the generato'r
// javascript runtime.
func EvaluateFile(fs afero.Fs, vm *otto.Otto, path string) error {
	f, err := fs.Open(path)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	_, err = vm.Eval(data)
	if err != nil {
		return err
	}
	return nil
}

func defaultVM(sys *System) *otto.Otto {
	return otto.New()
}

func Configure(ctx *Context) error {
	if ctx.WorkDir == "" {
		wd, err := os.Getwd()
		if err != nil {
			return err
		}
		ctx.WorkDir = wd
	}
	if ctx.Out == nil {
		ctx.Out = os.Stdout
	}
	if ctx.Job == nil {
		if stream.Sinks == nil {
			stream.AddSink(&health.WriterSink{Writer: ctx.Out})
		}
		ctx.Job = stream.NewJob("generate:" + ctx.WorkDir)
	}
	if ctx.FS == nil {
		ctx.FS = afero.NewOsFs()
	}

	// ensure everything is relative to the working directory
	// The working directory is where the  directory in which the project to be
	// built lives.
	//
	// All file operations, happens within the diretory. So to access the
	// configuration file for example it is located in /config.json.
	ctx.FS = afero.NewBasePathFs(ctx.FS, ctx.WorkDir)
	af := afero.Afero{Fs: ctx.FS}
	data, err := af.ReadFile(configFile)
	if err != nil {
		return err
	}
	cfg := &Config{}
	err = json.Unmarshal(data, cfg)
	if err != nil {
		return err
	}
	if ctx.Sys == nil {
		ctx.Sys = defaultSystem()
	}
	ctx.Sys.Config = cfg

	if ctx.VM == nil {
		ctx.VM = defaultVM(ctx.Sys)
	}
	// Add reuire
	req := newRequire(ctx.FS, scriptsDir)
	err = RegisterBuiltins(ctx, req)
	if err != nil {
		return err
	}
	_ = ctx.VM.Set("require", req.load)
	return nil
}

//PlanExecution is the planning phase of the generator. First a project specific plan
//file is evaluated, the file is loacetd in /scripts_dir/plan_fir/index.js for
//example /scripts/plan/index.js. Second the existing system plan is merged with
//default plan.
//
//TODO: Use user defined plan only when it is set, and use default plan only
//when there is no user defined plan.
//
// This is executed to prepare the Plan object, which is the blueprint on how
// the whole execution is going to take place.
func PlanExecution(ctx *Context) error {
	pFile := filepath.Join(scriptsDir, planDir, indexFile)
	err := EvaluateFile(ctx.FS, ctx.VM, pFile)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}
	v, err := ctx.VM.Call("getCurrentSys", nil)
	if err != nil {
		return err
	}
	str, err := v.ToString()
	if err != nil {
		return err
	}
	sys := &System{}
	err = json.Unmarshal([]byte(str), sys)
	if err != nil {
		return err
	}
	ctx.Sys = sys
	if ctx.Sys.Plan == nil {
		ctx.Sys.Plan = defaultPlan()
	}
	return nil
}

//Execute  executes the plans that are outilend by the plans found in ctx.
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
func Execute(ctx *Context) error {
	o, err := ctx.VM.Call("fileTree", nil)
	if err != nil {
		return err
	}
	ov, err := o.Export()
	if err != nil {
		return err
	}
	if ov == nil {
		// No files to mess with
		return errors.New("no files to build")
	}
	files, ok := ov.([]string)
	if !ok {
		// Some fish
		return err
	}

	var wg sync.WaitGroup
	var list FileList
	var errlist []error
	p := ctx.Sys.Plan
	for i := range files {
		wg.Add(1)
		go func(fn string) {
			defer wg.Done()
			s, err := p.FindStrategy(fn)
			if err != nil {
				errlist = append(errlist, err)
				return
			}
			f, err := ExecStrategy(ctx, fn, s)
			if err != nil {
				errlist = append(errlist, err)
				return
			}
			list = append(list, f)
		}(files[i])

	}
	wg.Wait()
	return ExecPlaan(ctx, list, p)
}

func ExecStrategy(ctx *Context, filename string, s *Strategy) (*File, error) {
	return nil, nil
}

func ExecPlaan(ctx *Context, fl FileList, s *Plan) error {
	return nil
}

//RegisterBuiltins adds important modules  like underscore and fs.
func RegisterBuiltins(ctx *Context, r *require) error {
	f := &fileSys{}
	f.Fs = ctx.FS
	v := f.export().ToValue(ctx.VM)
	r.addToCache("fs", v)
	r.addToCache("markdown", markdown().ToValue(ctx.VM))
	v, err := ctx.VM.Run(underscore.Source())
	if err != nil {
		return err
	}
	r.addToCache("underscore", v)
	return nil
}
