#!/usr/bin/env pwsh
Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"
function exitIfFailed { if ($LASTEXITCODE -ne 0) { exit } }

Remove-Item -Force -ErrorAction SilentlyContinue ./cmd/rename/rename

Set-Location -Path cmd/rename
go build -o rename
Set-Location -Path ../..
exitIfFailed

./cmd/rename/rename
Remove-Item -Force -ErrorAction SilentlyContinue ./cmd/rename/rename
