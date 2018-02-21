#!/bin/bash
set -u -e -o pipefail
find cmd pkg -name "*.go" | xargs wc -l
echo ""
wc -l "tmpl/app.js"
echo ""
wc -l "tmpl/main.css"

