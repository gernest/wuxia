package fs

import (
	"bytes"

	"github.com/robertkrimen/otto"
)

type File struct {
	buf *bytes.Buffer
}

func (f *File) Read(call otto.FunctionCall) otto.Value {
	arg1 := call.Argument(0)
	if arg1.IsFunction() {
	}
}
