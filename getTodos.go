package main

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func GetTodos(c *fiber.Ctx) error {
	query := bson.D{{}}
	cursor, err := mg.Db.Collection("Todo-1").Find(c.Context(), query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var todos []Todo = make([]Todo, 0)

	if err := cursor.All(c.Context(), &todos); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.JSON(todos)
}
