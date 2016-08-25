package wuxia

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

type Stage interface {
	Name() string
	Exec(*Context) error
}

type baseStage struct {
	name     string
	execFunc func(*Context) error
}

func (b *baseStage) Name() string {
	return b.name
}
func (b *baseStage) Exec(ctx *Context) error {
	return b.execFunc(ctx)
}
