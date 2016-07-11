package gen

func entryScript() string {
	return `
var system=sys();
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
}
