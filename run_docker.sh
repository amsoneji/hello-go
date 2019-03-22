#!/bin/bash

# Script helping users run a go docker image to encapsulate the runtime.
# 
#   This script helps to remove the need for maintaining your own local go environment and runtime.
#   Through the use of a golang:1.11.4 docker image, as well as local bind mounts, the container
#   is able to access your files in this directory directly, and is able to run them natively.

docker run -it --rm \
  --name hello-go \
  --mount type=bind,source="$(pwd)",target=/go/src/app \
  --mount type=bind,source="$(pwd)/.bashrc",target=/root/.bashrc \
  --workdir="/go/src/app/" \
  -p 3000:3000 \
  golang:1.11.6 \
  /bin/bash