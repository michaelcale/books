#!/bin/bash
set -u -e -o pipefail -o verbose

cd cmd
go get -v -u ./...
cd ../pkg
go get -v -u ./...
cd ..
