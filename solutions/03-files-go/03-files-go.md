# 03-files-go
> Write and read into a file
## Create a new project with source code

```bash
mkdir 03-files-go
cd 03-files-go
go mod init files-go
touch main.go
```

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("😡", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("👋 Hello, World! 🌍")
	if err != nil {
		fmt.Println("😡", err)
		return
	}

	data, err := os.ReadFile("hello.txt")
	if err != nil {
		fmt.Println("😡", err)
	}

	fmt.Print(string(data))

}
```

## Build 

```bash
tinygo build -o main.wasm -target wasi ./main.go

ls -lh *.wasm
```

## Run

```bash
wasmedge main.wasm
```
> Yous should get an error

```bash
wasmedge --dir . main.wasm 
```

> Ref https://wasmedge.org/docs/start/build-and-run/cli/

