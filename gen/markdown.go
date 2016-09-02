package gen

import (
	"github.com/gernest/blackfriday"
	"github.com/robertkrimen/otto"
)

const (
	fileNameKey    = "name"
	fileMetaKey    = "meta"
	fileContentKey = "contents"
)

// markdown plugin for the otto runtime. This exposes one method exec, which
// accepts a File object as an argument. It renders the Contents of the file to
// markdown.
//
// A HTML method is provided just in case you want to render a piece of text
// as markdown.
func markdown() Export {
	e := make(Export)
	e.Set("exec", func(call otto.FunctionCall) otto.Value {
		v, err := call.Argument(0).Export()
		if err != nil {
			panicOtto(err.Error())
		}
		o, ok := v.(map[string]interface{})
		if !ok {
			panicOtto("markdown: wrong ragument type, expected a file object")
		}
		if c, ok := o[fileContentKey]; ok {
			if cv, cok := c.(string); cok {
				if cv != "" {
					o[fileContentKey] = blackfriday.MarkdownCommon([]byte(cv))
				}
			}
		}
		ov, _ := call.Otto.ToValue(o)
		return ov
	})
	e.Set("HTML", func(call otto.FunctionCall) otto.Value {
		v, err := call.Argument(0).ToString()
		if err != nil {
			panicOtto(err.Error())
		}
		v = string(blackfriday.MarkdownCommon([]byte(v)))
		ov, _ := call.Otto.ToValue(v)
		return ov
	})
	return e
}
