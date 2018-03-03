// we're applying react-like state => UI
var currentState = {
  searchInputFocused: false,
  searchResults: [],
  // index within searchResults array, -1 means not selected
  selectedSearchResultIdx: -1
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

// accessor functions for items in gBookToc array:
// 	[${chapter or aticle url}, ${parentIdx}, ${title}, ${synonim 1}, ${synonim 2}, ...],
// as generated in gen_book_toc_search.go and stored in books/${book}/toc_search.js

var itemIdxIsExpanded = 0;
var itemIdxURL = 1;
var itemIdxParent = 2;
var itemIdxFirstChild = 3;
var itemIdxTitle = 4;
var itemIdxFirstSynonym = 5;

function tocItemIsExpanded(item) {
  return item[itemIdxIsExpanded];
}

function tocItemSetIsExpanded(item, isExpanded) {
  item[itemIdxIsExpanded] = isExpanded;
}

function tocItemURL(item) {
  while (true) {
    var uri = item[itemIdxURL];
    if (uri != "") {
      return uri;
    }
    item = tocItemParent(item);
    if (!item) {
      return "";
    }
  }
}

function tocItemFirstChildIdx(item) {
  return item[itemIdxFirstChild];
}

function tocItemHasChildren(item) {
  return tocItemFirstChildIdx(item) != -1;
}

function tocItemParent(item) {
  var idx = tocItemParentIdx(item);
  if (idx == -1) {
    return null;
  }
  return gBookToc[idx];
}

function tocItemIsRoot(item) {
  return tocItemParentIdx(item) == -1;
}

function tocItemParentIdx(item) {
  return item[itemIdxParent];
}

function tocItemTitle(item) {
  return item[itemIdxTitle];
}

// all searchable items: title + search synonyms
function tocItemSearchable(item) {
  return item.slice(itemIdxTitle);
}

// from https://github.com/component/escape-html/blob/master/index.js
var matchHtmlRegExp = /["'&<>]/;
function escapeHTML(string) {
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
// ("foobar", 2) => ["fo", "obar"]
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

function tagClose(tagName) {
  return "</" + tagName + ">";
}

function inTag(tagName, contentHTML, cls, id) {
  return tagOpen(tagName, cls, id) + contentHTML + tagClose(tagName);
}

function inTagRaw(tagName, content, cls, id) {
  var contentHTML = escapeHTML(content);
  return tagOpen(tagName, cls, id) + contentHTML + tagClose(tagName);
}
function attr(name, val) {
  return name + "='" + val + "'";
}

function span(s, cls) {
  return inTagRaw("span", s, cls);
}

function div(html, opt) {
  return inTag("div", html, opt && opt.cls, opt && opt.id);
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

function setState(newState, now = false) {
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

function isChapterOrArticleURL(s) {
  return s.startsWith("a-");
}

function getLocationLastElement() {
  var loc = window.location.pathname;
  var parts = loc.split("/");
  var lastIdx = parts.length - 1;
  return parts[lastIdx];
}

function getLocationLastElementWithHash() {
  var loc = window.location.pathname;
  var parts = loc.split("/");
  var lastIdx = parts.length - 1;
  return parts[lastIdx] + window.location.hash;
}

function navigateToSearchResult(idx) {
  var loc = window.location.pathname;
  var parts = loc.split("/");
  var lastIdx = parts.length - 1;
  var lastURL = parts[lastIdx];
  var selected = currentState.searchResults[idx];
  var tocItem = selected.tocItem;

  // either replace chapter/article url or append to book url
  var uri = tocItemURL(tocItem);
  if (isChapterOrArticleURL(lastURL)) {
    parts[lastIdx] = uri;
  } else {
    parts.push(uri);
  }
  loc = parts.join("/");
  clearSearchResults();
  window.location = loc;
}

// create HTML to highlight part of s starting at idx and with length len
function hilightSearchResult(txt, matches) {
  var prevIdx = 0;
  var n = matches.length;
  var res = "";
  var s = "";
  // alternate non-higlighted and highlihted strings
  for (var i = 0; i < n; i++) {
    var el = matches[i];
    var idx = el[0];
    var len = el[1];

    var nonHilightLen = idx - prevIdx;
    if (nonHilightLen > 0) {
      s = txt.substring(prevIdx, prevIdx + nonHilightLen);
      res += span(s);
    }
    s = txt.substring(idx, idx + len);
    res += span(s, "hili");
    prevIdx = idx + len;
  }
  var txtLen = txt.length;
  nonHilightLen = txtLen - prevIdx;
  if (nonHilightLen > 0) {
    s = txt.substring(prevIdx, prevIdx + nonHilightLen);
    res += span(s);
  }
  return res;
}

// return true if term is a search synonym inside tocItem
function isMatchSynonym(tocItem, term) {
  term = term.toLowerCase();
  var title = tocItemTitle(tocItem).toLowerCase();
  return title != term;
}

function getParentTitle(tocItem) {
  var parentIdx = tocItemParentIdx(tocItem);
  if (parentIdx == -1) {
    return null;
  }
  var parent = gBookToc[parentIdx];
  return tocItemTitle(parent);
}

// if search matched synonym returns "${chapterTitle} / ${articleTitle}"
// otherwise empty string
function getArticlePath(tocItem, term) {
  if (!isMatchSynonym(tocItem, term)) {
    return null;
  }
  var title = tocItemTitle(tocItem);
  var parentTitle = getParentTitle(tocItem);
  if (parentTitle == "") {
    return title;
  }
  return parentTitle + " / " + title;
}

/* results is array of items:
{
  tocItem: [],
  term: "",
  match: [[idx, len], ...],
}
*/
function buildResultsHTML(results, selectedIdx) {
  var a = [];
  var n = results.length;
  for (var i = 0; i < n; i++) {
    var r = results[i];
    var tocItem = r.tocItem;
    var term = r.term;
    var matches = r.match;

    var html = hilightSearchResult(term, matches);
    var articlePath = getArticlePath(tocItem, term);
    if (articlePath) {
      var s = "in: " + articlePath;
      html += " " + inTagRaw("span", s, "search-result-in");
    } else {
      var parentTitle = getParentTitle(tocItem);
      if (parentTitle) {
        var s = "in: " + parentTitle;
        html += " " + inTagRaw("span", s, "search-result-in");
      }
    }

    var opt = {
      id: "search-result-no-" + i,
      cls: "search-result"
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
  if (document.layers) {
    // NS4
    elY = el.y;
    elH = el.height;
  } else {
    for (var p = el; p && p.tagName != "BODY"; p = p.offsetParent) {
      elY += p.offsetTop;
    }
    elH = el.offsetHeight;
  }
  if (topOfPage + heightOfPage < elY + elH) {
    el.scrollIntoView(false);
  } else if (elY < topOfPage) {
    el.scrollIntoView(true);
  }
}

function rebuildSearchResultsUI() {
  var html;
  var results = currentState.searchResults;
  var selectedIdx = currentState.selectedSearchResultIdx;
  var searchWindow = document.getElementById("search-results-window");
  var blurOverlay = document.getElementById("blur-overlay");
  var searchResults = document.getElementById("search-results");
  if (results.length == 0) {
    if (currentSearchTerm == "") {
      searchWindow.style.display = "none";
      blurOverlay.style.display = "none";
    } else {
      searchWindow.style.display = "block";
      blurOverlay.style.display = "block";
      html =
        "<div class='no-search-results'>No search results for '" +
        currentSearchTerm +
        "'</div>";
      searchResults.innerHTML = html;
    }
    return;
  }
  searchWindow.style.display = "block";
  blurOverlay.style.display = "block";
  html = buildResultsHTML(results, selectedIdx);
  searchResults.innerHTML = html;

  // ensure element is scrolled into view
  window.requestAnimationFrame(() => {
    if (selectedIdx < 0) {
      return;
    }
    var id = "search-result-no-" + selectedIdx;
    var el = document.getElementById(id);
    scrollIntoViewIfOutOfView(el);
  });
}

function getItemsIdxForParent(parentIdx) {
  var res = [];
  var n = gBookToc.length;
  for (var i = 0; i < n; i++) {
    var tocItem = gBookToc[i];
    if (tocItemParentIdx(tocItem) == parentIdx) {
      res.push(i);
    }
  }
  return res;
}

function buildTOCHTMLLevel(level, parentIdx) {
  var opt = {};
  var opt, tocItemIdx, tocItem, parent;
  var itemsIdx = getItemsIdxForParent(parentIdx);
  if (itemsIdx.length == 0) {
    return "";
  }
  var currURI = getLocationLastElementWithHash();
  //console.log("currURI:", currURI);
  var html = "";
  var n = itemsIdx.length;
  for (var i = 0; i < n; i++) {
    tocItemIdx = itemsIdx[i];
    tocItem = gBookToc[tocItemIdx];
    opt.cls = "lvl" + level;
    var title = tocItemTitle(tocItem);
    var uri = tocItemURL(tocItem);
    // var arrow = "&#x25B8; "; // http://graphemica.com/%E2%96%B8
    var arrow = "&#x25BA; ";
    if (tocItemIsExpanded(tocItem)) {
      //arrow = "&#x25BE; "; // http://graphemica.com/%E2%96%BE
      arrow = "&#x25BC; ";
    }

    if (!tocItemHasChildren(tocItem)) {
      arrow = "";
    }

    if (tocItemIsExpanded(tocItem)) {
      if (opt.cls == "") {
        opt.cls = "bold";
      } else {
        opt.cls += " bold";
      }
    }
    var el = '<a href="' + uri + '">' + arrow + title + "</a>";
    el = div(el, opt);

    // console.log("uri:", uri);

    if (currURI === uri) {
      if (!tocItemHasChildren(tocItem)) {
        el = div(escapeHTML(title), opt);
      }
    }

    html += el;

    if (tocItemIsExpanded(tocItem)) {
      var htmlChild = buildTOCHTMLLevel(level + 1, tocItemIdx);
      html += htmlChild;
    }
  }
  return html;
}

function buildTOCHTML() {
  return buildTOCHTMLLevel(0, -1);
}

function setIsExpandedUpwards(i) {
  var tocItem = gBookToc[i];
  tocItemSetIsExpanded(tocItem, true);
  tocItem = tocItemParent(tocItem);
  while (tocItem != null) {
    tocItemSetIsExpanded(tocItem, true);
    tocItem = tocItemParent(tocItem);
  }
}

function tocUnexpandAll() {
  var tocItem;
  var n = gBookToc.length;
  for (var i = 0; i < n; i++) {
    tocItem = gBookToc[i];
    tocItemSetIsExpanded(tocItem, false);
  }
}

function setTocExpandedForCurrentURL() {
  var currURI = getLocationLastElementWithHash();
  var n = gBookToc.length;
  var tocItem, uri;
  for (var i = 0; i < n; i++) {
    tocItem = gBookToc[i];
    uri = tocItemURL(tocItem);
    if (uri === currURI) {
      setIsExpandedUpwards(i);
      return;
    }
  }
}

function locationHashChanged(e) {
  tocUnexpandAll();
  setTocExpandedForCurrentURL();
  recreateTOC();
}

function createTOC() {
  var el = document.getElementById("toc");
  var html = buildTOCHTML();
  el.innerHTML = html;
}

function recreateTOC() {
  var el = document.getElementById("toc");
  var scrollTop = el.scrollTop;
  createTOC();
  el = document.getElementById("toc");
  el.scrollTop = scrollTop;
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
    selectedSearchResultIdx: -1
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

// searches s for toFind and toFindArr.
// returns null if no match
// returns array of [idx, len] position in $s where $toFind or $toFindArr matches
function searchMatch(s, toFind, toFindArr) {
  s = s.toLowerCase();

  // try exact match
  var idx = s.indexOf(toFind);
  if (idx != -1) {
    return [[idx, toFind.length]];
  }

  // now see if matches for search for AND of components in toFindArr
  if (!toFindArr) {
    return null;
  }

  var n = toFindArr.length;
  var res = Array(n);
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

/*
returns null if no match
returns: {
  term: "",
  match: [[idx, len], ...]
}
*/
function searchMatchMulti(toSearchArr, toFind) {
  var toFindArr = toFind.split(" ").filter(notEmptyString);
  var n = toSearchArr.length;
  for (var i = 0; i < n; i++) {
    var toSearch = toSearchArr[i];
    var match = searchMatch(toSearch, toFind, toFindArr);
    if (match) {
      return {
        term: toSearch,
        match: match,
        tocItem: null // will be filled later
      };
    }
  }
  return null;
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
  searchTerm = searchTerm.toLowerCase();
  currentSearchTerm = searchTerm;
  if (searchTerm.length == 0) {
    clearSearchResults();
    return;
  }

  // console.log("search for:", searchTerm);
  var a = gBookToc; // loaded via toc_search.js, generated in gen_book_toc_search.go
  var n = a.length;
  var res = [];
  for (var i = 0; i < n && res.length < maxSearchResults; i++) {
    var tocItem = a[i];
    var searchable = tocItemSearchable(tocItem);
    var match = searchMatchMulti(searchable, searchTerm);
    if (!match) {
      continue;
    }
    match.tocItem = tocItem;
    res.push(match);
  }
  // console.log("search results:", res);
  setState({
    searchResults: res,
    selectedSearchResultIdx: 0
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
  var nStr = parts[parts.length - 1];
  var n = parseInt(nStr, 10);
  return isNaN(n) ? -1 : n;
}

function findEnclosingResultNode(el) {
  while (el) {
    var idx = getIdxFromSearchResultElementId(el.id);
    if (idx >= 0) {
      return idx;
    }
    el = el.parentNode;
  }
  return -1;
}

// if search result item is
function onClick(ev) {
  var el = ev.target;
  var idx = findEnclosingResultNode(el);
  // console.log("el:", el, "idx:", idx);
  if (idx < 0) {
    setState({
      selectedSearchResultIdx: -1
    });
    return;
  }
  navigateToSearchResult(idx);
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
  setState(
    {
      selectedSearchResultIdx: idx
    },
    true
  );
  ev.stopPropagation();
}

function onEnter(ev) {
  var selIdx = currentState.selectedSearchResultIdx;
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
  var dir = ev.key == "ArrowUp" ? -1 : 1;
  var results = currentState.searchResults;
  var n = results.length;
  var selIdx = currentState.selectedSearchResultIdx;
  if (n <= 0 || selIdx < 0) {
    return;
  }
  var newIdx = selIdx + dir;
  if (newIdx >= 0 && newIdx < n) {
    setState({
      selectedSearchResultIdx: newIdx
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

function onSearchInputChanged(ev) {
  var s = ev.target.value;
  var fn = doSearch.bind(this, s);
  searchInputDebouncer(fn);
}

function start() {
  //console.log("started");

  document.addEventListener("keydown", onKeyDown, true);
  var el = getSearchInputElement();
  el.addEventListener("input", onSearchInputChanged, true);
  document.addEventListener("mousemove", onMouseMove, true);
  document.addEventListener("click", onClick, false);

  // if this is chapter or article, we generate toc
  var uri = getLocationLastElement();
  if (isChapterOrArticleURL(uri)) {
    window.onhashchange = locationHashChanged;
    tocUnexpandAll();
    setTocExpandedForCurrentURL();
    createTOC();
  }
}

// we don't want to run javascript on about etc. pages
function isAppPage() {
  var loc = window.location.pathname;
  return loc.indexOf("essential/") != -1;
}

if (isAppPage()) {
  // we don't want this in e.g. about page
  document.addEventListener("DOMContentLoaded", start);
}

/*
Font awesome JavaScript from https://raw.githubusercontent.com/FortAwesome/Font-Awesome/master/svg-with-js/js/fa-regular.js
*/

(function() {
  var _WINDOW = {};
  try {
    if (typeof window !== "undefined") _WINDOW = window;
  } catch (e) {}

  var _ref = _WINDOW.navigator || {};
  var _ref$userAgent = _ref.userAgent;
  var userAgent = _ref$userAgent === undefined ? "" : _ref$userAgent;

  var WINDOW = _WINDOW;

  var IS_IE = ~userAgent.indexOf("MSIE") || ~userAgent.indexOf("Trident/");

  var NAMESPACE_IDENTIFIER = "___FONT_AWESOME___";

  var PRODUCTION = (function() {
    try {
      return undefined === "production";
    } catch (e) {
      return false;
    }
  })();

  var oneToTen = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
  var oneToTwenty = oneToTen.concat([11, 12, 13, 14, 15, 16, 17, 18, 19, 20]);

  var RESERVED_CLASSES = [
    "xs",
    "sm",
    "lg",
    "fw",
    "ul",
    "li",
    "border",
    "pull-left",
    "pull-right",
    "spin",
    "pulse",
    "rotate-90",
    "rotate-180",
    "rotate-270",
    "flip-horizontal",
    "flip-vertical",
    "stack",
    "stack-1x",
    "stack-2x",
    "inverse",
    "layers",
    "layers-text",
    "layers-counter"
  ]
    .concat(
      oneToTen.map(function(n) {
        return n + "x";
      })
    )
    .concat(
      oneToTwenty.map(function(n) {
        return "w-" + n;
      })
    );

  function bunker(fn) {
    try {
      fn();
    } catch (e) {
      if (!PRODUCTION) {
        throw e;
      }
    }
  }

  var w = WINDOW || {};

  if (!w[NAMESPACE_IDENTIFIER]) w[NAMESPACE_IDENTIFIER] = {};
  if (!w[NAMESPACE_IDENTIFIER].styles) w[NAMESPACE_IDENTIFIER].styles = {};
  if (!w[NAMESPACE_IDENTIFIER].hooks) w[NAMESPACE_IDENTIFIER].hooks = {};
  if (!w[NAMESPACE_IDENTIFIER].shims) w[NAMESPACE_IDENTIFIER].shims = [];

  var namespace = w[NAMESPACE_IDENTIFIER];

  var _extends =
    Object.assign ||
    function(target) {
      for (var i = 1; i < arguments.length; i++) {
        var source = arguments[i];

        for (var key in source) {
          if (Object.prototype.hasOwnProperty.call(source, key)) {
            target[key] = source[key];
          }
        }
      }

      return target;
    };

  function define(prefix, icons) {
    var normalized = Object.keys(icons).reduce(function(acc, iconName) {
      var icon = icons[iconName];
      var expanded = !!icon.icon;

      if (expanded) {
        acc[icon.iconName] = icon.icon;
      } else {
        acc[iconName] = icon;
      }
      return acc;
    }, {});

    if (typeof namespace.hooks.addPack === "function") {
      namespace.hooks.addPack(prefix, normalized);
    } else {
      namespace.styles[prefix] = _extends(
        {},
        namespace.styles[prefix] || {},
        normalized
      );
    }

    /**
     * Font Awesome 4 used the prefix of `fa` for all icons. With the introduction
     * of new styles we needed to differentiate between them. Prefix `fa` is now an alias
     * for `fas` so we'll easy the upgrade process for our users by automatically defining
     * this as well.
     */
    if (prefix === "fas") {
      define("fa", icons);
    }
  }

  var icons_regular = {
    edit: [
      576,
      512,
      [],
      "f044",
      "M402.3 344.9l32-32c5-5 13.7-1.5 13.7 5.7V464c0 26.5-21.5 48-48 48H48c-26.5 0-48-21.5-48-48V112c0-26.5 21.5-48 48-48h273.5c7.1 0 10.7 8.6 5.7 13.7l-32 32c-1.5 1.5-3.5 2.3-5.7 2.3H48v352h352V350.5c0-2.1.8-4.1 2.3-5.6zm156.6-201.8L296.3 405.7l-90.4 10c-26.2 2.9-48.5-19.2-45.6-45.6l10-90.4L432.9 17.1c22.9-22.9 59.9-22.9 82.7 0l43.2 43.2c22.9 22.9 22.9 60 .1 82.8zM460.1 174L402 115.9 216.2 301.8l-7.3 65.3 65.3-7.3L460.1 174zm64.8-79.7l-43.2-43.2c-4.1-4.1-10.8-4.1-14.8 0L436 82l58.1 58.1 30.9-30.9c4-4.2 4-10.8-.1-14.9z"
    ],
    star: [
      576,
      512,
      [],
      "f005",
      "M528.1 171.5L382 150.2 316.7 17.8c-11.7-23.6-45.6-23.9-57.4 0L194 150.2 47.9 171.5c-26.2 3.8-36.7 36.1-17.7 54.6l105.7 103-25 145.5c-4.5 26.3 23.2 46 46.4 33.7L288 439.6l130.7 68.7c23.2 12.2 50.9-7.4 46.4-33.7l-25-145.5 105.7-103c19-18.5 8.5-50.8-17.7-54.6zM388.6 312.3l23.7 138.4L288 385.4l-124.3 65.3 23.7-138.4-100.6-98 139-20.2 62.2-126 62.2 126 139 20.2-100.6 98z"
    ]
  };

  bunker(function() {
    define("far", icons_regular);
  });

  var icons_brands = {
    github: [
      496,
      512,
      [],
      "f09b",
      "M165.9 397.4c0 2-2.3 3.6-5.2 3.6-3.3.3-5.6-1.3-5.6-3.6 0-2 2.3-3.6 5.2-3.6 3-.3 5.6 1.3 5.6 3.6zm-31.1-4.5c-.7 2 1.3 4.3 4.3 4.9 2.6 1 5.6 0 6.2-2s-1.3-4.3-4.3-5.2c-2.6-.7-5.5.3-6.2 2.3zm44.2-1.7c-2.9.7-4.9 2.6-4.6 4.9.3 2 2.9 3.3 5.9 2.6 2.9-.7 4.9-2.6 4.6-4.6-.3-1.9-3-3.2-5.9-2.9zM244.8 8C106.1 8 0 113.3 0 252c0 110.9 69.8 205.8 169.5 239.2 12.8 2.3 17.3-5.6 17.3-12.1 0-6.2-.3-40.4-.3-61.4 0 0-70 15-84.7-29.8 0 0-11.4-29.1-27.8-36.6 0 0-22.9-15.7 1.6-15.4 0 0 24.9 2 38.6 25.8 21.9 38.6 58.6 27.5 72.9 20.9 2.3-16 8.8-27.1 16-33.7-55.9-6.2-112.3-14.3-112.3-110.5 0-27.5 7.6-41.3 23.6-58.9-2.6-6.5-11.1-33.3 2.6-67.9 20.9-6.5 69 27 69 27 20-5.6 41.5-8.5 62.8-8.5s42.8 2.9 62.8 8.5c0 0 48.1-33.6 69-27 13.7 34.7 5.2 61.4 2.6 67.9 16 17.7 25.8 31.5 25.8 58.9 0 96.5-58.9 104.2-114.8 110.5 9.2 7.9 17 22.9 17 46.4 0 33.7-.3 75.4-.3 83.6 0 6.5 4.6 14.4 17.3 12.1C428.2 457.8 496 362.9 496 252 496 113.3 383.5 8 244.8 8zM97.2 352.9c-1.3 1-1 3.3.7 5.2 1.6 1.6 3.9 2.3 5.2 1 1.3-1 1-3.3-.7-5.2-1.6-1.6-3.9-2.3-5.2-1zm-10.8-8.1c-.7 1.3.3 2.9 2.3 3.9 1.6 1 3.6.7 4.3-.7.7-1.3-.3-2.9-2.3-3.9-2-.6-3.6-.3-4.3.7zm32.4 35.6c-1.6 1.3-1 4.3 1.3 6.2 2.3 2.3 5.2 2.6 6.5 1 1.3-1.3.7-4.3-1.3-6.2-2.2-2.3-5.2-2.6-6.5-1zm-11.4-14.7c-1.6 1-1.6 3.6 0 5.9 1.6 2.3 4.3 3.3 5.6 2.3 1.6-1.3 1.6-3.9 0-6.2-1.4-2.3-4-3.3-5.6-2z"
    ],
    twitter: [
      512,
      512,
      [],
      "f099",
      "M459.37 151.716c.325 4.548.325 9.097.325 13.645 0 138.72-105.583 298.558-298.558 298.558-59.452 0-114.68-17.219-161.137-47.106 8.447.974 16.568 1.299 25.34 1.299 49.055 0 94.213-16.568 130.274-44.832-46.132-.975-84.792-31.188-98.112-72.772 6.498.974 12.995 1.624 19.818 1.624 9.421 0 18.843-1.3 27.614-3.573-48.081-9.747-84.143-51.98-84.143-102.985v-1.299c13.969 7.797 30.214 12.67 47.431 13.319-28.264-18.843-46.781-51.005-46.781-87.391 0-19.492 5.197-37.36 14.294-52.954 51.655 63.675 129.3 105.258 216.365 109.807-1.624-7.797-2.599-15.918-2.599-24.04 0-57.828 46.782-104.934 104.934-104.934 30.213 0 57.502 12.67 76.67 33.137 23.715-4.548 46.456-13.32 66.599-25.34-7.798 24.366-24.366 44.833-46.132 57.827 21.117-2.273 41.584-8.122 60.426-16.243-14.292 20.791-32.161 39.308-52.628 54.253z"
    ]
  };

  bunker(function() {
    define("fab", icons_brands);
  });

  var icons_solid = {
    home: [
      576,
      512,
      [],
      "f015",
      "M488 312.7V456c0 13.3-10.7 24-24 24H348c-6.6 0-12-5.4-12-12V356c0-6.6-5.4-12-12-12h-72c-6.6 0-12 5.4-12 12v112c0 6.6-5.4 12-12 12H112c-13.3 0-24-10.7-24-24V312.7c0-3.6 1.6-7 4.4-9.3l188-154.8c4.4-3.6 10.8-3.6 15.3 0l188 154.8c2.7 2.3 4.3 5.7 4.3 9.3zm83.6-60.9L488 182.9V44.4c0-6.6-5.4-12-12-12h-56c-6.6 0-12 5.4-12 12V117l-89.5-73.7c-17.7-14.6-43.3-14.6-61 0L4.4 251.8c-5.1 4.2-5.8 11.8-1.6 16.9l25.5 31c4.2 5.1 11.8 5.8 16.9 1.6l235.2-193.7c4.4-3.6 10.8-3.6 15.3 0l235.2 193.7c5.1 4.2 12.7 3.5 16.9-1.6l25.5-31c4.2-5.2 3.4-12.7-1.7-16.9z"
    ],
    edit: [
      576,
      512,
      [],
      "f044",
      "M402.6 83.2l90.2 90.2c3.8 3.8 3.8 10 0 13.8L274.4 405.6l-92.8 10.3c-12.4 1.4-22.9-9.1-21.5-21.5l10.3-92.8L388.8 83.2c3.8-3.8 10-3.8 13.8 0zm162-22.9l-48.8-48.8c-15.2-15.2-39.9-15.2-55.2 0l-35.4 35.4c-3.8 3.8-3.8 10 0 13.8l90.2 90.2c3.8 3.8 10 3.8 13.8 0l35.4-35.4c15.2-15.3 15.2-40 0-55.2zM384 346.2V448H64V128h229.8c3.2 0 6.2-1.3 8.5-3.5l40-40c7.6-7.6 2.2-20.5-8.5-20.5H48C21.5 64 0 85.5 0 112v352c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V306.2c0-10.7-12.9-16-20.5-8.5l-40 40c-2.2 2.3-3.5 5.3-3.5 8.5z"
    ]
  };

  bunker(function() {
    define("fas", icons_solid);
  });
})();
