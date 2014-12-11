#
$serviceName = if($deployed.serviceName) { $deployed.serviceName } else { $deployed.name }

# Verify if service has been installed
$service = Get-WmiObject Win32_Service -Filter ("name='$serviceName'")
if(!$service) {
    Write-Host "Cannot uninstall service [$serviceName] because it cannot be found."
    Exit 1
}

# Uninstall service
Write-Host "Uninstalling service [$serviceName]"
$service.Delete() | Out-Null

# Remove folder
Write-Host "Removing service folder from [$($deployed.targetPath)]."
Remove-Item -Recurse -Force $deployed.targetPath | Out-Null
