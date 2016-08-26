package gen

import "testing"

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
