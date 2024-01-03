#!/bin/bash
IMAGE_BASE_NAME="wasm-wasi"
IMAGE_TAG="0.0.2"
docker login -u ${DOCKER_USER} -p ${DOCKER_PWD}
docker buildx build --platform linux/amd64 --push -t ${DOCKER_USER}/${IMAGE_BASE_NAME}:${IMAGE_TAG} .

