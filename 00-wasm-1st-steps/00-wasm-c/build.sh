#!/bin/bash
clear
bat $0 --line-range 5:
echo ""
clang --target=wasm32 --no-standard-libraries -Wl,--export-all -Wl,--no-entry -o main.wasm main.c

ls -lh *.wasm
