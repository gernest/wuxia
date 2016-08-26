package gen

import (
	"github.com/gernest/blackfriday"
	"github.com/robertkrimen/otto"
)

const (
	fileNameKey   = "name"
	fileMetaKey   = "meta"
	fileCOnentKey = "contents"
)

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
		if c, ok := o[fileCOnentKey]; ok {
			if cv, cok := c.(string); cok {
				if cv != "" {
					o[fileCOnentKey] = blackfriday.MarkdownCommon([]byte(cv))
				}
			}
		}
		ov, _ := call.Otto.ToValue(o)
		return ov
	})
	return e
}
