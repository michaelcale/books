#!/usr/bin/env pwsh
Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"
function exitIfFailed { if ($LASTEXITCODE -ne 0) { exit } }

function isWindows { return $ENV:OS -eq "Windows_NT" }
$isWin = isWindows

# Write-Host "isWindows: $isWin"

Remove-Item -Force -ErrorAction SilentlyContinue ./cmd/gen-books/gen-books

Set-Location -Path cmd/gen-books
go build
Set-Location -Path ../..
exitIfFailed


./cmd/gen-books/gen-books -preview  -analytics UA-113489735-1
Remove-Item -Force -ErrorAction SilentlyContinue ./cmd/gen-books/gen-books
