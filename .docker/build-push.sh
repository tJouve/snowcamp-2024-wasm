#!/bin/bash
#IMAGE_BASE_NAME="snowcamp-2024-wasm-lite"
#IMAGE_TAG="0.0.1"
#docker login -u ${DOCKER_USER} -p ${DOCKER_PWD}
#docker buildx build --platform linux/amd64 --push -t ${DOCKER_USER}/${IMAGE_BASE_NAME}:${IMAGE_TAG} .

set -o allexport; source .env; set +o allexport

IMAGE_BASE_NAME="snowcamp-2024-wasm-lite"
IMAGE_TAG="0.0.2"
docker login -u ${DOCKER_USER} -p ${DOCKER_PWD}
docker buildx build \
--build-arg="GO_ARCH=${GO_ARCH}" \
--build-arg="GO_VERSION=${GO_VERSION}" \
--build-arg="TINYGO_ARCH=${TINYGO_ARCH}" \
--build-arg="TINYGO_VERSION=${TINYGO_VERSION}" \
--build-arg="EXTISM_VERSION=${EXTISM_VERSION}" \
--build-arg="EXTISM_ARCH=${EXTISM_ARCH}" \
--build-arg="EXTISM_JS_VERSION=${EXTISM_JS_VERSION}" \
--build-arg="EXTISM_JS_ARCH=${EXTISM_JS_ARCH}" \
--build-arg="EXTISM_JS_OS=${EXTISM_JS_OS}" \
--build-arg="BINARYEN_VERSION=${BINARYEN_VERSION}" \
--build-arg="BINARYEN_ARCH=${BINARYEN_ARCH}" \
--build-arg="BINARYEN_OS=${BINARYEN_OS}" \
--build-arg="NODE_MAJOR=${NODE_MAJOR}" \
--platform linux/amd64 \
--push -t ${DOCKER_USER}/${IMAGE_BASE_NAME}:${IMAGE_TAG} .

echo "to test it:"
echo "docker run -it ${DOCKER_USER}/${IMAGE_BASE_NAME}:${IMAGE_TAG}"