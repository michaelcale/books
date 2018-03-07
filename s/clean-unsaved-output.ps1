#!/usr/bin/env pwsh
Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"
function exitIfFailed { if ($LASTEXITCODE -ne 0) { exit } }

Set-Location -Path cached_output

rm -rf *
git checkout ./

Set-Location -Path ..
