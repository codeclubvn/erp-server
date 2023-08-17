# Powershell

# Build
docker compose up -d --force-recreate --build

# Remove dangling images
$noneImages = docker images -f "dangling=true" -q
$numOfNoneImages = $noneImages | wc -w
Write-Output "Docker has $numOfNoneImages dangling image(s)"

if ($numOfNoneImages -gt 0) {
    docker rmi $noneImages
    Write-Output "Remove $numOfNoneImages dangling image(s) in Docker successfully"
}

# Remove dangling volumes
$noneVolumes = docker volume ls -f "dangling=true" -q
$numOfNoneVolumes = $noneVolumes | wc -w
Write-Output "Docker has $numOfNoneVolumes dangling volume(s)"

if ($numOfNoneVolumes -gt 0) {
    docker volume rm $noneVolumes
    Write-Output "Remove $numOfNoneVolumes dangling volume(s) in Docker successfully"
}
