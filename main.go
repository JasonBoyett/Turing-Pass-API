package main

import (
	"log"
	"os"
	"strconv"
  encrypt "github.com/JasonBoyett/Turing-Pass-API/encrypt"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/cors"
)


func main() {

	app := fiber.New()

  app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, from Turing Pass!")
	})

  app.Get("/jsonp", func(c *fiber.Ctx) error {
    var length int
    callbackFunc := c.Query("callback")
    passWord := c.Query("passWord")
    siteName := c.Query("siteName")
    symbols := c.Query("symbols")
    lengthFromQuery := c.Query("len")

    length, err := strconv.Atoi(lengthFromQuery)
    if err != nil {
      length = 16
    }
    log.Println(length)

    newPass, err := encrypt.Encrypt(passWord, siteName, symbols == "true", length)
    if err != nil {
      return c.JSON(
        fiber.Map{
          "pass": "Error", 
        },
      )
    }

    if callbackFunc != "" {
      return c.JSONP(
        fiber.Map{
          "pass": newPass,
        },
        callbackFunc)
    } else {
      return c.JSON(
        fiber.Map{
          "pass": newPass, 
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
