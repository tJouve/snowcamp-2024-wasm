#!/bin/bash
./best-wasm-server \
../08-extism-rust-plug-in/target/wasm32-wasi/release/hello_extism_rust.wasm \
hello \
8081 \
'{"text":"Hello I am Jane Doe 😊"}'

