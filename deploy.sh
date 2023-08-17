#!/bin/bash

# Build
docker compose up -d --force-recreate --build

# Remove dangling images
none_images=$(docker images -f "dangling=true" -q)
num_of_none_images=$(echo "$none_images" | wc -w)
echo "Docker has '$num_of_none_images' dangling image(s)"

if [ "$num_of_none_images" -gt 0 ]; then
    docker rmi "$none_images"
    echo "Remove '$num_of_none_images' dangling image(s) in Docker successfully"
fi

# Remove dangling volumes
none_volumes=$(docker volume ls -f "dangling=true" -q)
num_of_none_volumes=$(echo "$none_volumes" | wc -w)
echo "Docker has '$num_of_none_volumes' dangling volume(s)"

if [ "$num_of_none_volumes" -gt 0 ]; then
    docker volume rm "$none_volumes"
    echo "Remove '$num_of_none_volumes' dangling volume(s) in Docker successfully"
fi
