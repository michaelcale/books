// we're applying react-like state => UI
var currentState = {
  // TODO: use searchInputFocused
  searchInputFocused: false,
  searchResults: [],
  // index within searchResults array, -1 means not selected
  selectedSearchResult: -1
};

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
  var uri = selected[0];
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
  var results = currentState.searchResults;
  var n = results.length;
  var selIdx = currentState.selectedSearchResult;
  if (n <= 0 || selIdx < 0) {
    return;
  }
  var dir = (ev.key == "ArrowUp") ? -1 : 1;
  var newIdx = selIdx + dir;
  if (newIdx >= 0 && newIdx < n) {
    setState({
      selectedSearchResult: newIdx
    });
    ev.preventDefault();
  }
}

function onKeyUp(ev) {
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
function hilightSearchResult(s, idx, len) {
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
    var title = r[1];
    var idx = r[2];
    var len = r[3];
    var html = hilightSearchResult(title, idx, len);
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

function rebuildSearchResultsUI() {
  var results = currentState.searchResults;
  var selectedIdx = currentState.selectedSearchResult;
  var el = document.getElementById("search-results");
  if (results.length == 0) {
    el.style.display = "none";
    return;
  }
  el.style.display = "block";
  var html = buildResultsHTML(results, selectedIdx);
  el.innerHTML = html;
  // TOOD: ensure selected search result is visible
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
  setState({
    searchResults: [],
    selectedSearchResult: -1
  });
}

var currentSearchTerm;

var maxSearchResults = 25;

function doSearch(searchTerm) {
  searchTerm = searchTerm.trim();
  if (searchTerm == currentSearchTerm) {
    return;
  }
  if (searchTerm.length == 0) {
    clearSearchResults();
    return;
  }
  var toFind = searchTerm.toLowerCase();
  console.log("search for:", toFind);
  var a = gBookTocSearchData; // loaded via toc_search.js
  var n = a.length;
  var res = [];
  for (var i = 0; i < n && res.length < maxSearchResults; i++) {
    var el = a[i];
    var title = el[1].toLowerCase();
    var idx = title.indexOf(toFind);
    if (idx == -1) {
      continue;
    }
    var uri = el[0];
    var resEl = [uri, el[1], idx, toFind.length];
    res.push(resEl);
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

  document.addEventListener("keyup", onKeyUp, true);
  var el = getSearchInputElement();
  el.addEventListener("input", onSearchInputChanged, true);
  document.addEventListener("mousemove", onMouseMove, true);
  document.addEventListener("click", onClick, true);
}

// parsed and DOM ready but before loading external resources like images or stylesheets
// https://javascript.info/onload-ondomcontentloaded
document.addEventListener("DOMContentLoaded", start);
