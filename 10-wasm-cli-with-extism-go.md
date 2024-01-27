# Create an Extism Host application (a CLI)

The application will need 4 arguments:
- the path to the wasm file
- the name of the exported function
- a string parameter to pass to the function
- a json string to pass a config value

The goal is to be able to re-use and call the 3 previous plugins like this:

```bash
hostapp \
../07-extism-go-plug-in/go-plugin.wasm \
hello \
"👋 John Doe" \
'{"url":"https://jsonplaceholder.typicode.com/todos/1"}'
```

```bash
hostapp \
../08-extism-rust-plug-in/target/wasm32-wasi/release/hello_extism_rust.wasm \
hello \
"👩 Jane Doe" \
'{"text":"Hello I am Jane Doe 😊"}'
```

```bash
hostapp \
go run main.go \
../09-extism-js-plug-in/hello-js.wasm \
hello \
"Bob Morane" \
'{}'
```

## With Go 

To create an instance of the wasm plug-in, the code looks like this:

```go
wasmPlugin, err := extism.NewPlugin(ctx, pluginManifest, pluginConfig, nil)
```
- the type of `pluginManifest` is `extism.Manifest`
- the type of `pluginConfig` is `extism.PluginConfig`

To call the exported function, the code looks like this:

```go
_, result, err := wasmPlugin.Call(
    functionName,
    []byte(input),
)
```

- Read this: https://github.com/extism/go-sdk?tab=readme-ov-file#creating-a-plug-in
- Create the application

### Some help

```golang
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


// configMap is a map[string]string
var configMap map[string]string
err := json.Unmarshal([]byte(config), &configMap)
if err != nil {
    fmt.Println("Error converting config to map:", err)
    os.Exit(1)
}
```


## Build and run the application

You can build the application with:

```bash
go build -ldflags="-s -w"
ls -lh hostapp
```

And then call the `hostapp` application

or run directly the code with `go run main.go`