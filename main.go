package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

  "github.com/gofiber/fiber/v2/middleware/cors"
)

type returnValue struct {
	property string
}

func main() {
	app := fiber.New()

  app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/jsonp", func(c *fiber.Ctx) error {
		// data := returnValue{value: "Hello World"}

		return c.JSONP(returnValue{property: "Hello World"})

	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}
