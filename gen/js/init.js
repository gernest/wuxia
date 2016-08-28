var System=sys();
var Tpl={};
var fs=require('fs');
Tpl.funcs={};
Tpl.getTplFuncs=function(){
	var rst=[];
	for (var prop in Tpl.funcs){
		if (Tpl.funcs.hasOwnProperty(prop)){
			rst.push(prop);
		}
	}
	return rst;
};

function getCurrentSys(){
	return JSON.stringify(System);
}

function process(fileName){
  var file ={
    name: fileName,
    contents: "",
    meta:{},
  };
  try{
    var f=fs.open(fileName);
    file.contents=f.read();
    f.close();
  }catch(e){
    throw e;
  }
  return file;
}

function prepare(plan){
    if (plan){
        _.each(plan.Dependency, function(el){
                require(el);
        });
        return true;
    }
    return false;
}


