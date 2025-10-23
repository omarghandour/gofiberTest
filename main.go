package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Test",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
		x := 10
		p := &x
		*p = 20
		fmt.Println(x)
		return c.SendString("zHello,  .World! .")
	})

	app.Listen(":3001")
}
