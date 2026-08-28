$ErrorActionPreference = "Stop"

$moduleRoot = Split-Path $PSScriptRoot -Parent
$repositoryRoot = Split-Path $moduleRoot -Parent | Split-Path -Parent
$iasiSource = Join-Path $repositoryRoot "iasi"
$versionSource = Join-Path $repositoryRoot "VERSION"
$destination = Join-Path $moduleRoot "internal\source\embedded\iasi"
$versionDestination = Join-Path $moduleRoot "internal\source\embedded\VERSION"

if (-not (Test-Path $iasiSource -PathType Container)) {
    throw "IASI source not found: $iasiSource"
}
if (-not (Test-Path $versionSource -PathType Leaf)) {
    throw "Version source not found: $versionSource"
}

Remove-Item $destination -Recurse -Force -ErrorAction SilentlyContinue
Remove-Item (Join-Path $moduleRoot "internal\source\embedded\agentics") -Recurse -Force -ErrorAction SilentlyContinue
Remove-Item (Join-Path $moduleRoot "internal\source\embedded\adapters") -Recurse -Force -ErrorAction SilentlyContinue
New-Item (Split-Path $destination -Parent) -ItemType Directory -Force | Out-Null
Copy-Item $iasiSource $destination -Recurse
Copy-Item $versionSource $versionDestination -Force