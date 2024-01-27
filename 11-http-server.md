# Create an HTTP server to server "wasm micro services"

> various ways to do that

## HTTP server skeleton

This is the source code of an HTTP Server using the **fiber** framework

```go
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

	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	app.Post("/", func(c *fiber.Ctx) error {

		params := c.Body()

		c.Status(http.StatusOK)

		return c.SendString("hello")

	})

	fmt.Println("🌍 http server is listening on:", httpPort)
	app.Listen(":" + httpPort)
}
```

With what you learn with the exercise number `10`, you can now create a simple HTTP server that will serve the **wasm plugins** done in the previous exercise (`07`, `08`, `09`).

The `hello` function of the plug-ins will be called at every HTTP request.

- To build the server, see: `./build.sh`
- To run the server, see: `./start.go.sh`, `./start.js.sh`, `./start.rust.sh`
- **Test the curl request** for each plugin, see: `./query.sh`
