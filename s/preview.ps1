#!/usr/bin/env pwsh

# you can pass additional args like:
# -update-go-deps
Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"
function exitIfFailed { if ($LASTEXITCODE -ne 0) { exit } }

Remove-Item -Force -ErrorAction SilentlyContinue ./cmd/gen-books/gen-books.exe

Set-Location -Path cmd/gen-books
go build -o gen-books.exe
Set-Location -Path ../..
exitIfFailed

./cmd/gen-books/gen-books.exe -preview  -analytics UA-113489735-1 $args
Remove-Item -Force -ErrorAction SilentlyContinue ./cmd/gen-books/gen-books.exe
