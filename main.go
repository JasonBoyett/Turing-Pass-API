package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"

  "github.com/gofiber/fiber/v2/middleware/cors"
)


func main() {

  type returnValue struct {
    property string
    someOtherProperty int
  }

	app := fiber.New()

  app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

  app.Get("/jsonp", func(c *fiber.Ctx) error {
    data := returnValue{property: "value", someOtherProperty: 123}
    log.Println(data)
    callbackFunc := c.Query("callback")

    if callbackFunc != "" {
      return c.JSONP(
        returnValue{
          property: "value",
          someOtherProperty: 123,
        },
        callbackFunc)
    } else {
      return c.JSONP(
        returnValue{
          property: "value",
          someOtherProperty: 123,
        },
      )
    }
  })

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))
}
