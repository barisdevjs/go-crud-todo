package main

import (
	"github.com/gofiber/fiber/v2"
)

func GetOptions(c *fiber.Ctx) error {
	// Set the allowed HTTP methods for the route
	c.Set("Allow", "GET, POST, PUT, DELETE, OPTIONS")

	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	return c.SendStatus(fiber.StatusNoContent)
}
