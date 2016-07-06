package gen

import (
	"fmt"
	"testing"

	"github.com/robertkrimen/otto"
	"github.com/spf13/afero"
)

func TestRequire_resolve(t *testing.T) {
	fs := afero.NewMemMapFs()
	sample := []struct {
		path, script string
	}{
		{"/project/modules/echo.js", `
function echo(msg){
	return msg;
}
	`},
		{"/project/index.js", `
		var echo= require("echo.js");
		echo("index.js")
`},
	}
	for _, v := range sample {
		f, err := fs.Create(v.path)
		if err != nil {
			t.Fatal(err)
		}
		_, err = f.WriteString(v.script)
		if err != nil {
			t.Fatal(err)
		}
		f.Close()
	}
	req := &require{
		fs: fs,
		paths: []string{
			"/project",
			"/project/modules",
		},
	}
	expectation := []struct {
		path, expect string
	}{
		{"index", "/project/index.js"},
		{"index.js", "/project/index.js"},
	}
	for _, e := range expectation {
		result, err := req.resolve(e.path)
		if err != nil {
			t.Fatal(err)
		}

		if result != e.expect {
			t.Errorf("expected %s got %s", e.expect, result)
		}
	}
	// LoadModules
	vm := otto.New()
	vm.Set("require", req.load)
	v, err := vm.Run(sample[1].script)
	if err != nil {
		t.Fatal(err)
	}
	val, err := v.Export()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(val)
}
