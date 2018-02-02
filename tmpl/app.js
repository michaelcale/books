
function start() {
    console.log("started");
}

// parsed and DOM ready but before loading external resources like images or stylesheets
// https://javascript.info/onload-ondomcontentloaded
document.addEventListener("DOMContentLoaded", start);
