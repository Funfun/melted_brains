function set_cursor(){
  // skip comment & new lines
  var elements = arguments[0].getElementsByTagName('span');
  var collected = [];
  for(var i=0; i<elements.length; i++){
    if(["com", "pln"].indexOf(elements[i].className)==-1 ){
      collected.push(elements[i]);
    }
  }
  // TODO: set cursor on first element
}
function onLoad(){
  var codeBlock = document.getElementById('code');
  set_cursor(codeBlock);
}
window.onload = onLoad;