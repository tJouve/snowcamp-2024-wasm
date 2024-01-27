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
	valuePosition := uint32(result[0] >> 32)
	valueSize := uint32(result[0])

	// Read the value from the memory
	bytes, ok := module.Memory().Read(valuePosition, valueSize)

	if !ok {
		log.Panicf("😡 Out of range of memory size")
	} else {
		fmt.Println("😃 Returned value :", string(bytes))
	}

}
