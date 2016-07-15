package wuxia

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
)

func TestGenerator_Build(t *testing.T) {
	basePth := "fixture/site"
	pwd, err := os.Getwd()
	fs := afero.NewBasePathFs(afero.NewOsFs(), filepath.Join(pwd, basePth))
	g := NewGenerator(nil, nil, fs)
	err = g.Build()
	if err != nil {
		t.Error(err)
	}
}
