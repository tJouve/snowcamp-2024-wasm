#!/bin/bash
extism call go-plugin.wasm hello \
  --input "😀 Hello World 🌍! (from TinyGo)" \
  --log-level "debug" \
  --allow-host "*" \
  --set-config '{"url":"https://jsonplaceholder.typicode.com/todos/1"}' \
  --wasi
echo ""
