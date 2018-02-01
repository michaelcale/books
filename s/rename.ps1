#!/usr/bin/env pwsh
Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"
function exitIfFailed { if ($LASTEXITCODE -ne 0) { exit } }

# chapters and articles are order by their number prefix (0030-* etc.)
# we space them by 10 so that we can squeeze an article or 2 when
# we move articles/chapters around or when we add a new one
# this script re-balances the numbers to be evenly spaced by 10 again

Remove-Item -Force -ErrorAction SilentlyContinue ./cmd/rename/rename

Set-Location -Path cmd/rename
go build -o rename
Set-Location -Path ../..
exitIfFailed

./cmd/rename/rename
Remove-Item -Force -ErrorAction SilentlyContinue ./cmd/rename/rename
