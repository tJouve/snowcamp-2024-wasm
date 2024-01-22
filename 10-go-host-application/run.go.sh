#!/bin/bash
go run main.go \
../07-extism-go-plug-in/go-plugin.wasm \
hello \
"👋 John Doe" \
'{"url":"https://jsonplaceholder.typicode.com/todos/1"}'
