go build
if ($LASTEXITCODE -ne 0) {
    Write-Host "Build Failed."
    exit 1
}
.\gobackend
