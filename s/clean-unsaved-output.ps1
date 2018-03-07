#!/usr/bin/env pwsh
Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"
function exitIfFailed { if ($LASTEXITCODE -ne 0) { exit } }

Remove-Item "cached_output\*" | Where-Object { ! $_.PSIsContainer }

Set-Location -Path cached_output

git checkout ./

Set-Location -Path ..
