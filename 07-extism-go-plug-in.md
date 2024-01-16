# 07-Create an Extism plug-ins with TinyGo
> - run the plugins with the Extism CLI
> - use the host functions provided by the PDK

Extism Go PDK: https://github.com/extism/go-pdk

## Create a new project with source code

```bash
mkdir 07-extism-go-plug-in
cd 07-extism-go-plug-in
go mod init hello-extism-go
```

## Create a main.go file

```bash
touch main.go
```

```golang
package main

//export hello()
func hello() int32 {
  // read the function argument from the memory

  // display the argument

  // read the config object
  
  // display the value of url

  // use the url to make an http request

  // display the response

  // return a message: output := "🎉 Extism is 💜"


  return 0
}

func main() {}
```
> - `//export hello()`, then the function is "visible" for the host
> - the empty `main` function is mandatory
> - `int32`: the wasm functions return a number (if you make a void function it should work) 

## Install the Go plug-in development kit

```bash
go get github.com/extism/go-pdk
```

## Specifications

The wasm plugin will be launched with the [Extism CLI](https://github.com/extism/cli), like this:

```bash
extism call go-plugin.wasm hello \
  --input "😀 Hello World 🌍! (from TinyGo)" \
  --log-level "info" \
  --allow-host "*" \
  --set-config '{"url":"https://jsonplaceholder.typicode.com/todos/1"}' \
  --wasi
```

> We call the `hello` function in the wasm plugin
> - `--input` will pass the string to the plugin
> - `--log-level` will set the log level
> - `--allow-host "*"` will allow the plugin to make http requests
> - `--set-config` will set the configuration for the plugin (another way to pass data to the plugin)
> - `--wasi` will enable WASI

The function will:

- read the function argument from the memory
- display the result
- get the value of `url` from the config object
- display the value of `url`
- make a `GET` HTTP request to the `url`
- display the response
- return a value to the host (the CLI)


## Update the source code / how to use the Extism PDK

### How to read the function argument fro the shared memory

```go
// read the function argument from the memory
input := pdk.Input()
```

### How to display the result

```go
// display input
pdk.Log(pdk.LogInfo, string(input))
```

### How to read a key from the config object

```go
// read a key from the config object
url, _ := pdk.GetConfig("url")
```

### How to make an HTTP request

```go
// make an HTTP request
req := pdk.NewHTTPRequest("GET", url)
res := req.Send()
// you can get the result with res.Body()
```

### How to "return" a value to the Extism CLI

```go
// return a message
// copy string to host memory
mem := pdk.AllocateString("🎉 Extism is 💜")
pdk.OutputMemory(mem)
```

## Build the plugin

```bash
tinygo build -scheduler=none --no-debug \
  -o go-plugin.wasm \
  -target wasi main.go

ls -lh *.wasm
```

... And, then, run the plugin with the Extism CLI 🚀
