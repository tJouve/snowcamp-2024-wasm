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

	httpPort := "8081"


	app := fiber.New(fiber.Config{DisableStartupMessage: true})

	app.Post("/", func(c *fiber.Ctx) error {

		params := c.Body()

		return c.SendString("hello")

	})

	fmt.Println("🌍 http server is listening on:", httpPort)
	app.Listen(":" + httpPort)
}
