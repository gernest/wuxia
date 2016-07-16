// Thests for the open method for fs module. This opens a file for reading only.
// 
var fs=newFS();
var fileName="hello.txt";
var content="hello";
function testOpen(){
  console.log("-- Testing fs.open");
  try{
    f=fs.open(fileName);
    message=f.read();
    if(message!=content){
      throw "expected "+content+" got "+message;
    }
    f.close();
  }catch(e){
    throw e;
  }
}
testOpen();
