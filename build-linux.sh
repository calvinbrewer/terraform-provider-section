#!/usr/bin/env bash

cd "$(dirname "$0")" || exit

image_name=terraform-provider-section-builder:latest

mkdir bin || exit

docker build -t "${image_name}" . || exit

cid=$(docker create "${image_name}") || exit

docker cp "${cid}:/go/bin/terraform-provider-section" ./bin/ || exit

docker rm "${cid}" >/dev/null
