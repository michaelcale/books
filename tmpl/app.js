// we're applying react-like state => UI
var currentState = {
  searchResults: [],
  searchInputFocused: false
};

var rebuildUITimer = null;
function triggerUIRebuild() {
  rebuildUITimer = null;
  rebuildUIFromState();
}

function requestRebuildUI() {
  // debounce the requests
  if (rebuildUITimer != null) {
    window.cancelAnimationFrame(rebuildUITimer);
    rebuildUITimer = null;
  }
  window.requestAnimationFrame(triggerUIRebuild);
}

function areValuesEqual(v1, v2) {
  // TODO: implement me
  return false;
}

function setState(newState) {
  var vOld, vNew;
  var stateChanged = false;
  for (var k in newState) {
    vOld = currentState[k];
    vNew = newState[k];
    if (stateChanged) {
      // avoid calling areValuesEqual if we're updating the state anyway
      currentState[k] = vNew;
    } else if (!areValuesEqual(vOld, vNew)) {
      stateChanged = true;
      currentState[k] = vNew;
    }
  }
  if (stateChanged) {
    requestRebuildUI();
  }
}

function getSearchInputElement() {
  return document.getElementById("search-input");
}

function focusAndClearSearchInput() {
  var el = getSearchInputElement();
  el.value = "";
  el.focus();
}

function unfocusAndClearSearchInput() {
  var el = getSearchInputElement();
  el.value = "";
  el.blur();
  clearSearchResults();
}

function onKeyUp(ev) {
  // console.log(ev);
  if (ev.key == "/") {
    focusAndClearSearchInput();
    ev.preventDefault();
    return;
  }
  if (ev.key == "Escape") {
    unfocusAndClearSearchInput();
    ev.preventDefault();
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
function buildResultsHTML(results) {
  var a = [];
  var n = results.length;
  for (var i = 0; i < n; i++) {
    var r = results[i];
    var title = r[1];
    var idx = r[2];
    var len = r[3];
    var html = hilightSearchResult(title, idx, len);
    var opt = {id: "search-result-no-" + i};
    var s = tagOpen("div", opt) + html + "</div>";
    a.push(s);
  }
  return a.join("\n");
}

function rebuildSearchResultsUI() {
  var results = currentState.searchResults;
  var el = document.getElementById("search-results");
  if (results.length == 0) {
    el.style.display = "none";
    return;
  }
  el.style.display = "block";
  var html = buildResultsHTML(results);
  el.innerHTML = html;
}

function rebuildUIFromState() {
  rebuildSearchResultsUI();
}

function clearSearchResults() {
  setState({
    searchResults: []
  });
}

var currentSearchTerm;

var maxSearchResults = 25;

function doSearch(searchTerm) {
  if (searchTerm == currentSearchTerm) {
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
    searchResults: res
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

function onSearchInputChanged(ev) {
  var s = ev.target.value;
  var fn = doSearch.bind(this, s);
  searchInputDebouncer(fn);
}

function startGlobalKeyListening() {
  document.addEventListener("keyup", onKeyUp, true);
}

function startInputElementListening() {
  var el = getSearchInputElement();
  el.addEventListener("input", onSearchInputChanged, true);
}

function start() {
  console.log("started");
  startGlobalKeyListening();
  startInputElementListening();
}

// parsed and DOM ready but before loading external resources like images or stylesheets
// https://javascript.info/onload-ondomcontentloaded
document.addEventListener("DOMContentLoaded", start);
