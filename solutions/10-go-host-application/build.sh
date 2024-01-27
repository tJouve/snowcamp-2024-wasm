#!/bin/bash
# The -ldflags="-s -w" flag is used to optimize the size and reduce the noise of the generated binary. 
go build -ldflags="-s -w"
ls -lh hostapp
