#!/bin/bash

cat > ./03-go-plugin/main.go <<- EOM
package main
// go

//export say_hello
func say_hello() {

	// read input

	// read config

	// use a host function to make a request 

	// create output

}

func main() {}

EOM


cat > ./05-go-host-application/main.go <<- EOM
package main
//host

import (
	"context"
	"os"
)

func main() {

	ctx := context.Background()

	args := os.Args[1:]
	wasmFilePath := args[0]
	functionName := args[1]
	input := args[2]

	// Plugin config

	// Plugin manifest

	// Create a plugin instance

	// Call the function of the plugin

}

EOM