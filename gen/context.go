package gen

import (
	"io"

	"github.com/gocraft/health"
	"github.com/robertkrimen/otto"
	"github.com/spf13/afero"
)

//Context is the context of a static website project.
type Context struct {
	FS      afero.Fs
	VM      *otto.Otto
	Sys     *System
	WorkDir string
	Verbose bool
	Job     *health.Job
	Out     io.Writer
}

//Stage is an interface representing a generation step
type Stage interface {
	Name() string
	Exec(*Context) error
}

type baseStage struct {
	name     string
	execFunc func(*Context) error
}

//NewStage returns struct that satisfies the Stage interface. The deails are
//hidden deilibarately as a way to quickly define stages based on functions of
//struct methods.
//
// It is recommended to implement Stage interface explicity. Use this only when
// it is necessary.
func NewStage(name string, e func(*Context) error) Stage {
	return &baseStage{name: name, execFunc: e}
}

func (b *baseStage) Name() string {
	return b.name
}
func (b *baseStage) Exec(ctx *Context) error {
	return b.execFunc(ctx)
}
