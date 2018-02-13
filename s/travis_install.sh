#!/bin/bash
set -u -e -o pipefail -o verbose

go get -v -u github.com/netlify/netlifyctl
cd cmd
go get -v -u ./...
cd ../pkg
go get -v -u ./...
cd ..
