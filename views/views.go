package views

//View is an interface for rendering of templates.
type View interface {
	Render(tpl string, data interface{}) error
}
