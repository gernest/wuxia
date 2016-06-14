package gen

import (
	"testing"

	"github.com/robertkrimen/otto"
)

func TestGetPlan(t *testing.T) {
	var src = `
var o={};
o.title="simple";
o.dependencies=["base","simple"];
o.before=["setup"];
o.exec=["build"];
o.after=["cleanup"];
return o;
`
	vm := otto.New()
	p, err := GetPlan(vm, src)
	if err != nil {
		t.Fatal(err)
	}
	if p.Title != "simple" {
		t.Errorf("expected title got %s", p.Title)
	}
}
