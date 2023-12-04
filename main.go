package main

import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World! versi 1")
    })

	app.Get("/salah", func(c *fiber.Ctx) error {
        return fiber.ErrServiceUnavailable

    	// 503 On vacation!
   		 return fiber.NewError(fiber.StatusServiceUnavailable, "On vacation!")
    })

    app.Listen(":8080")
}