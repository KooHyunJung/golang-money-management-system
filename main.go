package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/start", func(c *fiber.Ctx) error {
		const a int = 10
		const b float64 = 15.12345
		var sum float64 = float64(a) + b
		fmt.Println(sum)

		return c.SendString("")
	})

	app.Listen(":3000")
}

//  참고 문헌 : https://docs.gofiber.io/
// https://go.dev/
// https://github.com/gofiber/fiber
