# Powershell

docker compose up -d --force-recreate --build

$noneImages = docker images -f "dangling=true" -q
$numOfNoneImages = $noneImages | wc -w
Write-Output "Docker has $numOfNoneImages dangling image(s)"

if ($numOfNoneImages -gt 0) {
    docker rmi $noneImages
    Write-Output "Remove $numOfNoneImages dangling image(s) in Docker successfully"
}
