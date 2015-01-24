function parseTokens(){
  // skip comment & new lines
  var elements = arguments[0].getElementsByTagName('span');
  var collected = [];
  for(var i=0; i<elements.length; i++){
    if(["com"].indexOf(elements[i].className)==-1 && elements[i].innerHTML.replace(/\n|\t/g, '') != ''){
      collected.push(elements[i]);
    }
  }
  return collected;
}

function MacOsXCharCode(charCode){
  if(charCode == 13){
    return 10;
  }
  return charCode;
}

function onKeyPress(pEvent){
  var charCode = pEvent.charCode,
      newOffset = null;

  console.log(window.tokenElem.value);
  console.log(charCode);
  if(window.tokenElem.value.charCodeAt(window.tokenElem.offSet) == MacOsXCharCode(charCode)){
    console.log('correct', charCode, window.tokenElem.value, window.tokenElem.offSet)
    newOffset = window.tokenElem.offSet++;
    window.tokenElem.el.innerHTML = window.tokenElem.value.substr(0, newOffset+1) + "|" + window.tokenElem.value.substr(newOffset+1);
    console.log(newOffset, window.tokenElem.value.length);
    if( (newOffset+1) >= window.tokenElem.value.length){
      console.log("Next");

      window.tokenElem = chooseNextAt(window.tokenElem.index+1);
    }
  }
}
function chooseNextAt(idx){
  console.log("idx", idx);
  return {
    index: idx,
    el: window.tokens[idx],
    value: window.tokens[idx].innerHTML,
    offSet: 0
  }
}
function onLoad(){
  var codeBlock = document.getElementById('code'),
      tokens = parseTokens(codeBlock);

  // we need global
  window.tokens = tokens;
  window.tokenElem = chooseNextAt(0);
  document.onkeypress = onKeyPress;

}
window.onload = onLoad;