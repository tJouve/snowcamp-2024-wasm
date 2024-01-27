package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	extism "github.com/extism/go-sdk"
	"github.com/gofiber/fiber/v2"
	"github.com/tetratelabs/wazero"
)

// store all your plugins in a normal Go hash map, protected by a Mutex
// (reproduce something like the node.js event loop)
// to avoid "memory collision 💥"
var m sync.Mutex
var plugins = make(map[string]*extism.Plugin)

func StorePlugin(plugin *extism.Plugin) {
	plugins["code"] = plugin
}

func GetPlugin() (extism.Plugin, error) {
	if plugin, ok := plugins["code"]; ok {
		return *plugin, nil
	} else {
		return extism.Plugin{}, errors.New("🔴 no plugin")
	}
}

func main() {
	ctx := context.Background()

	args := os.Args[1:]
	wasmFilePath := args[0]
	functionName := args[1]
	httpPort := args[2]
	config := args[3]

	var configMap map[string]string
	err := json.Unmarshal([]byte(config), &configMap)
	if err != nil {
		fmt.Println("Error converting config to map:", err)
		os.Exit(1)
	}

	pluginConfig := extism.PluginConfig{
		ModuleConfig: wazero.NewModuleConfig().WithSysWalltime(),
		EnableWasi:   true,
		LogLevel:     extism.LogLevelInfo,
	}

	pluginManifest := extism.Manifest{
		Wasm: []extism.Wasm{
			extism.WasmFile{
				Path: wasmFilePath},
		},
		AllowedHosts: []string{"*"},
		Config:       configMap,
	}

	pluginInst, err := extism.NewPlugin(ctx, pluginManifest, pluginConfig, nil) // new
	if err != nil {
		log.Println("🔴 !!! Error when loading the plugin", err)
		os.Exit(1)
	}

	//-----------------------------------------
	StorePlugin(pluginInst)
	//-----------------------------------------

	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	app.Post("/", func(c *fiber.Ctx) error {

		params := c.Body()

		m.Lock()
		// don't forget to release the lock on the Mutex
		defer m.Unlock()
		pluginInst, err := GetPlugin()


		_, out, err := pluginInst.Call(functionName, params)

		if err != nil {
			fmt.Println(err)
			c.Status(http.StatusConflict)
			return c.SendString(err.Error())
		} else {
			c.Status(http.StatusOK)

			return c.SendString(string(out))
		}

	})

	fmt.Println("🌍 http server is listening on:", httpPort)
	app.Listen(":" + httpPort)
}
