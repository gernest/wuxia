package views

import "io"

//View is an interface for rendering of templates.
type View interface {
	Render(tpl string, data interface{}, out io.Writer) error
}
