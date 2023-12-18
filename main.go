package main

import "github.com/gofiber/fiber/v2"

import "fmt"

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World! versi 5")
    })

	app.Get("/salah", func(c *fiber.Ctx) error {
        return fiber.ErrServiceUnavailable

    	// 503 On vacation!
   		 return fiber.NewError(fiber.StatusServiceUnavailable, "On vacation!")
    })

    app.Post("/test", func(c *fiber.Ctx) error {
        // Get raw body from POST request:
        fmt.Println(string(c.BodyRaw()))
        return c.Send(c.BodyRaw()) // []byte("user=john")
      })

    app.Listen(":8080")
}