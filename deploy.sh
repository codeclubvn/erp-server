#!/bin/bash

docker compose up -d --force-recreate --build

none_images=$(docker images -f "dangling=true" -q)
num_of_none_images=$(echo "$none_images" | wc -w)
echo "Docker has '$num_of_none_images' dangling image(s)"

if [ "$num_of_none_images" -gt 0 ]; then
    docker rmi "$none_images"
    echo "Remove '$num_of_none_images' dangling image(s) in Docker successfully"
fi
