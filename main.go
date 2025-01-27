package main

import (
	"log"
	"otfch_be/db_connection"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//Initializing the Application
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	app.Listen(":3000")

	//Initializing the database
	pool, err := db_connection.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer pool.Close()
}
