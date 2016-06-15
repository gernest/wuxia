package gen

import (
	"github.com/gernest/valeria/gen"
	"github.com/gernest/valeria/vm"
	"github.com/spf13/afero"
)

type Generator struct {
	vm     *vm.VM
	config *gen.Config
	fs     afero.Fs
}
