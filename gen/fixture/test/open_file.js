
var fs=newFS();
var fileName="hello_open.txt";
var content="hello";
function testOpen(){
  console.log("-- Testing fs.openFile");
  try{
    f=fs.openFile(fileName,"w+c",0600);
    f.write(content);
    f.close();
  }catch(e){
    throw e;
  }
}
testOpen();
