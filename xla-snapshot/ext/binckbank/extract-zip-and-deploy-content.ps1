# Determine temporary extraction dir
$targetDir = [System.IO.Path]::GetDirectoryName($deployed.targetPath)
$targetName = $([System.IO.Path]::GetFileName($deployed.targetPath))

$unzipped = "$([System.IO.Path]::GetDirectoryName($deployed.file))\$targetName"
New-Item -Type Directory $unzipped | Out-Null

#
Write-Host "Unzipping [$($deployed.file)] to [$unzipped]."
$shellApplication = New-Object -Com shell.application
$zipNS = $shellApplication.NameSpace($deployed.file)
$unzippedNS = $shellApplication.NameSpace($unzipped)
$unzippedNS.CopyHere($zipNS.Items()) | Out-Null

#
If(!(Test-Path $deployed.targetPath)) {
	Write-Host "Creating target directory [$($deployed.targetPath)]."
	New-Item -Type Directory $deployed.targetPath | Out-Null
}

#
Write-Host "Comparing new content at [$unzipped] with existing content at [$($deployed.targetPath)]."
$existingFiles = @(Get-Childitem -Recurse $deployed.targetPath)
$newFiles = @(Get-ChildItem -Recurse $unzipped)
$filesToRemove = @(Compare-Object $existingFiles $newFiles | Where-Object {$_.SideIndicator -eq "<=" })

#
Write-Host "Copying added and modified files from [$unzipped] to [$targetDir\$targetName]."
Copy-Item -Recurse -Force $unzipped $targetDir

#
if($filesToRemove) {
	Write-Host "Removing deleted files."
	for ($i = $filesToRemove.length - 1; $i -ge 0; $i--) {
		$f = $filesToRemove[$i].InputObject
		Write-Host "Removing [$($f.FullName)]."
		Remove-Item -Recurse -Force -Path $f.FullName
	}
}
