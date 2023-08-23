package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type returnValue struct {
  value string
}

func main() {
  app := fiber.New()

  app.Get("/", func (c *fiber.Ctx) error {
    return c.SendString("Hello, World!")
  })
    
  app.Get("/json", func(c *fiber.Ctx) error {
    data := returnValue{value: "Hello World"}

    return c.JSON(data)

    return c.JSON(fiber.Map{
      "value": "Hello World",
    })
  })

  port := os.Getenv("PORT")
  if port == "" {
    port = "3000"
  }

log.Fatal(app.Listen("0.0.0.0:" + port))
}
