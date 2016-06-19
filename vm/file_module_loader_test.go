package vm

import (
	"fmt"
	"testing"

	"github.com/robertkrimen/otto"
	"github.com/spf13/afero"
)

func TestFile_Load(t *testing.T) {
	var echo = `
function echo(msg){
  return msg;
}

exports.echo=echo;
`
	var index = `
var echo=require("./echo.js");
echo.echo("hello");
`
	memFs := afero.NewMemMapFs()
	efile, err := memFs.Create("fixture/echo.js")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = efile.Close() }()
	fmt.Fprint(efile, echo)

	ifile, err := memFs.Create("fixture/index.js")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = ifile.Close() }()
	fmt.Fprint(ifile, index)
	f := &File{}
	f.fs = memFs
	req := NewRequre(f)
	req.SetWorkingDir("fixture")
	vm := otto.New()
	vm.Set("require", req.ToValue())
	result, err := vm.Run(index)
	if err != nil {
		t.Error(err)
	}
	o, _ := result.ToString()
	if o != "hello" {
		t.Errorf("expected hello got %s", o)
	}

	// Loading json
	var j = `{"title": "hello"}`
	jfile, err := memFs.Create("fixture/config.json")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = jfile.Close() }()
	fmt.Fprint(jfile, j)

	var jscript = `
var echo=require("./echo.js");
var cfg=require("./config.json");
echo.echo(cfg.title);
`
	result, err = vm.Run(jscript)
	if err != nil {
		t.Error(err)
	}
	o, _ = result.ToString()
	if o != "hello" {
		t.Errorf("expected hello got %s", o)
	}
}
