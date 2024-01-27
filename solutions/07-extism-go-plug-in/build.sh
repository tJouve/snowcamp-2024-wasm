#!/bin/bash
tinygo build -scheduler=none --no-debug \
  -o go-plugin.wasm \
  -target wasi main.go

ls -lh *.wasm
