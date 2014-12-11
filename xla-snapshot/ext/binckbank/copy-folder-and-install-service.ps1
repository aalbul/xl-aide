#
$serviceName = if($deployed.serviceName) { $deployed.serviceName } else { $deployed.name }
$displayName = if($deployed.serviceDisplayName) { $deployed.serviceDisplayName } else { $serviceName }
$description = if($deployed.serviceDescription) { $deployed.serviceDescription } else { $serviceName }

# Remove old service folder if it's still there
if (Test-Path $deployed.targetPath) {
	Write-Host "Removing old content from [$($deployed.targetPath)]."
	Remove-Item -Recurse -Force $deployed.targetPath
}

# Copy new service folder
Write-Host "Copying service folder for [$serviceName] to [$($deployed.targetPath)]."
Copy-Item -Recurse -Force $deployed.file $deployed.targetPath

# Install service
Write-Host "Installing service [$serviceName]."

if($deployed.username -and $deployed.password) {
    $securePassword = $deployed.password | ConvertTo-SecureString -asPlainText -Force
    $cred = New-Object System.Management.Automation.PSCredential($deployed.username, $securePassword)
    New-Service -Name $serviceName -BinaryPathName $deployed.binaryPathName -DependsOn $deployed.dependsOn -Description $description -DisplayName $displayName -StartupType $deployed.startupType -Credential $cred
} else {
    New-Service -Name $serviceName -BinaryPathName $deployed.binaryPathName -DependsOn $deployed.dependsOn -Description $description -DisplayName $displayName -StartupType $deployed.startupType | Out-Null
}
