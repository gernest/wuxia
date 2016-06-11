package vm

import (
	"io/ioutil"
	"testing"

	"github.com/robertkrimen/otto"
)

func TestFile_Load(t *testing.T) {
	req := NewRequre(&File{})
	req.SetWorkingDir("fixture")
	vm := otto.New()
	vm.Set("require", req.ToValue())
	data, err := ioutil.ReadFile("fixture/index.js")
	if err != nil {
		t.Error(err)
	}
	result, err := vm.Run(data)
	if err != nil {
		t.Error(err)
	}
	o, _ := result.ToString()
	if o != "hello" {
		t.Errorf("expected hello got %s", o)
	}
}
