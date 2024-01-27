package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	extism "github.com/extism/go-sdk"
	"github.com/gofiber/fiber/v2"
	"github.com/tetratelabs/wazero"
)

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

	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	app.Post("/", func(c *fiber.Ctx) error {

		params := c.Body()

		pluginInst, err := extism.NewPlugin(ctx, pluginManifest, pluginConfig, nil) // new
		if err != nil {
			log.Println("🔴 !!! Error when loading the plugin", err)
			os.Exit(1)
		}

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
