package views

import (
	"html/template"
	"io"
)

//View is an interface for rendering of templates.
type View interface {
	Render(tpl string, data interface{}, out io.Writer) error
}

type AssetVew struct {
	tpl *template.Template
}

func New(name string, funcs ...template.FuncMap) (View, error) {
	v := &AssetVew{tpl: template.New(name)}
	if len(funcs) > 0 {
		v = &AssetVew{tpl: template.New(name).Funcs(funcs[0])}
	}
	err := v.Load()
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (v *AssetVew) Load() error {
	for _, n := range AssetNames() {
		t := v.tpl.New(n)
		data, err := Asset(n)
		if err != nil {
			return err
		}
		_, err = t.Parse(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}
func (v *AssetVew) Render(tpl string, data interface{}, out io.Writer) error {
	return v.tpl.ExecuteTemplate(out, tpl, data)
}
