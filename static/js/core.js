function parseTokens(){
  // skip comment & new lines
  var elements = arguments[0].getElementsByTagName('span');
  var collected = [];
  for(var i=0; i<elements.length; i++){
    if(["com"].indexOf(elements[i].className)==-1 && elements[i].innerHTML.replace(/\n|\t/g, '') != ''){
      elements[i].classList.add('arc');
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
function numberOFDupInStr(str){
  var res = str.match(/(\t{1,})/g);

  if(str[0] == "\t" && res){
    console.log("Tab first", str[0] == "\t");
    console.log("Tab size", res[0].length);
    return res[0].length;
  }
  else {
    return 0;
  }

}

function onKeyPress(pEvent){
  var charCode = pEvent.charCode,
      newOffset = null,
      isNextTab,
      i;
  if(MacOsXCharCode(charCode) == window.tokenElem.value.charCodeAt(window.tokenElem.offSet)){
    isNextTabs = numberOFDupInStr(window.tokenElem.value.substr(window.tokenElem.offSet+1));
    newOffset = window.tokenElem.offSet++;

    if(isNextTabs > 0){
      window.tokenElem.offSet = window.tokenElem.offSet + isNextTabs;
      newOffset = window.tokenElem.offSet - 1;
    }

    window.tokenElem.el.innerHTML = setCarretAndSkipTab(newOffset);

    // next span element
    if( (newOffset+1) >= window.tokenElem.value.length || window.tokenElem.value == "&amp;"){
      window.tokenElem.el.innerHTML = window.tokenElem.value;
      window.tokenElem.el.classList.remove('arc');
      window.tokenElem = chooseNextAt(window.tokenElem.index+1);
      var str = window.tokenElem.value;
      window.tokenElem.el.innerHTML = setCarret(str);
    }
  }
}
function setCarretAndSkipTab(offSet){
  // case01
  if(window.tokenElem.value == "&amp;"){
    return (window.tokenElem.value + "|");
  }
  return window.tokenElem.value.substr(0, offSet+1) + "|" + window.tokenElem.value.substr(offSet+1);
}
function setCarret(str){
  return "|" + str;
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