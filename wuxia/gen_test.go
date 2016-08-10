package wuxia

import (
	"testing"

	"github.com/spf13/afero"
)

func TestGenerator_Build(t *testing.T) {
	g := NewGenerator(nil, nil, afero.NewOsFs())
	p := "fixture/site"
	g.workDir = p
	g.Verbose = true
	err := g.Config()
	if err != nil {
		t.Error(err)
	}

	err = g.Init()
	if err != nil {
		t.Error(err)
	}
	err = g.Plan()
	if err != nil {
		t.Error(err)
	}
}
