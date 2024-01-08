# 05-call-functions-go

## Create a new project with source code

```bash
mkdir 05-call-functions-go
cd 05-call-functions-go
go mod init call-functions-go
touch main.go
```

```go
package main

// ✋ make this function callable from host
//export add
func add(x int, y int) int {
  return x + y
}

func main() {}
```

## Build 

```bash
tinygo build -o main.wasm -target wasi ./main.go

ls -lh *.wasm
```

## Run: call the function `add``
> Ref https://wasmedge.org/docs/start/build-and-run/cli/

```bash
wasmedge --reactor main.wasm add 38 4
```

## Add an `hello` function

```go
//export hello
func hello(name string) string {
  return "hello " + name
}
```

## Build, then, run: call the function `hello`
```bash
wasmedge --reactor main.wasm hello "Bob"
```
