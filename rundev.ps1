go build -o dist/gobackend.exe
if ($LASTEXITCODE -ne 0) {
    Write-Host "Build Failed."
    exit 1
}
.\dist\gobackend.exe
