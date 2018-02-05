
function getSearchInputElement() {
    return document.getElementById("search-input");
}

function focusSearchInput() {
    getSearchInputElement().focus();
}

function unfocusSearchInput() {
    getSearchInputElement().blur();
}

function onKeyUp(ev) {
    // console.log(ev);
    if (ev.key == "/") {
        focusSearchInput();
        ev.preventDefault();
        return;
    }
    if (ev.key == "Escape") {
        unfocusSearchInput();
        ev.preventDefault();
        return;
    }
}

var currentSearchTerm;

function startSearch(searchTerm) {
    if (searchTerm == currentSearchTerm) {
        return;
    }
    console.log("search for:", searchTerm);
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
    var fn = startSearch.bind(this, s);
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
