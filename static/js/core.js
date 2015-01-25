window.GO_RACER_NS = {}

window.GO_RACER_NS['parseTokens'] = function (){
  // skip comment & new lines
  var elements = arguments[0].getElementsByTagName('span');
  var collected = [];
  for(var i=0; i<elements.length; i++){
    if(["com"].indexOf(elements[i].className)==-1 && elements[i].innerHTML.replace(/\n|\t/g, '') != ''){
      elements[i].classList.add('arc');
      collected.push(elements[i]);
    }
  }
  debugger;
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
    return res[0].length;
  }
  else {
    return 0;
  }
}

window.GO_RACER_NS['carretMoveLogic'] = function(charCode){
  var newOffset = null,
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
      window.tokenElem = window.GO_RACER_NS.chooseNextAt(window.tokenElem.index+1);
      var str = window.tokenElem.value;
      window.tokenElem.el.innerHTML = setCarret(str);
    }
  }
}

window.GO_RACER_NS['onKeyPress'] = function (pEvent){
  window.ws.send(pEvent.charCode);

  if(pEvent.keyCode == 32 && pEvent.target == document.body) {
    pEvent.preventDefault();
    return false;
  }
}
function setCarretAndSkipTab(offSet){
  // case01
  if(window.tokenElem.value == "&amp;"){
    return (window.tokenElem.value + window.carret);
  }
  return window.tokenElem.value.substr(0, offSet+1) + window.carret + window.tokenElem.value.substr(offSet+1);
}
function setCarret(str){
  return window.carret + str;
}
window.GO_RACER_NS['chooseNextAt'] = function(idx){
  return {
    index: idx,
    el: window.tokens[idx],
    value: window.tokens[idx].innerHTML,
    offSet: 0
  }
}

window.GO_RACER_NS['prepareGameField'] = function(){
  var codeBlock = document.getElementById('code');
  // we need global
  window.carret = "<span class='carret blink'></span>";
  window.tokens = window.GO_RACER_NS.parseTokens(codeBlock);
  window.tokenElem = window.GO_RACER_NS.chooseNextAt(0);
  window.onkeypress = window.GO_RACER_NS.onKeyPress;
}
