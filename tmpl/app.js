// we're applying react-like state => UI
var currentState = {
  searchInputFocused: false,
  searchResults: [],
  // index within searchResults array, -1 means not selected
  selectedSearchResult: -1,
};

var currentSearchTerm = "";

// polyfil for Object.is
// https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Object/is
if (!Object.is) {
  Object.is = function(x, y) {
    // SameValue algorithm
    if (x === y) {
      // Steps 1-5, 7-10
      // Steps 6.b-6.e: +0 != -0
      return x !== 0 || 1 / x === 1 / y;
    } else {
      // Step 6.a: NaN == NaN
      return x !== x && y !== y;
    }
  };
}

var rebuildUITimer = null;
function triggerUIRebuild() {
  rebuildUITimer = null;
  rebuildUIFromState();
}

function requestRebuildUI(now) {
  // debounce the requests
  if (rebuildUITimer != null) {
    window.cancelAnimationFrame(rebuildUITimer);
    rebuildUITimer = null;
  }
  if (now) {
    triggerUIRebuild();
  } else {
    window.requestAnimationFrame(triggerUIRebuild);
  }
}

function setState(newState, now=false) {
  var vOld, vNew;
  var stateChanged = false;
  for (var k in newState) {
    vOld = currentState[k];
    vNew = newState[k];
    if (stateChanged) {
      // avoid calling areValuesEqual if we're updating the state anyway
      currentState[k] = vNew;
    } else if (!Object.is(vOld, vNew)) {
      stateChanged = true;
      currentState[k] = vNew;
    }
  }
  if (stateChanged) {
    requestRebuildUI(now);
  }
}

function isChapterOrArticle(s) {
  return s.startsWith("ch-") || s.startsWith("a-");
}

function navigateToSearchResult(idx) {
  var loc = window.location.pathname;
  var parts = loc.split("/");
  var lastIdx = parts.length - 1;
  var lastEl = parts[lastIdx];
  var selected = currentState.searchResults[idx];
  var el = selected[0];
  var uri = el[0];
  if (isChapterOrArticle(lastEl)) {
    parts[lastIdx] = uri;
  } else {
    parts.push(uri);
  }
  loc = parts.join("/");
  clearSearchResults();
  window.location = loc;
}

function onEnter(ev) {
  var selIdx = currentState.selectedSearchResult;
  if (selIdx == -1) {
    return;
  }
  navigateToSearchResult(selIdx);
}

function onKeySlash(ev) {
  setState({
    searchInputFocused: true
  });
  ev.preventDefault();
}

function onEscape(ev) {
  setState({
    searchInputFocused: false
  });
  ev.preventDefault();
}

function onUpDown(ev) {
  var dir = (ev.key == "ArrowUp") ? -1 : 1;
  var results = currentState.searchResults;
  var n = results.length;
  var selIdx = currentState.selectedSearchResult;
  if (n <= 0 || selIdx < 0) {
    return;
  }
  var newIdx = selIdx + dir;
  if (newIdx >= 0 && newIdx < n) {
    setState({
      selectedSearchResult: newIdx
    });
    ev.preventDefault();
  }
}

function onKeyDown(ev) {
  // console.log(ev);
  if (ev.key == "/") {
    onKeySlash(ev);
    return;
  }

  if (ev.key == "Escape") {
    onEscape(ev);
    return;
  }

  if (ev.key == "Enter") {
    onEnter(ev);
    return;
  }

  if (ev.key == "ArrowUp" || ev.key == "ArrowDown") {
    onUpDown(ev);
    return;
  }
}

