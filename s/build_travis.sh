#!/bin/bash
set -u -e -o pipefail -o verbose

cd cmd/gen-books
go build -o gen-books
cd ../..

./cmd/gen-books/gen-books -analytics UA-113489735-1

# netlifyctl deploy
