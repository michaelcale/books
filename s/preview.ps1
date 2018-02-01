#!/usr/bin/env pwsh
Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"
function exitIfFailed { if ($LASTEXITCODE -ne 0) { exit } }

Remove-Item -Force -ErrorAction SilentlyContinue ./cmd/gen-books/gen-books

Set-Location -Path cmd/preview
go build -o preview
Set-Location -Path ../..
exitIfFailed

./cmd/preview/preview
Remove-Item -Force -ErrorAction SilentlyContinue ./cmd/preview/preview
