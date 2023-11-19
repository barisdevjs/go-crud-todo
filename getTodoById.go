package main

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTodo(c *fiber.Ctx) error {
	todoID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.SendStatus(400)
	}

	query := bson.D{{Key: "_id", Value: todoID}}
	result := mg.Db.Collection("Todo-1").FindOne(c.Context(), &query)

	var todo Todo
	err = result.Decode(&todo)
	if err != nil {
		return c.SendStatus(500)
	}

	return c.Status(200).JSON(todo)

}
