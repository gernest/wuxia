package gen

import (
	"os"
	"testing"

	"github.com/spf13/afero"
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

func TestPlanExecution(t *testing.T) {
	// check whether the default plan is correctly set
	//TODO: find a better way, as the refenece error that comes from otto is a
	//red flag
	nctx := &Context{
		WorkDir: "fixture/site",
		Verbose: true,
		FS:      afero.NewMemMapFs(),
	}
	_ = Configure(nctx)
	_ = Initilize(nctx)
	_ = PlanExecution(nctx)
	np := nctx.Sys.Plan
	if np == nil {
		t.Fatal("expected default plan")
	}
	if np.Title != "default_plan" {
		t.Errorf("expected default_plan got 5s", np.Title)
	}

	// Testcase for missing modules
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
	ctx.Sys.Boot.PlanFile = "nothing.js"
	err = PlanExecution(ctx)
	if err == nil {
		t.Error("expected an error")
	}
}
