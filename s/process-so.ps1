#!/usr/bin/env pwsh
Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"
function exitIfFailed { if ($LASTEXITCODE -ne 0) { exit } }

Remove-Item -Force -ErrorAction SilentlyContinue ./process_so
go build -o process_so
exitIfFailed

./process_so
Remove-Item -Force -ErrorAction SilentlyContinue ./process_so
