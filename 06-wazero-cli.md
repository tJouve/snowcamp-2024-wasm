# Make you own wasm CLI
> With the wazero runtime

## Main objectives
- Load the wasm file: `./function/demo.wasm`
- Instantiate the wasm module
- Call the `hello` wasm function

### `hello` wasm function
> 👋 do not look at the source code of the function now

- Arguments: `strPos uint32, strSize uint32` (position and size of the string into the wasm memory)
- Return: `uint64` (position and size of the returned value into the wasm memory ... packed in only one `uint64`)

That means:
- The CLI will copy the string into the wasm memory and get the position and size
- The CLI will call the wasm function with the position and size
- The CLI will read the returned value from the wasm memory
- The CLI will print the returned value

## Create a new project with source code
> inspiration: https://github.com/tetratelabs/wazero/blob/main/examples/allocation/tinygo/greet.go
```bash
mkdir 06-wasm-cli
cd 06-wasm-cli
go mod init wasm-cli
touch main.go
```

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

func main() {

	// Create instance of a wazero runtime
	ctx := context.Background()

	// Create a new runtime.
	runtime := wazero.NewRuntime(ctx)

	// This closes everything this Runtime created.
	defer runtime.Close(ctx)

	// Use WASI
	wasi_snapshot_preview1.MustInstantiate(ctx, runtime)

	// Load the WebAssembly module
	wasmPath := "../function/demo.wasm"
	wasmDemo, err := os.ReadFile(wasmPath)

	if err != nil {
		log.Panicln(err)
	}

	// Instantiate the Wasm plugin/program.
	module, err := runtime.Instantiate(ctx, wasmDemo)
	if err != nil {
		log.Panicln(err)
	}
	// These function are exported by TinyGo
	malloc := module.ExportedFunction("malloc")
	free := module.ExportedFunction("free")

	// Get the reference to the Wasm function: "hello"
	helloFunction := module.ExportedFunction("hello")

	// Passing parameters to the Wasm function: "hello"
	// Function argument
	name := "Jane Doe"
	nameSize := uint64(len(name))

	// Allocate Memory for "Jane Doe"
	results, err := malloc.Call(ctx, nameSize)
	if err != nil {
		log.Panicln(err)
	}
	namePosition := results[0]

	// Free the pointer when finished
	defer free.Call(ctx, namePosition)

	// Copy "Jane Doe" to memory
	success := module.Memory().Write(uint32(namePosition), []byte(name))
	if !success {
		log.Panicf("out of range of memory size")
	}

	// Call hello(pos, size)
	// Call "hello" with the position and the size of "Jane Doe" (into the memory)
	// The result type is []uint64
	result, err := helloFunction.Call(ctx, namePosition, nameSize)
	if err != nil {
		log.Panicln(err)
	}

	// Extract the position and size of from result
	valuePosition := uint32(result[0] >> 32) // shift right bit operation
	valueSize := uint32(result[0]) // AND mask operation

	// Read the value from the memory
	bytes, ok := module.Memory().Read(valuePosition, valueSize)

	if !ok {
		log.Panicf("😡 Out of range of memory size")
	} else {
		fmt.Println("😃 Returned value :", string(bytes))
	}

}
```

## `hello` wasm function
> inspiration: https://github.com/tetratelabs/wazero/blob/main/examples/allocation/tinygo/testdata/greet.go

- Have a look at the source code of the `hello` function in `./function/main.go`

## `hey` wasm function
- Add a `hey` function
- Update the CLI to call the `hey` function (and the `hello` function)
