package gen

import (
	"os"

	"github.com/spf13/afero"
)

func sampleProjectFs() afero.Fs {
	data := []struct {
		path     string
		mode     os.FileMode
		contents []byte
	}{
		{"README.md", 0600, []byte(readmeFile)},
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
Project v is a moderm static content generator for building anything static. It is fast, modular and flexible, combining the power of Go and the easy of use of Javascript.
`
