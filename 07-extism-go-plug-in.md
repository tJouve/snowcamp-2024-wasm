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

The wasm plugin will be launched with the [Extism CLI](https://github.com/extism/cli)