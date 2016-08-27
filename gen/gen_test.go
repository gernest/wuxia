package gen

import (
	"os"
	"testing"
)

func TestGenerator_Build(t *testing.T) {
	p := "fixture/site"
	ctx := &Context{
		WorkDir: p,
		Verbose: true,
	}
	err := Configure(ctx)
	if err != nil {
		t.Error(err)
	}

	err = Initilize(ctx)
	if err != nil {
		t.Error(err)
	}
	err = PlanExecution(ctx)
	if err != nil {
		t.Error(err)
	}

	err = Execute(ctx)
	if err != nil {
		t.Error(err)
	}
}

func TestConfigure(t *testing.T) {
	// default work directory
	ctx := &Context{}
	wd, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}

	_ = Configure(ctx)
	if err != nil {
		t.Error(err)
	}
	if ctx.WorkDir != wd {
		t.Error("expected %s got %s", wd, ctx.WorkDir)
	}
}
