#!/usr/bin/env pwsh
Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"
function exitIfFailed { if ($LASTEXITCODE -ne 0) { exit } }

Remove-Item -Force -ErrorAction SilentlyContinue ./cmd/import-stack-overflow/import-stack-overflow

Set-Location -Path cmd/import-stack-overflow
go build -o import-stack-overflow
Set-Location -Path ../..
exitIfFailed

./cmd/import-stack-overflow/import-stack-overflow
Remove-Item -Force -ErrorAction SilentlyContinue ./cmd/import-stack-overflow/import-stack-overflow
