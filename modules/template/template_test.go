package template

import (
	"testing"

	"github.com/robertkrimen/otto"
)

func TestTemplate(t *testing.T) {
	vm := otto.New()
	vm.Set("Template", func(call otto.FunctionCall) otto.Value {
		name, _ := call.Argument(0).ToString()
		return NewObject(vm, name)
	})
	vm.Set("error", func(call otto.FunctionCall) otto.Value {
		name, _ := call.Argument(0).ToString()
		t.Error(name)
		return otto.UndefinedValue()
	})

	// Template.Name
	var templateTane = `
function TestTemplateName(name){
	var tpl =Template(name);
	if (tpl.name()!=name){
		throw "expected "+name+" got "+tpl.name();
	}
}
TestTemplateName("home");
`
	_, err := vm.Eval(templateTane)
	if err != nil {
		t.Error(err)
	}

	//
	// Template.Parse
	//
	var templateParseBad = `
function TestTemplateParseBad(){
	var tpl =Template("test");
	var err="";
	try{
		tpl.parse("hello {{")
	}catch(e){
		err=e;
	}
	if(err==""){
		error("expected exection got nothing instead");
	}
}
TestTemplateParseBad();
`
	_, err = vm.Eval(templateParseBad)
	if err != nil {
		t.Error(err)
	}

	var templateParseGood = `
function TestTemplateParseGood(){
	var tpl =Template("test");
	var err="";
	try{
		tpl.parse("hello {{.name}}")
	}catch(e){
		err=e;
	}
	if(err!=""){
		error("expected no exception got "+err);
	}
}
TestTemplateParseGood();
`
	_, err = vm.Eval(templateParseGood)
	if err != nil {
		t.Error(err)
	}
}
