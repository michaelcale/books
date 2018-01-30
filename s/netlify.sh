#!/bin/bash

cd cmd/gen-books
go build -o gen-books
cd ../..
./cmd/gen-books/gen-books
