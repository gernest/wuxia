package gen

import (
	"bytes"
	"testing"

	"github.com/robertkrimen/otto"
)

func TestTemplate(t *testing.T) {
	vm := otto.New()
	tplFunc := `
var Tpl={};
Tpl.funcs={};
Tpl.funcs.world=function(name){
	return name+",world"
}
Tpl.getTplFuncs=function(){
	var rst=[]
	for (var prop in Tpl.funcs){
		if (Tpl.funcs.hasOwnProperty(prop)){
			rst.push(prop)
		}
	}
	return rst
}
`
	_, err := vm.Eval(tplFunc)
	if err != nil {
		t.Error(err)
	}
	tpl := &Template{vm: vm}
	tpl = tpl.New()
	sample := `{{"hello"|world}}`
	tp, err := tpl.Parse(sample)
	if err != nil {
		t.Error(err)
	}
	buf := &bytes.Buffer{}
	err = tp.Execute(buf, nil)
	if err != nil {
		t.Error(err)
	}
	if buf.String() != "hello,world" {
		t.Errorf("expected hello,world got %s", buf)
	}
}
