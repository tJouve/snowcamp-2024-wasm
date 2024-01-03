# 02-hello-go

```bash
mkdir 02-hello-go
cd 02-hello-go
go mod init hello-go
touch main.go
```

```go
package main

import (
	"fmt"
	"os" 
)

func main() {

	fmt.Println("👋 Hello World from Go 🌍")
	args := os.Args
	argsWithoutCaller := os.Args[1:]

	fmt.Println(args)
	fmt.Println(argsWithoutCaller)

}
```


```bash
tinygo build -o main.wasm -target wasi ./main.go

ls -lh *.wasm
```

## Run

```bash
wasmtime main.wasm hello world
wasmedge main.wasm hello world
wazero run main.wasm hello world
```

## Change the source code
> of the main function

```go
// Create a new scanner to read from standard input
scanner := bufio.NewScanner(os.Stdin)

fmt.Println("Enter some text (press Ctrl+D or Ctrl+Z to end):")

// Read input line by line
for scanner.Scan() {
    text := scanner.Text() // Get the current line of text
    if text == "" {
        break // Exit loop if an empty line is entered
    }
    fmt.Println("You entered:", text)
}

if err := scanner.Err(); err != nil {
    fmt.Println("Error:", err)
}
```

> Re-build and re-run
