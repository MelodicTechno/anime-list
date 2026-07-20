# Build script for Windows
param(
    [string]$Output = "bin/anime-list-server.exe"
)

$ErrorActionPreference = "Stop"
Push-Location $PSScriptRoot\..
try {
    go build -o $Output ./cmd/api/
    Write-Host "Build succeeded: $Output"
} finally {
    Pop-Location
}
