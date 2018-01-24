#!/usr/bin/env pwsh
Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"
function exitIfFailed { if ($LASTEXITCODE -ne 0) { exit } }

Remove-Item -Force -Recurse -ErrorAction SilentlyContinue ./books
Remove-Item -Force -Recurse -ErrorAction SilentlyContinue ./books_html/book

# import from so
Remove-Item -Force -ErrorAction SilentlyContinue ./cmd/import-stack-overflow/import-stack-overflow

Set-Location -Path cmd/import-stack-overflow
go build -o import-stack-overflow
Set-Location -Path ../..
exitIfFailed

./cmd/import-stack-overflow/import-stack-overflow
Remove-Item -Force -ErrorAction SilentlyContinue ./cmd/import-stack-overflow/import-stack-overflow

# rebuild html
Remove-Item -Force -ErrorAction SilentlyContinue ./cmd/gen-books/gen-books

Set-Location -Path cmd/gen-books
go build -o gen-books
Set-Location -Path ../..
exitIfFailed

./cmd/gen-books/gen-books
Remove-Item -Force -ErrorAction SilentlyContinue ./cmd/gen-books/gen-books
