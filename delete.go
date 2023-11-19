package main

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteTodo(c *fiber.Ctx) error {
	todoID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.SendStatus(400)
	}

	query := bson.D{{Key: "_id", Value: todoID}}
	result, err := mg.Db.Collection("Todo-1").DeleteOne(c.Context(), &query)

	if err != nil {
		return c.SendStatus(500)
	}

	if result.DeletedCount < 1 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON("record deleted")
}
