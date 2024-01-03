#!/bin/bash
# The -ldflags="-s -w" flag is used to optimize the size and reduce the noise of the generated binary. 
clear
bat $0 --line-range 5:
echo ""
go build -ldflags="-s -w"
ls -lh hostapp

