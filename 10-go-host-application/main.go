package main

//host

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	extism "github.com/extism/go-sdk"
	"github.com/tetratelabs/wazero"
)

func main() {

	ctx := context.Background()

	/*
		- read the arguments from the command line
		- read the object configuration
		- create a plugin config
		- create a plugin manifest
		- create a plugin instance
		- use the call method of the plugin instance to call the function
		- print the result

	*/

	args := os.Args[1:]


	// foo...

	if err != nil {
		fmt.Println("😡", err)
		os.Exit(1)
	} else {
		fmt.Println("🙂", string(result))
		os.Exit(0)
	}

}
