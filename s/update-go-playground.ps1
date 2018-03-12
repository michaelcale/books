#!/usr/bin/env pwsh
Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"
function exitIfFailed { if ($LASTEXITCODE -ne 0) { exit } }

Remove-Item -Force -ErrorAction SilentlyContinue ./cmd/gen-books/update-playground.exe

Set-Location -Path cmd/gen-books
go build -o update-playground.exe
Set-Location -Path ../..
exitIfFailed

./cmd/gen-books/update-playground.exe -update-go-playground
Remove-Item -Force -ErrorAction SilentlyContinue ./cmd/gen-books/update-playground.exe
