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

	args := os.Args[1:]
	wasmFilePath := args[0]
	functionName := args[1]
	input := args[2]
	config := args[3]

	
	// config is a json string, convert it to a map
	var configMap map[string]string
	err := json.Unmarshal([]byte(config), &configMap)
	if err != nil {
		fmt.Println("Error converting config to map:", err)
		os.Exit(1)
	}

	// Plugin config
	pluginConfig := extism.PluginConfig{
		ModuleConfig: wazero.NewModuleConfig().WithSysWalltime(),
		EnableWasi:   true,
		LogLevel:     extism.LogLevelInfo,
	}

	// Plugin manifest
	pluginManifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmFile{Path: wasmFilePath},
		},
		AllowedHosts: []string{"*"},
		Config:       configMap,
	}

	// Create a plugin instance
	wasmPlugin, err := extism.NewPlugin(ctx, pluginManifest, pluginConfig, nil)

	if err != nil {
		panic(err)
	}

	// Call the function of the plugin
	_, result, err := wasmPlugin.Call(
		functionName,
		[]byte(input),
	)

	if err != nil {
		fmt.Println("😡", err)
		os.Exit(1)
	} else {
		fmt.Println("🙂", string(result))
		os.Exit(0)
	}

}
