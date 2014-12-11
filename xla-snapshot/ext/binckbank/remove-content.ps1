# Remove content
Write-Host "Removing content from [$($deployed.targetPath)]."
Remove-Item -Recurse -Force $deployed.targetPath | Out-Null
