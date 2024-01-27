#!/bin/bash
extism call hello-js.wasm \
  hello \
  --input "👩 Jane Doe" \
  --log-level "info" \
  --wasi

