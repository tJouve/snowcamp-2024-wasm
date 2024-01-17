#!/bin/bash
extism call ./target/wasm32-wasi/release/hello_extism_rust.wasm \
  hello \
  --input "👩 Jane Doe" \
  --wasi
echo ""
extism call ./target/wasm32-wasi/release/hello_extism_rust.wasm \
  hello \
  --input "👨 John Doe" \
  --wasi
echo ""
