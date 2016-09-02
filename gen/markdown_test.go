package gen

import (
	"encoding/json"
	"testing"

	"github.com/gernest/blackfriday"
	"github.com/robertkrimen/otto"
)

func TestMarkdown(t *testing.T) {
	m := Markdown()
	vm := otto.New()
	vm.Set("md", m.ToValue(vm))
	f := &File{
		Name:     "test file",
		Contents: "hello, world",
	}
	b, _ := json.Marshal(f)
	vm.Set("data", string(b))
	expect := string(blackfriday.MarkdownCommon([]byte(f.Contents)))
	vm.Set("expect", expect)
	var mdTest = `
	var file=JSON.parse(data);
	var out =md.exec(file);
	if (out.contents!=expect){
		throw Error("exec: expected "+expect+" got "+out.contents)
	}
	var hello ="hello, world";
	var h=md.HTML(hello);
	if (h!=expect){
		throw Error("render: expected "+expect+" got "+h)
	}
	`
	_, err := vm.Eval(mdTest)
	if err != nil {
		t.Error(err)
	}
}
