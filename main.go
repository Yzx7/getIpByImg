package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!, tu ip es:" + c.IP())
	})

	app.Get("/poderosaimagen", func(c *fiber.Ctx) error {
		// "127.0.0.1"
		fmt.Println(c.IP())
		return c.SendFile("./fancygopher.jpg")
		// ...
	})

	log.Fatal(app.Listen(":3000"))
}
