package gen

import "testing"

func TestGenerator_init(t *testing.T) {
	g := &Generator{}
	err := g.init()
	if err != nil {
		t.Error(err)
	}
	_, err = g.vm.Eval(`
	if (system.boot.config_file!="config.json"){
		throw Error("failed to set init system object");
	}
`)
	if err != nil {
		t.Error(err)
	}
}
