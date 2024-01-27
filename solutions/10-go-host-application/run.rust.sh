#!/bin/bash
go run main.go \
../08-extism-rust-plug-in/target/wasm32-wasi/release/hello_extism_rust.wasm \
hello \
"👩 Jane Doe" \
'{"text":"Hello I am Jane Doe 😊"}'