// from https://github.com/component/escape-html/blob/master/index.js
var matchHtmlRegExp = /["'&<>]/;
function escapeHtml(string) {
  var str = "" + string;
  var match = matchHtmlRegExp.exec(str);

  if (!match) {
    return str;
  }

  var escape;
  var html = "";
  var index = 0;
  var lastIndex = 0;

  for (index = match.index; index < str.length; index++) {
    switch (str.charCodeAt(index)) {
      case 34: // "
        escape = "&quot;";
        break;
      case 38: // &
        escape = "&";
        break;
      case 39: // '
        escape = "&#39;";
        break;
      case 60: // <
        escape = "&lt;";
        break;
      case 62: // >
        escape = "&gt;";
        break;
      default:
        continue;
    }

    if (lastIndex !== index) {
      html += str.substring(lastIndex, index);
    }

    lastIndex = index + 1;
    html += escape;
  }

  return lastIndex !== index ? html + str.substring(lastIndex, index) : html;
}

// splits a string in two parts at a given index
// ("foobar", 2) = ["fo", "obar"]
function splitStringAt(s, idx) {
  var res = ["", ""];
  if (idx == 0) {
    res[1] = s;
  } else {
    res[0] = s.substring(0, idx);
    res[1] = s.substring(idx);
  }
  return res;
}

function span(s, cls) {
  s = escapeHtml(s);
  if (!cls) {
    return "<span>" + s + "</span>";
  }
  return "<span class='" + cls + "'>" + s + "</span>";
}

// create HTML to highlight part of s starting at idx and with length len
function hilightSearchResult(s, matches) {
  // TODO: handle multiple matches
  var idx = matches[0][0];
  var len = matches[0][1];
  var s1 = s.substring(0, idx);
  var s2 = s.substring(idx, idx + len);
  var s3 = s.substring(idx + len);
  var a = [s1, s2, s3];
  var n = a.length;
  var isHighlighted;
  var res = "";
  // alternate non-higlighted and highlihted strings
  // this allows for future showing of multiple highlights
  for (var i = 0; i < n; i++) {
    s = a[i];
    isHighlighted = i % 2 == 0; // they alternate
    if (isHighlighted) {
      res += span(s);
    } else {
      res += span(s, "hili");
    }
  }
  return res;
}

function attr(name, val) {
  return name + "='" + val + "'";
}

function tagOpen(name, cls, id) {
  var s = "<" + name;
  if (cls) {
    s += " " + attr("class", cls);
  }
  if (id) {
    s += " " + attr("id", id);
  }
  return s + ">";
}

function div(html, opt) {
  var s = tagOpen("div", opt && opt.cls, opt && opt.id);
  return s + html + "</div>";
}

// results is in format:
// [uri, title, match idx, match len];
function buildResultsHTML(results, selectedIdx) {
  var a = [];
  var n = results.length;
  for (var i = 0; i < n; i++) {
    var r = results[i];
    var el = r[0];
    var title = el[1];
    var matches = r[1];

    var html = hilightSearchResult(title, matches);
    var opt = {
      id: "search-result-no-" + i,
      cls: "search-result",
    };
    if (i == selectedIdx) {
      opt.cls += " search-result-selected";
    }
    var s = div(html, opt);
    a.push(s);
  }
  return a.join("\n");
}

// https://stackoverflow.com/questions/6215779/scroll-if-element-is-not-visible
// TODO: improve on scroll up
function scrollIntoViewIfOutOfView(el) {
  var parent = el.parentElement;
  var topOfPage = parent.scrollTop;
  var heightOfPage = parent.clientHeight;
  var elY = 0;
  var elH = 0;
  if (document.layers) { // NS4
      elY = el.y;
      elH = el.height;
  } else {
      for(var p=el; p&&p.tagName!='BODY'; p=p.offsetParent){
          elY += p.offsetTop;
      }
      elH = el.offsetHeight;
  }
  if ((topOfPage + heightOfPage) < (elY + elH)) {
      el.scrollIntoView(false);
  }
  else if (elY < topOfPage) {
      el.scrollIntoView(true);
  }
}

function rebuildSearchResultsUI() {
  var html;
  var results = currentState.searchResults;
  var selectedIdx = currentState.selectedSearchResult;
  var el = document.getElementById("search-results");
  var blurOverlay = document.getElementById("blur-overlay");
  if (results.length == 0) {
    if (currentSearchTerm == "") {
      el.style.display = "none";
      blurOverlay.style.display = "none";
    } else {
      el.style.display = "block";
      blurOverlay.style.display = "block";
      html = "<div class='no-search-results'>No search results for '" + currentSearchTerm + "'</div>";
      el.innerHTML = html;
    }
    return;
  }
  el.style.display = "block";
  blurOverlay.style.display = "block";
  html = buildResultsHTML(results, selectedIdx);
  el.innerHTML = html;

  // ensure element is scrolled into view
  window.requestAnimationFrame(() => {
    var id = "search-result-no-" + selectedIdx;
    var el = document.getElementById(id);
    scrollIntoViewIfOutOfView(el);
  });
}

function getSearchInputElement() {
  return document.getElementById("search-input");
}

function setSearchInputFocus() {
  // console.log("setSearchInputFocus:", currentState.searchInputFocused);
  var el = getSearchInputElement();
  var wantsFocus = currentState.searchInputFocused;
  var isFocused = document.activeElement === el;
  if (isFocused == wantsFocus) {
    return;
  }
  if (wantsFocus) {
    el.value = "";
    el.focus();
  } else {
    el.value = "";
    el.blur();
    clearSearchResults();
  }
}

function rebuildUIFromState() {
  setSearchInputFocus();
  rebuildSearchResultsUI();
}

function clearSearchResults() {
  currentSearchTerm = "";
  setState({
    searchResults: [],
    selectedSearchResult: -1
  });
}

var maxSearchResults = 25;

// el is [idx, len]
// sort by idx.
// if idx is the same, sort by reverse len
// (i.e. bigger len is first)
function sortSearchByIdx(el1, el2) {
  var res = el1[0] - el2[0];
  if (res == 0) {
    res = el2[1] - el1[1];
  }
  return res;
}

// [[idx, len], ...]
// sort by idx, if there is an overlap, drop overlapped elements
function sortSearchMatches(a) {
  if (a.length < 2) {
    return a;
  }
  a.sort(sortSearchByIdx);
  var lastIdx = a[0][0] + a[0][1]; // start + len
  var n = a.length;
  var res = [a[0]];
  for (var i = 1; i < n; i++) {
    var el = a[i];
    var idx = el[0];
    var len = el[1];
    if (idx >= lastIdx) {
      res.push(el);
      lastIdx = idx + len;
    }
  }
  return a;
}

// returns an array of [[idx, len], ...]
// searchTerms might be an empty array
function searchMatch(s, toFind, toFindArr) {
  s = s.toLowerCase();
  var idx = s.indexOf(toFind);
  if (idx != -1) {
    return [[idx, toFind.length]];
  }
  if (!toFindArr) {
    return null;
  }
  var n = toFindArr.length;
  var res = Array(n)
  for (var i = 0; i < n; i++) {
    toFind = toFindArr[i];
    idx = s.indexOf(toFind);
    if (idx == -1) {
      return null;
    }
    res[i] = [idx, toFind.length];
  }
  return sortSearchMatches(res);
}

function notEmptyString(s) {
  return s.length > 0;
}

// if search term is multiple words like "blank id",
// we search for both the exact match and if we match all
// terms ("blank", "id") separately
function doSearch(searchTerm) {
  searchTerm = searchTerm.trim();
  if (searchTerm == currentSearchTerm) {
    return;
  }
  currentSearchTerm = searchTerm;
  if (searchTerm.length == 0) {
    clearSearchResults();
    return;
  }

  searchTerm = searchTerm.toLowerCase();
  var searchTerms = searchTerm.split(" ").filter(notEmptyString);

  console.log("search for:", searchTerm);
  var a = gBookTocSearchData; // loaded via toc_search.js
  var n = a.length;
  var res = [];
  for (var i = 0; i < n && res.length < maxSearchResults; i++) {
    var el = a[i];
    var match = searchMatch(el[1], searchTerm, searchTerms);
    if (!match) {
      continue;
    }
    res.push([el, match]);
  }
  console.log("search results:", res);
  setState({
    searchResults: res,
    selectedSearchResult: 0
  });
}

// returns a debouncer function. Usage:
// var debouncer = makeDebouncer(250);
// function fn() { ... }
// debouncer(fn)
function makeDebouncer(timeInMs) {
  let interval;
  return function(f) {
    clearTimeout(interval);
    interval = setTimeout(() => {
      interval = null;
      f();
    }, timeInMs);
  };
}
// TODO: maybe just use debouncer from https://gist.github.com/nmsdvid/8807205
// and do addEventListener("input", debounce(onSearchInputChanged, 250, false))
var searchInputDebouncer = makeDebouncer(250);

function getIdxFromSearchResultElementId(id) {
  if (!id) {
    return -1;
  }
  if (!id.startsWith("search-result-no-")) {
    return -1;
  }
  var parts = id.split("-");
  var nStr = parts[parts.length-1];
  var n = parseInt(nStr, 10);
  return isNaN(n) ? -1 : n;
}

function findEnclosingResultNode(el) {
  while (el) {
    var idx = getIdxFromSearchResultElementId(el.id);
    if (idx > 0) {
      return idx;
    }
    el = el.parentNode;
  }
  return -1;
}

// if search result item is
function onClick(ev) {
  var el = ev.target;
  console.log("el:", el);
  var idx = findEnclosingResultNode(el);
  if (idx < 0) {
    setState({
      selectedSearchResult: -1,
    })
    console.log("stopped propagation");
    ev.stopPropagation();
    return;
  }
  if (idx >= 0) {
    navigateToSearchResult(idx);
  }
  ev.stopPropagation();
}

// when we're over elements with id "search-result-no-${id}", set this one
// as selected element
function onMouseMove(ev) {
  var el = ev.target;
  var idx = getIdxFromSearchResultElementId(el.id);
  if (idx < 0) {
    return;
  }
  //console.log("ev.target:", el, "id:", el.id, "idx:", idx);
  setState({
    selectedSearchResult: idx,
  }, true);
  ev.stopPropagation();
}

function onSearchInputChanged(ev) {
  var s = ev.target.value;
  var fn = doSearch.bind(this, s);
  searchInputDebouncer(fn);
}

function start() {
  console.log("started");

  document.addEventListener("keydown", onKeyDown, true);
  var el = getSearchInputElement();
  el.addEventListener("input", onSearchInputChanged, true);
  document.addEventListener("mousemove", onMouseMove, true);
  document.addEventListener("click", onClick, true);
}

// parsed and DOM ready but before loading external resources like images or stylesheets
// https://javascript.info/onload-ondomcontentloaded
document.addEventListener("DOMContentLoaded", start);
