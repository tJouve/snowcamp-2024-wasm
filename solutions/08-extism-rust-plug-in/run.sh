#!/bin/bash
extism call ./target/wasm32-wasi/release/hello_extism_rust.wasm \
  hello \
  --input "👩 Jane Doe" \
  --log-level "info" \
  --set-config '{"text":"Hello I am Jane Doe 😊"}' \
  --wasi

echo ""
extism call ./target/wasm32-wasi/release/hello_extism_rust.wasm \
  hello \
  --input "👨 John Doe" \
  --set-config '{"text":"Hello I am John Doe 👋"}' \
  --log-level "info" \
  --wasi
echo ""
