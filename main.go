package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// app := fiber.New()
	app := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		Prefork:      true,
	})

	if fiber.IsChild() {
		fmt.Println("Child process")
	} else {
		fmt.Println("Parent process")
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	err := app.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}
