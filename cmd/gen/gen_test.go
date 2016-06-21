package gen

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/gernest/valeria/gen"
	"github.com/spf13/afero"
)

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

func sampleProjectFs() afero.Fs {
	data := []struct {
		path     string
		mode     os.FileMode
		contents []byte
	}{
		{"README.md", 0600, []byte(readmeFile)},
		{"config.json", 0600, sampleConfigFile()},
	}
	fs := afero.NewMemMapFs()
	for i := 0; i < len(data); i++ {
		file, err := fs.OpenFile(data[i].path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, data[i].mode)
		if err != nil {
			panic(err)
		}
		_, err = file.Write(data[i].contents)
		if err != nil {
			panic(err)
		}
		if err = file.Close(); err != nil {
			panic(err)
		}
	}
	return fs
}

var readmeFile = `
# v
Project v is a moderm static content generator for building anything static. It 
is fast, modular and flexible, combining the power of Go and the easy of use of 
Javascript.
`

func sampleConfigFile() []byte {
	cfg := gen.Config{
		Source:      "src",
		Destination: "dest",
		Safe:        true,
		Excluede: []string{
			"tmp/", "bin",
		},
		KeepFiles: []string{
			"LICENCE",
		},
		Encoding: "utf-8",
		Port:     1900,
		Host:     "localhost",
	}
	data, _ := json.Marshal(cfg)
	return data
}
