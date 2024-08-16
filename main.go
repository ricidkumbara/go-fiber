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

	// Todo
	app.Use(func(c *fiber.Ctx) error {
		fmt.Println("Middleware before")
		err := c.Next()
		fmt.Println("Middleware after")
		return err
	})

	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("Middleware before for /api")
		err := c.Next()
		fmt.Println("Middleware after for /api")
		return err
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})
	app.Get("/api/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	if fiber.IsChild() {
		fmt.Println("Child process")
	} else {
		fmt.Println("Parent process")
	}

	err := app.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}
