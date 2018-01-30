#!/usr/bin/env pwsh
Set-StrictMode -Version Latest
$ErrorActionPreference = "Stop"
function exitIfFailed { if ($LASTEXITCODE -ne 0) { exit } }

# Each chapter and article should have a unique id
# Run this script to generate a random unique id that has a reasonably compact
# string representation

go run ./cmd/gen-id/genid.go
